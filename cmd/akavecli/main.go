// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package main provides a command-line interface for managing Akave SDK resources
// such as buckets and files.
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

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

	nodeRPCAddress        string
	privateKey            string
	encryptionKey         string
	maxConcurrency        int
	blockPartSize         int64
	useConnectionPool     bool
	disableErasureCoding  bool
	enableSDKMonkitStats  bool
	accountName           string
	useMetadataEncryption bool

	// tracing.
	mon = monkit.Package()

	tracingAgentAddr = "localhost:6831"
	serviceName      = "akavecli"

	bucketListOffset int64
	bucketListLimit  int64
	fileListOffset   int64
	fileListLimit    int64

	archivalMetadataVerbose bool
	archivalDownload        bool
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
	rootCmd.AddCommand(bucketCmd)
	rootCmd.AddCommand(fileCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(walletCmd)

	// Initialize file and buckets commands
	initStorageCommands()

	// Initialize wallet commands
	initWalletCommands()
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func initFlags() {
	rootCmd.PersistentFlags().StringVar(&nodeRPCAddress, "node-address", "127.0.0.1:5000", "The address of the node RPC")
	rootCmd.PersistentFlags().IntVar(&maxConcurrency, "maxConcurrency", 10, "Maximum concurrency level")
	rootCmd.PersistentFlags().Int64Var(&blockPartSize, "blockPartSize", (memory.KiB * 128).ToInt64(), "Size of each block part")
	rootCmd.PersistentFlags().BoolVar(&useConnectionPool, "useConnectionPool", true, "Use connection pool")
	rootCmd.PersistentFlags().BoolVar(&enableSDKMonkitStats, "print-stats", false, "Enable printing SDK monkit stats on shutdown")
	rootCmd.PersistentFlags().StringVar(&accountName, "account", "", "Optional: Wallet name to use. If not provided, will use the first available wallet")
	rootCmd.PersistentFlags().StringVar(&privateKey, "private-key", "", "Private key for signing transactions")
	rootCmd.PersistentFlags().StringVarP(&encryptionKey, "encryption-key", "e", "", "Encryption key for encrypting file data")
	rootCmd.PersistentFlags().BoolVar(&useMetadataEncryption, "metadata-encryption", false, "enable metadata encryption")
	bucketListCmd.Flags().Int64Var(&bucketListOffset, "offset", 0, "offset for bucket list pagination")
	bucketListCmd.Flags().Int64Var(&bucketListLimit, "limit", 20, "limit for bucket list pagination")
	fileListCmd.Flags().Int64Var(&fileListOffset, "offset", 0, "offset for file list pagination")
	fileListCmd.Flags().Int64Var(&fileListLimit, "limit", 20, "limit for file list pagination")
	fileDownloadCmd.Flags().BoolVar(&archivalDownload, "archival", false, "download from archival storage")
	archivalMetadataCmd.Flags().BoolVarP(&archivalMetadataVerbose, "verbose", "v", false, "Print detailed metadata including chunks and blocks")

	for _, cmd := range []*cobra.Command{fileUploadCmd, fileDownloadCmd} {
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

func run() error {
	initFlags()
	environment.Register(monkit.Default)
	log.SetOutput(os.Stderr)

	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	logger, err := zap.NewProduction()
	if err != nil {
		rootCmd.PrintErrf("failed to create logger: %v\n", err)
		return err
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
		return err
	}

	rootCmd.DisableFlagParsing = false
	// parse flags early to display usage on error.
	err = cmd.ParseFlags(os.Args[1:])
	if err != nil {
		rootCmd.PrintErrf("Error: failed to parse flags: %v\n", err)
		_ = cmd.Usage()
		return err
	}

	if err = rootCmd.Execute(); err != nil {
		rootCmd.PrintErrf("Error: %v\n", err)

		paramErr := &CmdParamsError{}
		if errors.As(err, &paramErr) {
			_ = cmd.Usage()
		}

		return err
	}

	if enableSDKMonkitStats {
		printMonkitStats(sdk.GetMonkitStats(), os.Stderr)
	}

	return nil
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
