// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var (
	walletCmd = &cobra.Command{
		Use:   "wallet",
		Short: "Manage wallets and accounts",
	}

	walletCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Creates a new wallet",
		Args:  cobra.ExactArgs(1),
		RunE:  cmdCreateWallet,
	}

	walletListCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all wallets",
		Args:  cobra.NoArgs,
		RunE:  cmdListWallets,
	}

	walletGetKeyCmd = &cobra.Command{
		Use:   "export-key",
		Short: "Exports private key for a wallet",
		Args:  cobra.ExactArgs(1),
		RunE:  cmdGetWalletKey,
	}

	walletImportCmd = &cobra.Command{
		Use:   "import",
		Short: "Import a wallet using a private key",
		Args:  cobra.ExactArgs(2),
		RunE:  cmdImportWallet,
	}

	walletFileExt = ".json"
)

func initWalletCommands() {
	walletCmd.AddCommand(walletCreateCmd)
	walletCmd.AddCommand(walletListCmd)
	walletCmd.AddCommand(walletGetKeyCmd)
	walletCmd.AddCommand(walletImportCmd)
}

// getDefaultKeystoreDir returns the default keystore directory
func getDefaultKeystoreDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ".akave_wallets"
	}
	return filepath.Join(homeDir, ".akave_wallets")
}

// getWalletPath returns the path to a wallet file
func getWalletPath(walletName string) (string, error) {
	keystoreDir := getDefaultKeystoreDir()
	if err := os.MkdirAll(keystoreDir, 0700); err != nil {
		return "", fmt.Errorf("failed to create keystore directory: %w", err)
	}
	return filepath.Join(keystoreDir, walletName+walletFileExt), nil
}

func cmdCreateWallet(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	if len(args) != 1 {
		return NewCmdParamsError("create wallet command expects exactly 1 argument [wallet name]")
	}
	
	walletName := args[0]
	walletPath, err := getWalletPath(walletName)
	if err != nil {
		return err
	}

	if _, err := os.Stat(walletPath); err == nil {
		return fmt.Errorf("wallet with name '%s' already exists", walletName)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to check if wallet exists: %w", err)
	}

	//generate private key
	privateKey, err := crypto.GenerateKey()
    if err != nil {
        cmd.PrintErrln("Failed to generate private key:", err)
		return err
    }

    privateKeyBytes := crypto.FromECDSA(privateKey)

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        cmd.PrintErrln("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil
    }

    address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// Create wallet info struct
	walletInfo := struct {
		Address    string `json:"address"`
		PrivateKey string `json:"private_key"`
	}{
		Address:    address,
		PrivateKey: hexutil.Encode(privateKeyBytes)[2:],
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(walletInfo, "", "    ")
	if err != nil {
		cmd.PrintErrln("Failed to marshal wallet info:", err)
		return err
	}

	// Write to file
	if err := os.WriteFile(walletPath, jsonData, 0600); err != nil {
		cmd.PrintErrln("Failed to write wallet file:", err)
		return err
	}

	cmd.Printf("Wallet (%s) created successfully at %s\n Fund your wallet with AKVF at https://faucet.akave.ai/ \n", walletName, walletPath)

	return nil
}

func cmdListWallets(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	// List all wallets	
	keystoreDir := getDefaultKeystoreDir()
	entries, err := os.ReadDir(keystoreDir)
	if err != nil {
		if os.IsNotExist(err) {
			cmd.PrintErrf("No wallets found in %s\n", keystoreDir)
			return nil
		}
		return fmt.Errorf("failed to read directory: %w", err)
	}

	// Print header
	cmd.Printf("%-20s\t%s\n", "NAME", "ADDRESS")
	cmd.Printf("%-20s\t%s\n", "----", "-------")

	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), walletFileExt) {
			continue
		}
		
		// Read wallet file
		walletPath := filepath.Join(keystoreDir, entry.Name())
		data, err := os.ReadFile(walletPath)
		if err != nil {
			cmd.PrintErrf("Failed to read wallet %s: %v\n", entry.Name(), err)
			continue
		}

		// Parse wallet info
		var walletInfo struct {
			Address string `json:"address"`
		}
		if err := json.Unmarshal(data, &walletInfo); err != nil {
			cmd.PrintErrf("Failed to parse wallet %s: %v\n", entry.Name(), err)
			continue
		}

		name := strings.TrimSuffix(entry.Name(), walletFileExt)
		cmd.Printf("%-20s\t%s\n", name, walletInfo.Address)
	}

	return nil
}

