// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package main provides a command-line interface for managing Akave SDK resources
// such as buckets and files.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spacemonkeygo/monkit/v3"
	"github.com/spacemonkeygo/monkit/v3/environment"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	mJaeger "storj.io/monkit-jaeger"

	"akave.ai/akavesdk/private/memory"
	"akave.ai/akavesdk/private/version"
	"akave.ai/akavesdk/sdk"
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

	fileCmd = &cobra.Command{
		Use:   "file",
		Short: "Manage files in buckets",
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

	fileListCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all files in a bucket",
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
		RunE: cmdListFiles,
	}

	fileInfoCmd = &cobra.Command{
		Use:   "info",
		Short: "Retrieves file information",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}

			if len(args) != 2 {
				return NewCmdParamsError(fmt.Sprintf("file info command expects exactly 2 arguments [bucket name] [file name]; got %d", len(args)))
			}

			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}

			if args[1] == "" {
				return NewCmdParamsError("file name is required")
			}

			return nil
		},
		RunE: cmdFileInfo,
	}

	fileUploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "Uploads a file to a bucket",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}

			if len(args) != 2 {
				return NewCmdParamsError(fmt.Sprintf("file upload command expects exactly 2 arguments [bucket name] [file path]; got %d", len(args)))
			}

			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}

			if args[1] == "" {
				return NewCmdParamsError("file path is required")
			}

			return nil
		},
		RunE: cmdFileUpload,
	}

	fileDownloadCmd = &cobra.Command{
		Use:   "download",
		Short: "Downloads a file from a bucket",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}

			if len(args) != 3 {
				return NewCmdParamsError(fmt.Sprintf("file download command expects exactly 3 arguments [bucket name] [file name] [destination path]; got %d", len(args)))
			}

			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}

			if args[1] == "" {
				return NewCmdParamsError("file name is required")
			}

			if args[2] == "" {
				return NewCmdParamsError("destination path is required")
			}

			return nil
		},
		RunE: cmdFileDownload,
	}

	fileDownloadV2Cmd = &cobra.Command{
		Use:   "download2",
		Short: "Downloads a file from a bucket",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}

			if len(args) != 3 {
				return NewCmdParamsError(fmt.Sprintf("file download command expects exactly 3 arguments [bucket name] [file name] [destination path]; got %d", len(args)))
			}

			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}

			if args[1] == "" {
				return NewCmdParamsError("file name is required")
			}

			if args[2] == "" {
				return NewCmdParamsError("destination path is required")
			}

			return nil
		},

		RunE: cmdFileDownloadV2,
	}

	fileRangeDownloadCmd = &cobra.Command{
		Use:     "download-range",
		Short:   "Downloads a file from a bucket with given range",
		Example: "akavecli file download-range foo-bucket test.txt 10-90 .",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}

			if len(args) != 4 {
				return NewCmdParamsError(fmt.Sprintf("file download command expects exactly 4 arguments [bucket name] [file name] [range] [destination path]; got %d", len(args)))
			}

			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}

			if args[1] == "" {
				return NewCmdParamsError("file name is required")
			}

			if len(strings.Split(args[2], "-")) != 2 {
				return NewCmdParamsError("range should be in the format start-end")
			}

			if args[3] == "" {
				return NewCmdParamsError("destination path is required")
			}

			return nil
		},
		RunE: cmdFileRangeDownload,
	}

	nodeRPCAddress    string
	privateKey        string
	maxConcurrency    int
	blockPartSize     int64
	useConnectionPool bool

	// tracing.
	mon = monkit.Package()

	tracingAgentAddr = "localhost:6831"
	collectorAddr    = "http://localhost:14268/api/traces"
	serviceName      = "akavecli"
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

	fileCmd.AddCommand(fileListCmd)
	fileCmd.AddCommand(fileInfoCmd)
	fileCmd.AddCommand(fileUploadCmd)
	fileCmd.AddCommand(fileDownloadCmd)
	fileCmd.AddCommand(fileDownloadV2Cmd)
	fileCmd.AddCommand(fileRangeDownloadCmd)
	// streaming file API
	fileStreamingCmd.AddCommand(streamingFileListCmd)
	fileStreamingCmd.AddCommand(streamingFileInfoCmd)
	fileStreamingCmd.AddCommand(streamingFileUploadCmd)
	fileStreamingCmd.AddCommand(streamingFileDownloadCmd)
	fileStreamingCmd.AddCommand(streamingFileDeleteCmd)
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

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(bucketCmd)
	rootCmd.AddCommand(fileCmd)
	rootCmd.AddCommand(fileStreamingCmd)
	rootCmd.AddCommand(ipcCmd)
}

