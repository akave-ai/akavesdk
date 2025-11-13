// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/akave-ai/akavesdk/sdk"
)

var (
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
				return NewCmdParamsError(fmt.Sprintf("create bucket command expects exactly 1 argument [bucket name]; got %d", len(args)))
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

	fileDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Removes a file from a bucket",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}

			if len(args) != 2 {
				return NewCmdParamsError(fmt.Sprintf("file delete command expects exactly 2 arguments [bucket name] [file name]; got %d", len(args)))
			}

			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}

			if args[1] == "" {
				return NewCmdParamsError("file name is required")
			}

			return nil
		},
		RunE: cmdFileDelete,
	}

	archivalMetadataCmd = &cobra.Command{
		Use:   "archival-metadata",
		Short: "Retrieves archival metadata for a file",
		Args: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				args[i] = strings.TrimSpace(arg)
			}

			if len(args) != 2 {
				return NewCmdParamsError(fmt.Sprintf("archival-metadata command expects exactly 2 arguments [bucket name] [file name]; got %d", len(args)))
			}

			if args[0] == "" {
				return NewCmdParamsError("bucket name is required")
			}

			if args[1] == "" {
				return NewCmdParamsError("file name is required")
			}

			return nil
		},
		RunE: cmdArchivalMetadata,
	}
)

func initStorageCommands() {
	bucketCmd.AddCommand(bucketCreateCmd)
	bucketCmd.AddCommand(bucketViewCmd)
	bucketCmd.AddCommand(bucketListCmd)
	bucketCmd.AddCommand(bucketDeleteCmd)
	fileCmd.AddCommand(fileUploadCmd)
	fileCmd.AddCommand(fileDownloadCmd)
	fileCmd.AddCommand(fileListCmd)
	fileCmd.AddCommand(fileInfoCmd)
	fileCmd.AddCommand(fileDeleteCmd)
	fileCmd.AddCommand(archivalMetadataCmd)
}