func cmdGetWalletKey(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	if len(args) != 1 {
		return NewCmdParamsError("export-key command expects exactly 1 argument [wallet name]")
	}

	walletName := args[0]
	walletPath, err := getWalletPath(walletName)
	if err != nil {
		return err
	}

	jsonBytes, err := os.ReadFile(walletPath)
	if err != nil {
		return fmt.Errorf("failed to read wallet file: %w", err)
	}

	// Parse wallet info
	var walletInfo struct {
		PrivateKey string `json:"private_key"`
	}
	if err := json.Unmarshal(jsonBytes, &walletInfo); err != nil {
		return fmt.Errorf("failed to parse wallet file: %w", err)
	}

	cmd.Printf("Private key: %s\n", walletInfo.PrivateKey)

	return nil
}


func cmdImportWallet(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	walletName := args[0]
	privateKeyHex := args[1]

	// Add 0x prefix if not present
	if !strings.HasPrefix(privateKeyHex, "0x") {
		privateKeyHex = "0x" + privateKeyHex
	}

	// Validate private key format and derive public key/address
	privateKeyBytes, err := hexutil.Decode(privateKeyHex)
	if err != nil {
		return fmt.Errorf("invalid private key format: %w", err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// Check if wallet name already exists
	walletPath, err := getWalletPath(walletName)
	if err != nil {
		return err
	}
	if _, err := os.Stat(walletPath); err == nil {
		return fmt.Errorf("wallet with name '%s' already exists", walletName)
	}

	// Create wallet file
	walletInfo := struct {
		Address    string `json:"address"`
		PrivateKey string `json:"private_key"`
	}{
		Address:    address,
		PrivateKey: strings.TrimPrefix(privateKeyHex, "0x"),
	}

	jsonData, err := json.MarshalIndent(walletInfo, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal wallet info: %w", err)
	}

	if err := os.WriteFile(walletPath, jsonData, 0600); err != nil {
		return fmt.Errorf("failed to write wallet file: %w", err)
	}

	cmd.Printf("Wallet imported successfully:\nName: %s\nAddress: %s\n", walletName, address)
	return nil
}


func getPrivateKeyFromWallet(walletName string) (string, string, error) {
	if walletName == "" {
		// If no wallet specified, try to use the first available wallet
		keystoreDir := getDefaultKeystoreDir()
		entries, err := os.ReadDir(keystoreDir)
		if err != nil {
			if os.IsNotExist(err) {
				return "", "", fmt.Errorf("no wallets found in %s", keystoreDir)
			}
			return "", "", fmt.Errorf("failed to read keystore directory: %w", err)
		}

		for _, entry := range entries {
			if !strings.HasSuffix(entry.Name(), walletFileExt) {
				continue
			}
			walletName = strings.TrimSuffix(entry.Name(), walletFileExt)
			break
		}

		if walletName == "" {
			return "", "", fmt.Errorf("no wallets found")
		}
	}

	walletPath, err := getWalletPath(walletName)
	if err != nil {
		return "", "", err
	}

	data, err := os.ReadFile(walletPath)
	if err != nil {
		return "", "", fmt.Errorf("failed to read wallet file: %w", err)
	}

	var walletInfo struct {
		Address    string `json:"address"`
		PrivateKey string `json:"private_key"`
	}
	if err := json.Unmarshal(data, &walletInfo); err != nil {
		return "", "", fmt.Errorf("failed to parse wallet file: %w", err)
	}

	return walletInfo.PrivateKey, walletInfo.Address, nil
}

// TODO: Add a command to check wallet balance