func initFlags() {
	rootCmd.PersistentFlags().StringVar(&nodeRPCAddress, "node-address", "127.0.0.1:5000", "The address of the node RPC")
	rootCmd.PersistentFlags().IntVar(&maxConcurrency, "maxConcurrency", 10, "Maximum concurrency level")
	rootCmd.PersistentFlags().Int64Var(&blockPartSize, "blockPartSize", (memory.KiB * 128).ToInt64(), "Size of each block part")
	rootCmd.PersistentFlags().BoolVar(&useConnectionPool, "useConnectionPool", true, "Use connection pool")
	ipcCmd.PersistentFlags().StringVar(&privateKey, "private-key", "", "Private key for signing IPC transactions")
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

	cmd.PrintErrf("Bucket created: ID=%s, CreatedAt=%s\n", result.ID, result.CreatedAt)

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

	cmd.PrintErrf("Bucket: ID=%s, Name=%s, CreatedAt=%s\n", result.ID, result.Name, result.CreatedAt)

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
		cmd.PrintErrf("Bucket: ID=%s, Name=%s, CreatedAt=%s\n", bucket.ID, bucket.Name, bucket.CreatedAt)
	}

	return nil
}

func cmdListFiles(cmd *cobra.Command, args []string) (err error) {
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

	files, err := sdk.ListFiles(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}

	if len(files) == 0 {
		cmd.PrintErrln("No files")
		return nil
	}
	for _, file := range files {
		cmd.PrintErrf("File: Name=%s, RootCID=%s, Size=%d, CreatedAt=%s\n", file.Name, file.RootCID, file.Size, file.CreatedAt)
	}

	return nil
}

func cmdFileInfo(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]

	sdk, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := sdk.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	file, err := sdk.FileInfo(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	cmd.PrintErrf("File: Name=%s, RootCID=%s, Size=%d, CreatedAt=%s\n", file.Name, file.RootCID, file.Size, file.CreatedAt)

	return nil
}

func cmdFileUpload(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	filePath := args[1]
	fileName := filepath.Base(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			cmd.PrintErrf("failed to close file: %v", cerr)
		}
	}()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	sdk, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := sdk.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	fileUpload, err := sdk.CreateFileUpload(ctx, bucketName, fileName, fileInfo.Size(), file)
	if err != nil {
		return fmt.Errorf("failed to create file upload: %w", err)
	}

	if err := sdk.Upload(ctx, fileUpload); err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	cmd.PrintErrf("File uploaded successfully: Name=%s, RootCID=%s\n", fileName, fileUpload.RootCID)

	return nil
}

func cmdFileDownload(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]
	destPath := args[2]

	sdk, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := sdk.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	fileDownload, err := sdk.CreateFileDownload(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to create file download: %w", err)
	}

	outFile, err := os.Create(filepath.Join(destPath, fileName))
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer func() {
		if cerr := outFile.Close(); cerr != nil && err == nil {
			cmd.PrintErrf("failed to close destination file: %v", cerr)
		}
	}()

	if err := sdk.Download(ctx, fileDownload, outFile); err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	cmd.PrintErrf("File downloaded successfully: Name=%s, Path=%s\n", fileName, filepath.Join(destPath, fileName))
	return nil
}

func cmdFileDownloadV2(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]
	destPath := args[2]

	sdk, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := sdk.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	fileDownload, err := sdk.CreateFileDownloadV2(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to create file download: %w", err)
	}

	outFile, err := os.Create(filepath.Join(destPath, fileName))
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer func() {
		if cerr := outFile.Close(); cerr != nil && err == nil {
			cmd.PrintErrf("failed to close destination file: %v", cerr)
		}
	}()

	if err := sdk.DownloadV2(ctx, fileDownload, outFile); err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	cmd.PrintErrf("File downloaded successfully: Name=%s, Path=%s\n", fileName, filepath.Join(destPath, fileName))
	return nil
}

func cmdFileRangeDownload(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]
	blockRangeStr := args[2]
	destPath := args[3]

	var start, end int64
	n, err := fmt.Sscanf(blockRangeStr, "%d-%d", &start, &end)
	if n != 2 || err != nil {
		return fmt.Errorf("failed to parse ranges in %s: %w", blockRangeStr, err)
	}

	sdk, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := sdk.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	fileDownload, err := sdk.CreateRangeFileDownload(ctx, bucketName, fileName, start, end)
	if err != nil {
		return fmt.Errorf("failed to create range file download: %w", err)
	}

	outFile, err := os.Create(filepath.Join(destPath, fileName))
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer func() {
		if cerr := outFile.Close(); cerr != nil && err == nil {
			cmd.PrintErrf("failed to close destination file: %v", cerr)
		}
	}()

	if err := sdk.Download(ctx, fileDownload, outFile); err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	cmd.PrintErrf("File part downloaded successfully: Name=%s, Blocks=%s, Path=%s\n", fileName, blockRangeStr, filepath.Join(destPath, fileName))
	return nil
}
