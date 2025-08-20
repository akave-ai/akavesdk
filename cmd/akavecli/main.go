// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package main provides a command-line interface for managing Akave SDK resources
// such as buckets and files.
package main

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spacemonkeygo/monkit/v3"
	"github.com/spacemonkeygo/monkit/v3/environment"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	mJaeger "storj.io/monkit-jaeger"

	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/version"
	"github.com/akave-ai/akavesdk/sdk"
)

var (
	rootCmd = &cobra.Command{
		Use:   "akavecli",
		Short: "A CLI for managing Akave resources",
	}

	versionCmd = &cobra.Command{
		Use:  "version",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(version.Info().Format())
		},
	}

	bucketCmd = &cobra.Command{
		Use:   "bucket",
		Short: "Manage buckets",
	}

	bucketCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Creates a new bucket",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return NewCmdParamsError(fmt.Sprintf("create bucket command expects exactly 1 argument [bucket name]; got %d", len(args)))
			}

			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}

			return nil
		},
		RunE: cmdCreateBucket,
	}

	bucketDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Removes a bucket",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return NewCmdParamsError(fmt.Sprintf("delete bucket command expects exactly 1 argument [bucket name]; got %d", len(args)))
			}

			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}

			return nil
		},
		RunE: cmdDeleteBucket,
	}

	bucketViewCmd = &cobra.Command{
		Use:   "view",
		Short: "Views a bucket's details",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}

			if len(args) != 1 {
				return NewCmdParamsError(fmt.Sprintf("create bucket command expects exactly 1 argument [bucket name]; got %d", len(args)))
			}

			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}

			return nil
		},
		RunE: cmdViewBucket,
	}

	bucketListCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all buckets",
		Args:  cobra.NoArgs,
		RunE:  cmdListBuckets,
	}

	nodeRPCAddress       string
	privateKey           string
	encryptionKey        string
	maxConcurrency       int
	blockPartSize        int64
	useConnectionPool    bool
	disableErasureCoding bool
	filecoinFlag         bool
	enableSDKMonkitStats bool
	accountName          string

	// tracing.
	mon = monkit.Package()

	tracingAgentAddr = "localhost:6831"
	serviceName      = "akavecli"

	bucketListOffset int64
	bucketListLimit  int64
	fileListOffset   int64
	fileListLimit    int64
)

// CmdParamsError represents an error related to positional arguments.
type CmdParamsError struct {
	Message string
}

// Error returns error message.
func (e *CmdParamsError) Error() string {
	return e.Message
}

// NewCmdParamsError creates new CmdParamsError error.
func NewCmdParamsError(message string) error {
	return &CmdParamsError{Message: message}
}

func init() {
	bucketCmd.AddCommand(bucketCreateCmd)
	bucketCmd.AddCommand(bucketDeleteCmd)
	bucketCmd.AddCommand(bucketViewCmd)
	bucketCmd.AddCommand(bucketListCmd)

	// streaming file API
	fileStreamingCmd.AddCommand(streamingFileListCmd)
	fileStreamingCmd.AddCommand(streamingFileInfoCmd)
	fileStreamingCmd.AddCommand(streamingFileUploadCmd)
	fileStreamingCmd.AddCommand(streamingFileDownloadCmd)
	fileStreamingCmd.AddCommand(streamingFileDeleteCmd)
	fileStreamingCmd.AddCommand(streamingFileVersionsCmd)
	// ipc API
	ipcCmd.AddCommand(ipcBucketCmd)
	ipcCmd.AddCommand(ipcFileCmd)
	ipcBucketCmd.AddCommand(ipcBucketCreateCmd)
	ipcBucketCmd.AddCommand(ipcBucketViewCmd)
	ipcBucketCmd.AddCommand(ipcBucketListCmd)
	ipcBucketCmd.AddCommand(ipcBucketDeleteCmd)
	ipcFileCmd.AddCommand(ipcFileUploadCmd)
	ipcFileCmd.AddCommand(ipcFileDownloadCmd)
	ipcFileCmd.AddCommand(ipcFileListCmd)
	ipcFileCmd.AddCommand(ipcFileInfoCmd)
	ipcFileCmd.AddCommand(ipcFileDeleteCmd)

	// Initialize wallet commands
	initWalletCommands()

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(bucketCmd)
	rootCmd.AddCommand(fileStreamingCmd)
	rootCmd.AddCommand(ipcCmd)
	rootCmd.AddCommand(walletCmd)
}

