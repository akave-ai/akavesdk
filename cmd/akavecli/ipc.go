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

	ipcFileDeleteCmd = &cobra.Command{
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
		RunE: cmdFileDeleteIPC,
	}
)

func cmdCreateBucketIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	privKey, err := getWalletPrivateKey(cmd)
	if err != nil {
		return err
	}

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privKey))
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

func cmdDeleteBucketIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	privKey, err := getWalletPrivateKey(cmd)
	if err != nil {
		return err
	}

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privKey))
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

	privKey, err := getWalletPrivateKey(cmd)
	if err != nil {
		return err
	}

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privKey))
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

	privKey, err := getWalletPrivateKey(cmd)
	if err != nil {
		return err
	}

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privKey))
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
		cmd.PrintErrf("Bucket: Name=%s, CreatedAt=%s\n", bucket.Name, bucket.CreatedAt)
	}

	return nil
}

func cmdListFilesIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]

	privKey, err := getWalletPrivateKey(cmd)
	if err != nil {
		return err
	}

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privKey))
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
		cmd.PrintErrf("File: Name=%s, RootCID=%s, EncodedSize=%d, CreatedAt=%s\n", file.Name, file.RootCID, file.EncodedSize, file.CreatedAt)
	}

	return nil
}

func cmdFileInfoIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]

	privKey, err := getWalletPrivateKey(cmd)
	if err != nil {
		return err
	}

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool, sdk.WithPrivateKey(privKey))
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

	cmd.PrintErrf("File: Name=%s, RootCID=%s, EncodedSize=%d, CreatedAt=%s\n", file.Name, file.RootCID, file.EncodedSize, file.CreatedAt)

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

	fi, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	privKey, err := getWalletPrivateKey(cmd)
	if err != nil {
		return err
	}

	key, err := encryptionKeyBytes()
	if err != nil {
		return err
	}

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool,
		sdk.WithPrivateKey(privKey),
		sdk.WithEncryptionKey(key),
		sdk.WithErasureCoding(parityBlocks()),
	)
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

	if err := ipc.CreateFileUpload(ctx, bucketName, fileName); err != nil {
		return fmt.Errorf("failed to create file upload: %w", err)
	}

	bar := progressbar.DefaultBytes(
		fi.Size(),
		"uploading",
	)
	r := progressbar.NewReader(file, bar)

	fileMeta, err := ipc.Upload(ctx, bucketName, fileName, &r)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	// Finish progress bar: fill the bar as in this point it's guaranteed that the file has been uploaded
	if err := bar.Finish(); err != nil {
		return fmt.Errorf("failed to finish progress bar: %w", err)
	}

	cmd.PrintErrf("File uploaded successfully: Name=%s, RootCID=%s, EncodedSize=%d\n", fileName, fileMeta.RootCID, fileMeta.EncodedSize)

	return nil
}

func cmdFileDownloadIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]
	destPath := args[2]

	privKey, err := getWalletPrivateKey(cmd)
	if err != nil {
		return err
	}

	key, err := encryptionKeyBytes()
	if err != nil {
		return err
	}

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool,
		sdk.WithPrivateKey(privKey),
		sdk.WithEncryptionKey(key),
		sdk.WithErasureCoding(parityBlocks()),
	)
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
	bar := progressbar.DefaultBytes(
		info.EncodedSize,
		"downloading",
	)

	if err := ipc.Download(ctx, fileDownload, io.MultiWriter(bar, outFile)); err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	writtenBytes := bar.State().CurrentNum
	// Finish progress bar: fill the bar as in this point it's guaranteed that the file has been downloaded
	if err := bar.Finish(); err != nil {
		return fmt.Errorf("failed to finish progress bar: %w", err)
	}

	cmd.PrintErrf("File downloaded successfully: Name=%s, Path=%s, Size=%d\n", fileName, filepath.Join(destPath, fileName), writtenBytes)
	return nil
}

func cmdFileDeleteIPC(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)
	bucketName := args[0]
	fileName := args[1]

	privKey, err := getWalletPrivateKey(cmd)
	if err != nil {
		return err
	}

	key, err := encryptionKeyBytes()
	if err != nil {
		return err
	}

	akaveSDK, err := sdk.New(nodeRPCAddress, maxConcurrency, blockPartSize, useConnectionPool,
		sdk.WithPrivateKey(privKey),
		sdk.WithEncryptionKey(key),
		sdk.WithErasureCoding(parityBlocks()),
	)
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

	cmd.PrintErrf("File successfully deleted: Name=%s", fileName)
	return nil
}

// getWalletPrivateKey returns the private key either from the flag or from a wallet.
// It also returns the wallet address and name if a wallet was used.
func getWalletPrivateKey(cmd *cobra.Command) (privKey string, err error) {
	if privateKey != "" {
		return privateKey, nil
	}

	var walletAddress string
	privKey, walletAddress, err = PrivateKeyFromWallet(accountName)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	name := accountName
	if name == "" {
		// If no account was specified, we used the first available wallet
		// Get the name from the address for display purposes
		entries, _ := os.ReadDir(keystoreDir)
		for _, entry := range entries {
			if !strings.HasSuffix(entry.Name(), walletFileExt) {
				continue
			}
			name = strings.TrimSuffix(entry.Name(), walletFileExt)
			break
		}
	}

	cmd.PrintErrf("Using wallet account: %s (%s)\n", name, walletAddress)
	return privKey, nil
}
