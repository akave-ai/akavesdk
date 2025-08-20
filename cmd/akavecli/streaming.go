// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"

	"github.com/akave-ai/akavesdk/sdk"
)

var (
	fileStreamingCmd = &cobra.Command{
		Use:   "files-streaming",
		Short: "Manage files in buckets using streaming API",
	}

	streamingFileListCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all files in a bucket",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}
			if len(args) != 1 {
				return NewCmdParamsError(fmt.Sprintf("file list command expects exactly 1 argument [bucket name]; got %d", len(args)))
			}
			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}

			return nil
		},
		RunE: cmdStreamingListFiles,
	}

	streamingFileVersionsCmd = &cobra.Command{
		Use:   "versions",
		Short: "Lists all versions of a given file",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}
			if len(args) != 2 {
				return NewCmdParamsError(fmt.Sprintf("file versions command expects exactly 2 argument [bucket name] [file name]; got %d", len(args)))
			}
			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}
			if args[1] == "" {
				return NewCmdParamsError("file name is required")
			}

			return nil
		},
		RunE: cmdStreamingFileVersions,
	}

	streamingFileInfoCmd = &cobra.Command{
		Use:   "info",
		Short: "Retrieves file information",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}
			if len(args) != 2 {
				return NewCmdParamsError(fmt.Sprintf("view bucket command expects exactly 2 arguments [bucket name] [file name]; got %d", len(args)))
			}
			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}
			if args[1] == "" {
				return NewCmdParamsError("file name is required")
			}

			return nil
		},
		RunE: cmdStreamingFileInfo,
	}

	streamingFileUploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "Uploads a file to a bucket",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}
			if len(args) != 2 {
				return NewCmdParamsError(fmt.Sprintf("view bucket command expects exactly 2 arguments [bucket name] [file path]; got %d", len(args)))
			}
			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}
			if args[1] == "" {
				return NewCmdParamsError("file path is required")
			}

			return nil
		},
		RunE: cmdStreamingFileUpload,
	}

	streamingFileDownloadCmd = &cobra.Command{
		Use:   "download",
		Short: "Downloads a file from a bucket",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}
			if len(args) != 3 {
				return NewCmdParamsError(fmt.Sprintf("view bucket command expects 3 or 4 arguments [bucket name] [file name] [destination path] [root_cid]; got %d", len(args)))
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
		RunE: cmdStreamingFileDownload,
	}

	streamingFileDeleteCmd = &cobra.Command{
		Use:   "delete [bucket name] [file name]",
		Short: "Delete a file from bucket",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}
			if len(args) != 2 {
				return NewCmdParamsError(fmt.Sprintf("delete bucket command expects exactly 2 arguments [bucket name] [file name]; got %d", len(args)))
			}
			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}
			if args[1] == "" {
				return NewCmdParamsError("file name is required")
			}

			return nil
		},
		RunE: cmdStreamingFileDelete,
	}
)

func cmdStreamingListFiles(cmd *cobra.Command, args []string) (err error) {
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

	files, err := sdk.StreamingAPI().ListFiles(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}

	if len(files) == 0 {
		cmd.PrintErrln("No files")
		return nil
	}
	for _, file := range files {
		cmd.PrintErrf("File: Name=%s, RootCID=%s, Size=%d, CreatedAt=%s\n", file.Name, file.RootCID, file.EncodedSize, file.CreatedAt)
	}

	return nil
}

func cmdStreamingFileInfo(cmd *cobra.Command, args []string) (err error) {
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

	file, err := sdk.StreamingAPI().FileInfo(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	cmd.PrintErrf("File: Name=%s, RootCID=%s, Size=%d, CreatedAt=%s\n", file.Name, file.RootCID, file.EncodedSize, file.CreatedAt)
	return nil
}

func cmdStreamingFileUpload(cmd *cobra.Command, args []string) (err error) {
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

	key, err := encryptionKeyBytes()
	if err != nil {
		return err
	}

	sdkOptions := []sdk.Option{
		sdk.WithEncryptionKey(key),
		sdk.WithErasureCoding(parityBlocks()),
	}

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdkOptions...)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := akaveSDK.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	streamingAPI := akaveSDK.StreamingAPI()

	fileUpload, err := streamingAPI.CreateFileUpload(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to create file upload: %w", err)
	}

	fi, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	bar := progressbar.DefaultBytes(
		fi.Size(),
		"uploading",
	)
	r := progressbar.NewReader(file, bar)

	info, err := streamingAPI.Upload(ctx, fileUpload, &r)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	// Finish progress bar: fill the bar as in this point it's guaranteed that the file has been uploaded
	if err := bar.Finish(); err != nil {
		return fmt.Errorf("failed to finish progress bar: %w", err)
	}

	cmd.PrintErrf("File uploaded successfully: Name=%s, RootCID=%s, Size=%d, TransferredSize=%d\n", fileName, info.RootCID, fi.Size(), info.EncodedSize)

	return nil
}

func cmdStreamingFileDownload(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]
	destPath := args[2]
	rootCID := ""
	if len(args) == 4 {
		rootCID = args[3]
	}

	key, err := encryptionKeyBytes()
	if err != nil {
		return err
	}

	sdkOptions := []sdk.Option{
		sdk.WithEncryptionKey(key),
		sdk.WithErasureCoding(parityBlocks()),
	}

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdkOptions...)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := akaveSDK.Close(); cerr != nil {
			cmd.PrintErrf("failed to close SDK: %v", cerr)
		}
	}()

	streamingAPI := akaveSDK.StreamingAPI()

	fileDownload, err := streamingAPI.CreateFileDownload(ctx, bucketName, fileName, rootCID)
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

	info, err := streamingAPI.FileInfo(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}
	// TODO: doesn't display correct bar with file size from Info as it contains some overhead
	bar := progressbar.DefaultBytes(
		info.EncodedSize,
		"downloading",
	)

	downloadFunc := streamingAPI.Download
	if filecoinFlag {
		downloadFunc = streamingAPI.DownloadV2
	}

	if err = downloadFunc(ctx, fileDownload, io.MultiWriter(bar, outFile)); err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	writtenBytes := bar.State().CurrentNum
	// Finish progress bar: fill the bar as in this point it's guaranteed that the file has been downloaded
	if err := bar.Finish(); err != nil {
		return fmt.Errorf("failed to finish progress bar: %w", err)
	}

	cmd.PrintErrf("File downloaded successfully: Name=%s, Path=%s, Size=%d, TransferredSize=%d\n", fileName, filepath.Join(destPath, fileName), writtenBytes, info.EncodedSize)

	return nil
}

func cmdStreamingFileDelete(cmd *cobra.Command, args []string) (err error) {
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

	err = sdk.StreamingAPI().FileDelete(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	cmd.PrintErrf("File deleted successfully: Name=%s\n", fileName)
	return nil
}

func cmdStreamingFileVersions(cmd *cobra.Command, args []string) (err error) {
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

	versions, err := sdk.StreamingAPI().FileVersions(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to get file versions: %w", err)
	}

	for _, version := range versions {
		cmd.PrintErrf("Version: RootCID=%s, Size=%d, CreatedAt=%s\n", version.RootCID, version.EncodedSize, version.CreatedAt)
	}

	return nil
}
