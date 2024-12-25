// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/spf13/cobra"
	"golang.org/x/term"
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
)

func initWalletCommands() {
	walletCmd.AddCommand(walletCreateCmd)
	walletCmd.AddCommand(walletListCmd)
	walletCmd.AddCommand(walletGetKeyCmd)
}

// getDefaultKeystoreDir returns the default keystore directory
func getDefaultKeystoreDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ".akave_wallets"
	}
	return filepath.Join(homeDir, ".akave_wallets")
}

// getKeystore initializes and returns a keystore for the given account
func getKeystore(accountName string) (*keystore.KeyStore, error) {
	if accountName == "" {
		return nil, fmt.Errorf("account name is required")
	}
	
	accountDir := filepath.Join(keystoreDir, accountName)
	if err := os.MkdirAll(accountDir, 0700); err != nil {
		return nil, fmt.Errorf("failed to create keystore directory: %w", err)
	}
	
	return keystore.NewKeyStore(accountDir, keystore.StandardScryptN, keystore.StandardScryptP), nil
}

// getAccount returns the first account from the keystore for the given account name
func getAccount(accountName string) (accounts.Account, error) {
	ks, err := getKeystore(accountName)
	if err != nil {
		return accounts.Account{}, err
	}

	accts := ks.Accounts()
	if len(accts) == 0 {
		return accounts.Account{}, fmt.Errorf("no accounts found for %s", accountName)
	}

	return accts[0], nil
}

func cmdCreateWallet(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	if len(args) != 1 {
		return NewCmdParamsError("create wallet command expects exactly 1 argument [wallet name]")
	}

	walletName := args[0]
	
	// Get password
	cmd.PrintErrf("Enter password for new wallet: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("failed to read password: %w", err)
	}
	cmd.PrintErrln()

	// Confirm password
	cmd.PrintErrf("Confirm password: ")
	confirmPassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("failed to read confirmation password: %w", err)
	}
	cmd.PrintErrln()

	// Verify passwords match
	if string(password) != string(confirmPassword) {
		return fmt.Errorf("passwords do not match")
	}

	// Create keystore and account
	ks, err := getKeystore(walletName)
	if err != nil {
		return fmt.Errorf("failed to initialize keystore: %w", err)
	}

	account, err := ks.NewAccount(string(password))
	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}

	cmd.PrintErrf("Wallet created: Address=%s\n", account.Address.Hex())
	cmd.PrintErrf("Keystore location: %s\n", filepath.Join(keystoreDir, walletName))

	return nil
}

func cmdListWallets(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	// List all subdirectories in keystoreDir
	entries, err := os.ReadDir(keystoreDir)
	if err != nil {
		if os.IsNotExist(err) {
			cmd.PrintErrln("No wallets found")
			return nil
		}
		return fmt.Errorf("failed to read keystore directory: %w", err)
	}

	found := false
	for _, entry := range entries {
		if entry.IsDir() {
			found = true
			// Get account details for this wallet
			ks, err := getKeystore(entry.Name())
			if err != nil {
				cmd.PrintErrf("Warning: failed to read wallet %s: %v\n", entry.Name(), err)
				continue
			}

			accounts := ks.Accounts()
			if len(accounts) > 0 {
				cmd.PrintErrf("Wallet: %s\n", entry.Name())
				for _, acc := range accounts {
					cmd.PrintErrf("  Address: %s\n", acc.Address.Hex())
				}
			}
		}
	}

	if !found {
		cmd.PrintErrln("No wallets found")
	}

	return nil
}

func cmdGetWalletKey(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	walletName := args[0]

	// Get the account
	account, err := getAccount(walletName)
	if err != nil {
		return fmt.Errorf("failed to get account: %w", err)
	}

	// Get password
	cmd.PrintErrf("Enter password for wallet %s: ", walletName)
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("failed to read password: %w", err)
	}
	cmd.PrintErrln()

	// Read the keystore file
	keystorePath := account.URL.Path
	jsonBytes, err := os.ReadFile(keystorePath)
	if err != nil {
		return fmt.Errorf("failed to read keystore file: %w", err)
	}

	// Decrypt the key
	key, err := keystore.DecryptKey(jsonBytes, string(password))
	if err != nil {
		return fmt.Errorf("failed to decrypt keystore (wrong password?): %w", err)
	}

	// Convert private key to hex
	privateKeyHex := fmt.Sprintf("%x", key.PrivateKey.D)
	cmd.PrintErrf("Private key: %s\n", privateKeyHex)

	return nil
} 