func cmdCreateBucket(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	sdkOptions, err := defaultSDKOptions(cmd)
	if err != nil {
		return err
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

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	result, err := ipc.CreateBucket(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to create bucket: %w", err)
	}

	cmd.PrintErrf("Bucket created: ID=%s, Name=%s, CreatedAt=%s\n", result.ID, result.Name, result.CreatedAt)

	return nil
}

func cmdDeleteBucket(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	sdkOptions, err := defaultSDKOptions(cmd)
	if err != nil {
		return err
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

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	err = ipc.DeleteBucket(ctx, bucketName)
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

	sdkOptions, err := defaultSDKOptions(cmd)
	if err != nil {
		return err
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

func cmdListBuckets(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	sdkOptions, err := defaultSDKOptions(cmd)
	if err != nil {
		return err
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

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	buckets, err := ipc.ListBuckets(ctx, bucketListOffset, bucketListLimit)
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

func cmdListFiles(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	sdkOptions, err := defaultSDKOptions(cmd)
	if err != nil {
		return err
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

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	files, err := ipc.ListFiles(ctx, bucketName, fileListOffset, fileListLimit)
	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}

	if len(files) == 0 {
		cmd.PrintErrln("No files")
		return nil
	}
	for _, file := range files {
		cmd.PrintErrf("File: Name=%s, RootCID=%s, EncodedSize=%d, ActualSize=%d, CreatedAt=%s\n", file.Name, file.RootCID, file.EncodedSize, file.ActualSize, file.CreatedAt)
	}

	return nil
}

func cmdFileInfo(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]

	sdkOptions, err := defaultSDKOptions(cmd)
	if err != nil {
		return err
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

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	file, err := ipc.FileInfo(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	cmd.PrintErrf("File: Name=%s, RootCID=%s, EncodedSize=%d, ActualSize=%d CreatedAt=%s\n", file.Name, file.RootCID, file.EncodedSize, file.ActualSize, file.CreatedAt)

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

	fi, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	sdkOpts, err := defaultSDKOptions(cmd)
	if err != nil {
		return err
	}
	sdkOpts = append(sdkOpts, sdk.WithErasureCoding(parityBlocks()))

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdkOpts...)
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

	fileUpload, err := ipc.CreateFileUpload(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to create file upload: %w", err)
	}

	var size int64
	if disableErasureCoding {
		size = fi.Size()
	} else {
		size = fi.Size() * 2 // TODO: not sure that it is correct to just multiply original size by 2
	}

	bar, done := trackProgress(cmd, size, fileUpload, "Uploading")

	fileMeta, err := ipc.Upload(ctx, fileUpload, file)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	close(done)
	if err := bar.Finish(); err != nil {
		return fmt.Errorf("failed to finish progress bar: %w", err)
	}
	if err := bar.Clear(); err != nil {
		return fmt.Errorf("failed to clear progress bar: %w", err)
	}

	cmd.PrintErrf("File uploaded successfully: Name=%s, RootCID=%s, ActualSize=%d, EncodedSize=%d\n", fileName, fileMeta.RootCID, fileMeta.Size, fileMeta.EncodedSize)

	return nil
}

func cmdFileDownload(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]
	destPath := args[2]

	sdkOpts, err := defaultSDKOptions(cmd)
	if err != nil {
		return err
	}
	sdkOpts = append(sdkOpts, sdk.WithErasureCoding(parityBlocks()))

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdkOpts...)
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

	info, err := ipc.FileInfo(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	bar, done := trackProgress(cmd, info.EncodedSize, fileDownload, "Downloading")

	// TODO: add test for archival mode
	if archivalDownload {
		if err := ipc.DownloadArchival(ctx, fileDownload, outFile); err != nil {
			return fmt.Errorf("failed to download file from pdp: %w", err)
		}
	} else {
		if err := ipc.Download(ctx, fileDownload, outFile); err != nil {
			return fmt.Errorf("failed to download file: %w", err)
		}
	}

	close(done)
	if err := bar.Finish(); err != nil {
		return fmt.Errorf("failed to finish progress bar: %w", err)
	}
	if err := bar.Clear(); err != nil {
		return fmt.Errorf("failed to clear progress bar: %w", err)
	}

	chunks, blocks, finalBytes := fileDownload.Stats()
	cmd.PrintErrf("File downloaded successfully: Name=%s, Path=%s, Size=%d, Chunks=%d, Blocks=%d\n",
		info.Name, filepath.Join(destPath, fileName), finalBytes, chunks, blocks)

	return nil
}

func cmdFileDelete(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]

	sdkOpts, err := defaultSDKOptions(cmd)
	if err != nil {
		return err
	}
	sdkOpts = append(sdkOpts, sdk.WithErasureCoding(parityBlocks()))

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdkOpts...)
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

	if err := ipc.FileDelete(ctx, bucketName, fileName); err != nil {
		return err
	}

	cmd.PrintErrf("File successfully deleted: Name=%s\n", fileName)
	return nil
}

func cmdArchivalMetadata(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	bucketName, fileName := args[0], args[1]

	sdkOptions, err := defaultSDKOptions(cmd)
	if err != nil {
		return err
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

	ipc, err := akaveSDK.IPC()
	if err != nil {
		return err
	}

	metadata, err := ipc.ArchivalMetadata(ctx, bucketName, fileName)
	if err != nil {
		return fmt.Errorf("failed to get archival metadata: %w", err)
	}

	if archivalMetadataVerbose {
		printVerboseArchivalMetadata(cmd, metadata)
	} else {
		printSimpleArchivalMetadata(cmd, metadata)
	}

	return nil
}

func printSimpleArchivalMetadata(cmd *cobra.Command, metadata sdk.ArchivalMetadata) {
	cmd.PrintErrf("Bucket: %s, File: %s\n", metadata.BucketName, metadata.Name)

	allBlocksHavePDP := true

outer:
	for _, chunk := range metadata.Chunks {
		for _, block := range chunk.Blocks {
			if block.PDPData == nil {
				allBlocksHavePDP = false
				break outer
			}
		}
	}

	if allBlocksHavePDP {
		cmd.PrintErrln("Status: Available for download from archival storage (all blocks have PDP data)")
	} else {
		cmd.PrintErrln("Status: Not fully available in archival storage (some blocks are missing PDP data)")
	}
}

func printVerboseArchivalMetadata(cmd *cobra.Command, metadata sdk.ArchivalMetadata) {
	cmd.PrintErrf("Bucket: %s, File: %s\n", metadata.BucketName, metadata.Name)
	cmd.PrintErrf("Total Chunks: %d\n\n", len(metadata.Chunks))

	for _, chunk := range metadata.Chunks {
		cmd.PrintErrf("  Chunk CID: %s, Size: %d\n", chunk.CID, chunk.Size)

		for _, block := range chunk.Blocks {
			if block.PDPData != nil {
				cmd.PrintErrf("    Block CID: %s, URL: %s, Offset: %d, Size: %d, Dataset ID: %d\n",
					block.CID, block.PDPData.URL, block.PDPData.Offset, block.PDPData.Size, block.PDPData.DataSetID)
			} else {
				cmd.PrintErrf("    Block CID: %s, Size: %d (No PDP data)\n", block.CID, block.Size)
			}
		}
		cmd.PrintErrln()
	}
}

// getWalletPrivateKey returns the private key either from the flag or from a wallet.
// It also prints the wallet address and name if a wallet was used.
func getWalletPrivateKey(cmd *cobra.Command) (privKey string, err error) {
	if privateKey != "" {
		return privateKey, nil
	}

	var walletAddress string
	privKey, walletAddress, name, err := PrivateKeyFromWallet(accountName)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	cmd.PrintErrln(fmt.Sprintf("Using wallet account: %s (%s)", name, walletAddress))
	return privKey, err
}

// defaultSDKOptions constructs the default SDK options based on CLI flags.
func defaultSDKOptions(cmd *cobra.Command) ([]sdk.Option, error) {
	privateKey, err := getWalletPrivateKey(cmd)
	if err != nil {
		return nil, err
	}

	var encryptionKeyBytes []byte
	if encryptionKey != "" {
		decodedKey, err := hex.DecodeString(encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("failed to decode encryption key: %w", err)
		}
		encryptionKeyBytes = decodedKey
	}

	opts := []sdk.Option{
		sdk.WithPrivateKey(privateKey),
		sdk.WithEncryptionKey(encryptionKeyBytes),
	}

	if useMetadataEncryption {
		opts = append(opts, sdk.WithMetadataEncryption())
	}

	return opts, nil
}