func initFlags() {
	rootCmd.PersistentFlags().StringVar(&nodeRPCAddress, "node-address", "127.0.0.1:5000", "The address of the node RPC")
	rootCmd.PersistentFlags().IntVar(&maxConcurrency, "maxConcurrency", 10, "Maximum concurrency level")
	rootCmd.PersistentFlags().Int64Var(&blockPartSize, "blockPartSize", (memory.KiB * 128).ToInt64(), "Size of each block part")
	rootCmd.PersistentFlags().BoolVar(&useConnectionPool, "useConnectionPool", true, "Use connection pool")
	rootCmd.PersistentFlags().BoolVar(&enableSDKMonkitStats, "print-stats", false, "Enable printing SDK monkit stats on shutdown")
	ipcCmd.PersistentFlags().StringVar(&accountName, "account", "", "Optional: Wallet name to use for IPC operations. If not provided, will use the first available wallet")
	ipcCmd.PersistentFlags().StringVar(&privateKey, "private-key", "", "Private key for signing IPC transactions")
	streamingFileDownloadCmd.Flags().BoolVar(&filecoinFlag, "filecoin", false, "downloads data from filecoin if they are already sealed there")
	ipcBucketListCmd.Flags().Int64Var(&bucketListOffset, "offset", 0, "offset for bucket list pagination")
	ipcBucketListCmd.Flags().Int64Var(&bucketListLimit, "limit", 20, "limit for bucket list pagination")
	ipcFileListCmd.Flags().Int64Var(&fileListOffset, "offset", 0, "offset for file list pagination")
	ipcFileListCmd.Flags().Int64Var(&fileListLimit, "limit", 20, "limit for file list pagination")

	for _, cmd := range []*cobra.Command{ipcFileUploadCmd, ipcFileDownloadCmd, streamingFileUploadCmd, streamingFileDownloadCmd} {
		cmd.Flags().StringVarP(&encryptionKey, "encryption-key", "e", "", "Encryption key for encrypting file data")
		cmd.Flags().BoolVar(&disableErasureCoding, "disable-erasure-coding", false, "Do not use erasure coding")
	}
}

func initTracing(log *zap.Logger) (*mJaeger.ThriftCollector, func()) {
	collector, err := mJaeger.NewThriftCollector(log, tracingAgentAddr, serviceName, []mJaeger.Tag{}, 0, 0, 0)
	if err != nil {
		log.Error("failed to create collector", zap.Error(err))
	}
	unreg := mJaeger.RegisterJaeger(monkit.Default, collector, mJaeger.Options{
		Fraction: 1,
	})

	return collector, unreg
}

func main() {
	initFlags()
	environment.Register(monkit.Default)
	log.SetOutput(os.Stderr)

	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = logger.Sync() }()

	ctx, cancel := context.WithCancel(context.Background())
	var eg errgroup.Group

	collector, unreg := initTracing(logger)
	eg.Go(func() error {
		collector.Run(ctx)
		return nil
	})
	defer func() {
		cancel()
		unreg()

		err := eg.Wait()
		if err != nil {
			rootCmd.PrintErrf("unexpected errgroup wait error: %v", err)
		}
	}()

	rootCmd.DisableFlagParsing = true
	// traverse early to get subcommand.
	cmd, _, err := rootCmd.Traverse(os.Args[1:])
	if err != nil {
		rootCmd.PrintErrf("Error: %v\n", err)
		_ = rootCmd.Usage()
		return
	}

	rootCmd.DisableFlagParsing = false
	// parse flags early to display usage on error.
	err = cmd.ParseFlags(os.Args[1:])
	if err != nil {
		rootCmd.PrintErrf("Error: failed to parse flags: %v\n", err)
		_ = cmd.Usage()
		return
	}

	if err = rootCmd.Execute(); err != nil {
		rootCmd.PrintErrf("Error: %v\n", err)

		paramErr := &CmdParamsError{}
		if errors.As(err, &paramErr) {
			_ = cmd.Usage()
		}
	}

	if enableSDKMonkitStats {
		printMonkitStats(sdk.GetMonkitStats(), os.Stderr)
	}
}

