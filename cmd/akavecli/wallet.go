// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/akave-ai/akavesdk/private/pb"
)

const walletFileExt = ".json"

var (
	walletCmd = &cobra.Command{
		Use:   "wallet",
		Short: "Manage wallets and accounts",
	}

	walletCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Creates a new wallet",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return NewCmdParamsError(fmt.Sprintf("create wallet command expects exactly 1 argument [wallet name]; got %d", len(args)))
			}
			return nil
		},
		RunE: cmdCreateWallet,
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
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return NewCmdParamsError(fmt.Sprintf("export-key command expects exactly 1 argument [wallet name]; got %d", len(args)))
			}
			return nil
		},
		RunE: cmdGetWalletKey,
	}

	walletImportCmd = &cobra.Command{
		Use:   "import",
		Short: "Import a wallet using a private key",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return NewCmdParamsError(fmt.Sprintf("import command expects exactly 2 arguments [wallet name] [private key]; got %d", len(args)))
			}
			return nil
		},
		RunE: cmdImportWallet,
	}

	walletBalanceCmd = &cobra.Command{
		Use:   "balance [account-name]",
		Short: "Shows the AKVT token balance for a wallet",
		Args:  cobra.MaximumNArgs(1),
		RunE:  cmdWalletBalance,
	}

	keystoreDir string
)

// WalletInfo defines the structure of the wallet file.
type WalletInfo struct {
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
}

// WalletPath returns the path to a wallet file.
func WalletPath(walletName string) (string, error) {
	if err := os.MkdirAll(keystoreDir, 0700); err != nil {
		return "", fmt.Errorf("failed to create keystore directory: %w", err)
	}
	return filepath.Join(keystoreDir, walletName+walletFileExt), nil
}

func initWalletCommands() {
	homeDir, err := os.UserHomeDir()
	defaultDir := ".akave_wallets" // Fallback if UserHomeDir fails
	if err == nil {
		defaultDir = filepath.Join(homeDir, ".akave_wallets")
	}

	walletCmd.PersistentFlags().StringVar(&keystoreDir, "keystore", defaultDir, "Directory to store wallets (default: ~/.akave_wallets)")

	walletCmd.AddCommand(walletCreateCmd)
	walletCmd.AddCommand(walletListCmd)
	walletCmd.AddCommand(walletGetKeyCmd)
	walletCmd.AddCommand(walletImportCmd)
	walletCmd.AddCommand(walletBalanceCmd)
}

