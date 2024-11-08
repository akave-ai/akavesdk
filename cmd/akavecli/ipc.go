// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"akave.ai/akavesdk/sdk"
)

var (
	ipcCmd = &cobra.Command{
		Use:   "ipc",
		Short: "Manage files and buckets using IPC API",
	}

	ipcBucketCmd = &cobra.Command{
		Use:   "bucket",
		Short: "Manage buckets ipc",
	}

	ipcFileCmd = &cobra.Command{
		Use:   "file",
		Short: "Manage files in buckets ipc",
	}

	ipcBucketCreateCmd = &cobra.Command{
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
		RunE: cmdCreateBucketIPC,
	}

	ipcBucketDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Removes a bucket",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return NewCmdParamsError(fmt.Sprintf("create bucket command expects exactly 1 argument [bucket name]; got %d", len(args)))
			}

			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}

			return nil
		},
		RunE: cmdDeleteBucketIPC,
	}

	ipcBucketViewCmd = &cobra.Command{
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
		RunE: cmdViewBucketIPC,
	}

	ipcBucketListCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all buckets",
		Args:  cobra.NoArgs,
		RunE:  cmdListBucketsIPC,
	}

	ipcFileListCmd = &cobra.Command{
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
		RunE: cmdListFilesIPC,
	}

	ipcFileInfoCmd = &cobra.Command{
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
		RunE: cmdFileInfoIPC,
	}

	ipcFileUploadCmd = &cobra.Command{
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
		RunE: cmdFileUploadIPC,
	}

	ipcFileDownloadCmd = &cobra.Command{
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
		RunE: cmdFileDownloadIPC,
	}
)

func cmdCreateBucketIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privateKey))
	if err != nil {
		return err
	}
	defer func() {
		if cerr := akaveSDK.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	result, err := ipc.CreateBucket(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to create bucket: %w", err)
	}

	cmd.PrintErrf("Bucket created: ID=%s, CreatedAt=%s\n", result.ID, result.CreatedAt)

	return nil
}

func cmdDeleteBucketIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privateKey))
	if err != nil {
		return err
	}
	defer func() {
		if cerr := akaveSDK.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	if err := ipc.DeleteBucket(ctx, bucketName); err != nil {
		return fmt.Errorf("failed to delete bucket: %w", err)
	}

	cmd.PrintErrf("Bucket deleted: Name=%s\n", bucketName)

	return nil
}

func cmdViewBucketIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privateKey))
	if err != nil {
		return err
	}
	defer func() {
		if cerr := akaveSDK.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	result, err := ipc.ViewBucket(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to get bucket: %w", err)
	}

	cmd.PrintErrf("Bucket: ID=%s, Name=%s, CreatedAt=%s\n", result.ID, result.Name, result.CreatedAt)

	return nil
}

func cmdListBucketsIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privateKey))
	if err != nil {
		return err
	}
	defer func() {
		if cerr := akaveSDK.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	buckets, err := ipc.ListBuckets(ctx)
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

func cmdListFilesIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privateKey))
	if err != nil {
		return err
	}
	defer func() {
		if cerr := akaveSDK.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	files, err := ipc.ListFiles(ctx, bucketName)
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

func cmdFileInfoIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privateKey))
	if err != nil {
		return err
	}
	defer func() {
		if cerr := akaveSDK.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	file, err := ipc.FileInfo(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	cmd.PrintErrf("File: Name=%s, RootCID=%s, Size=%d, CreatedAt=%s\n", file.Name, file.RootCID, file.Size, file.CreatedAt)

	return nil
}

func cmdFileUploadIPC(cmd *cobra.Command, args []string) (err error) {
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

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privateKey))
	if err != nil {
		return err
	}
	defer func() {
		if cerr := akaveSDK.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	fileUpload, err := ipc.CreateFileUpload(ctx, bucketName, fileName, fileInfo.Size(), file)
	if err != nil {
		return fmt.Errorf("failed to create file upload: %w", err)
	}

	if err := ipc.Upload(ctx, fileUpload); err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	cmd.PrintErrf("File uploaded successfully: Name=%s, RootCID=%s\n", fileName, fileUpload.RootCID)

	return nil
}

func cmdFileDownloadIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]
	destPath := args[2]

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privateKey))
	if err != nil {
		return err
	}
	defer func() {
		if cerr := akaveSDK.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	fileDownload, err := ipc.CreateFileDownload(ctx, bucketName, fileName)
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

	if err := ipc.Download(ctx, fileDownload, outFile); err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	cmd.PrintErrf("File downloaded successfully: Name=%s, Path=%s\n", fileName, filepath.Join(destPath, fileName))
	return nil
}