func cmdCreateBucket(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	sdk, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := sdk.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	result, err := sdk.CreateBucket(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to create bucket: %w", err)
	}

	cmd.PrintErrf("Bucket created: ID=%s, CreatedAt=%s\n", result.Name, result.CreatedAt)

	return nil
}

func cmdDeleteBucket(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	sdk, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := sdk.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	err = sdk.DeleteBucket(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to delete bucket: %w", err)
	}

	cmd.PrintErrf("Bucket deleted: Name=%s\n", bucketName)

	return nil
}

func cmdViewBucket(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	sdk, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := sdk.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	result, err := sdk.ViewBucket(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to get bucket: %w", err)
	}

	cmd.PrintErrf("Bucket: Name=%s, CreatedAt=%s\n", result.Name, result.CreatedAt)

	return nil
}

func cmdListBuckets(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	sdk, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := sdk.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	buckets, err := sdk.ListBuckets(ctx)
	if err != nil {
		return fmt.Errorf("failed to list buckets: %w", err)
	}

	if len(buckets) == 0 {
		cmd.PrintErrln("No buckets")
		return nil
	}
	for _, bucket := range buckets {
		cmd.PrintErrf("Bucket: Name=%s, CreatedAt=%s\n", bucket.Name, bucket.CreatedAt)
	}

	return nil
}

func encryptionKeyBytes() ([]byte, error) {
	decodedKey, err := hex.DecodeString(encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode encryption key: %w", err)
	}
	return decodedKey, nil
}

func parityBlocks() int {
	if !disableErasureCoding {
		return 16
	}
	return 0
}

func printMonkitStats(stats []sdk.MonkitStats, writer io.Writer) {
	_, _ = fmt.Fprintf(writer, "\n=== SDK Monkit Task Statistics ===\n")
	for _, stat := range stats {
		_, _ = fmt.Fprintf(writer, "Function: %s\n", stat.Name)
		_, _ = fmt.Fprintf(writer, "  Success calls: %d\n", stat.Successes)
		_, _ = fmt.Fprintf(writer, "  Highwater (max concurrent): %d\n", stat.Highwater)

		if len(stat.Errors) > 0 {
			_, _ = fmt.Fprintf(writer, "  Errors:\n")
			for errType, count := range stat.Errors {
				_, _ = fmt.Fprintf(writer, "    %s: %d\n", errType, count)
			}
		}

		// Get timing information
		if stat.SuccessTimes != nil && stat.SuccessTimes.Count > 0 {
			_, _ = fmt.Fprintf(writer, "  Success timing (avg): %.2fms\n", float64(stat.SuccessTimes.FullAverage().Milliseconds()))
			_, _ = fmt.Fprintf(writer, "  Success timing (low/high): %.2fms / %.2fms\n",
				float64(stat.SuccessTimes.Low.Milliseconds()), float64(stat.SuccessTimes.High.Milliseconds()))
		}

		if stat.FailureTimes != nil && stat.FailureTimes.Count > 0 {
			_, _ = fmt.Fprintf(writer, "  Failure timing (avg): %.2fms\n", float64(stat.FailureTimes.FullAverage().Milliseconds()))
			_, _ = fmt.Fprintf(writer, "  Failure timing (low/high): %.2fms / %.2fms\n",
				float64(stat.FailureTimes.Low.Milliseconds()), float64(stat.FailureTimes.High.Milliseconds()))
		}

		_, _ = fmt.Fprintf(writer, "\n")
	}
	_, _ = fmt.Fprintf(writer, "=== End SDK Monkit Statistics ===\n\n")
}