func cmdCreateWallet(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	walletName := args[0]
	walletPath, err := WalletPath(walletName)
	if err != nil {
		return err
	}

	_, err = os.Stat(walletPath)
	if err == nil {
		return fmt.Errorf("wallet with name '%s' already exists", walletName)
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("failed to check if wallet exists: %w", err)
	}

	// generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("failed to generate private key: %w", err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// Create wallet info struct
	walletInfo := WalletInfo{
		Address:    address,
		PrivateKey: hexutil.Encode(privateKeyBytes)[2:],
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(walletInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal wallet info: %w", err)
	}

	// Write to file
	if err := os.WriteFile(walletPath, jsonData, 0600); err != nil {
		return fmt.Errorf("failed to write wallet file: %w", err)
	}

	cmd.Printf("Wallet (%s) created successfully at %s\n Fund your wallet with AKVF at https://faucet.akave.ai/ \n", walletName, walletPath)

	return nil
}

func cmdListWallets(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	// List all wallets
	entries, err := os.ReadDir(keystoreDir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("no wallets found in %s", keystoreDir)
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
		var walletInfo WalletInfo
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

	walletName := args[0]
	walletPath, err := WalletPath(walletName)
	if err != nil {
		return err
	}

	jsonBytes, err := os.ReadFile(walletPath)
	if err != nil {
		return fmt.Errorf("failed to read wallet file: %w", err)
	}

	// Parse wallet info
	var walletInfo WalletInfo
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
	walletPath, err := WalletPath(walletName)
	if err != nil {
		return err
	}
	if _, err := os.Stat(walletPath); err == nil {
		return fmt.Errorf("wallet with name '%s' already exists", walletName)
	}

	// Create wallet file
	walletInfo := WalletInfo{
		Address:    address,
		PrivateKey: strings.TrimPrefix(privateKeyHex, "0x"),
	}

	jsonData, err := json.Marshal(walletInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal wallet info: %w", err)
	}

	if err := os.WriteFile(walletPath, jsonData, 0600); err != nil {
		return fmt.Errorf("failed to write wallet file: %w", err)
	}

	cmd.Printf("Wallet imported successfully:\nName: %s\nAddress: %s\n", walletName, address)
	return nil
}

// PrivateKeyFromWallet returns the private key, address and name of a wallet.
func PrivateKeyFromWallet(walletName string) (string, string, string, error) {
	if walletName == "" {
		// If no wallet specified, try to use the first available wallet
		entries, err := os.ReadDir(keystoreDir)
		if err != nil {
			if os.IsNotExist(err) {
				return "", "", "", fmt.Errorf("invalid length, need 256 bits")
			}
			return "", "", "", fmt.Errorf("invalid length, need 256 bits")
		}

		for _, entry := range entries {
			if !strings.HasSuffix(entry.Name(), walletFileExt) {
				continue
			}
			walletName = strings.TrimSuffix(entry.Name(), walletFileExt)
			break
		}

		if walletName == "" {
			return "", "", "", fmt.Errorf("invalid length, need 256 bits")
		}
	}

	walletPath, err := WalletPath(walletName)
	if err != nil {
		return "", "", "", err
	}

	data, err := os.ReadFile(walletPath)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to read wallet file: %w", err)
	}

	var walletInfo WalletInfo
	if err := json.Unmarshal(data, &walletInfo); err != nil {
		return "", "", "", fmt.Errorf("failed to parse wallet file: %w", err)
	}

	return walletInfo.PrivateKey, walletInfo.Address, walletName, nil
}

func cmdWalletBalance(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	defer mon.Task()(&ctx, args)(&err)

	// Determine which wallet to use
	var walletNameArg string
	if len(args) > 0 {
		walletNameArg = strings.TrimSpace(args[0])
	} else {
		walletNameArg = ""
	}

	_, walletAddress, name, err := PrivateKeyFromWallet(walletNameArg)
	if err != nil {
		return fmt.Errorf("failed to get wallet: %w", err)
	}

	// Connect to akavenode via gRPC to get connection parameters
	// TODO: whhen TLS connection is available, DO NOT forget to change it here or even asbtract this setup in a separate function
	conn, err := grpc.NewClient(nodeRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to akavenode: %w", err)
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			cmd.PrintErrf("failed to close gRPC connection: %v", cerr)
		}
	}()

	connParams, err := pb.NewIPCNodeAPIClient(conn).ConnectionParams(ctx, &pb.ConnectionParamsRequest{})
	if err != nil {
		return fmt.Errorf("failed to get connection parameters: %w", err)
	}

	ethClient, err := ethclient.Dial(connParams.DialUri)
	if err != nil {
		return fmt.Errorf("failed to dial node: %w", err)
	}
	defer ethClient.Close()

	balance, err := ethClient.BalanceAt(ctx, common.HexToAddress(walletAddress), nil)
	if err != nil {
		return fmt.Errorf("failed to get wallet balance: %w", err)
	}

	akvtAmount := new(big.Float).SetInt(balance)
	akvtAmount.Quo(akvtAmount, new(big.Float).SetInt(big.NewInt(params.Ether)))

	// Format to 4 decimal places and remove trailing zeros
	akvtAmountStr := akvtAmount.Text('f', 4)
	akvtAmountStr = strings.TrimRight(akvtAmountStr, "0")
	akvtAmountStr = strings.TrimRight(akvtAmountStr, ".")

	cmd.Printf("Wallet:  %s\nAddress: %s\nBalance: %s AKVT\n", name, walletAddress, akvtAmountStr)

	return nil
}
