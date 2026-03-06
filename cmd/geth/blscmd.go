package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/blskeystore"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls"
	"github.com/urfave/cli/v2"
)

// ##CROSS: bls seal

var (
	showSecretKeyFlag = &cli.BoolFlag{
		Name:    "show-secret-key",
		Aliases: []string{"show-private-key"},
		Usage:   "Show the BLS secret key used for this command",
	}
	chainIDFlag = &cli.StringFlag{
		Name:  "chain-id",
		Usage: "ID or name of the network that the validator will be used on",
	}
)

var (
	blsCommand = &cli.Command{
		Name:  "bls",
		Usage: "Manage BLS keys",
		Description: `

Manage BLS validator keys, list all existing keys, generate a new key pair,
import a private key from a plain text hex file, update an existing key,
or generate a proof of possession for a public key.

It supports interactive mode, when you are prompted for password as well as
non-interactive mode where passwords are supplied via a given password file.
Non-interactive mode is only meant for scripted use on test networks or known
safe environments.

Make sure you remember the password you gave when creating a new key (with
either new or import). Without it you are not able to decrypt your BLS key.

Keys are stored under <DATADIR>/bls-keystore.
It is safe to transfer the entire directory or the individual keys therein
between nodes by simply copying.

Make sure you backup your keys regularly.`,
		Subcommands: []*cli.Command{
			{
				Name:   "list",
				Usage:  "Print summary of existing BLS keys",
				Action: blsList,
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.BLSKeyStoreDirFlag,
				},
				Description: `
Prints a short summary of all BLS keys in the keystore`,
			},
			{
				Name:   "new",
				Usage:  "Generate a new BLS key pair",
				Action: blsNew,
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.BLSKeyStoreDirFlag,
					utils.PasswordFileFlag,
					utils.LightKDFFlag,
					showSecretKeyFlag,
				},
				Description: `
    geth bls new

Generates a new BLS key pair and stores it in encrypted format.

The key is saved in encrypted format, you are prompted for a password.

You must remember this password to decrypt your BLS key in the future.

For non-interactive use the password can be specified with the --password flag.
`,
			},
			{
				Name:      "update",
				Usage:     "Update an existing BLS key",
				Action:    blsUpdate,
				ArgsUsage: "<pubkey>",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.BLSKeyStoreDirFlag,
					utils.LightKDFFlag,
				},
				Description: `
    geth bls update <pubkey>

Updates an existing BLS key.

The key is re-encrypted with a new password, you are prompted for the old
password to unlock the key and a new password to save the updated file.

The pubkey argument is the hex-encoded BLS public key (or a unique prefix).
`,
			},
			{
				Name:      "import",
				Usage:     "Import a BLS private key from a hex file into encrypted storage",
				Action:    blsImport,
				ArgsUsage: "<keyFile>",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.BLSKeyStoreDirFlag,
					utils.PasswordFileFlag,
					utils.LightKDFFlag,
					showSecretKeyFlag,
				},
				Description: `
    geth bls import <keyfile>

Imports an unencrypted BLS private key from <keyfile> and stores it encrypted.
The keyfile is assumed to contain the private key in hexadecimal format (32 bytes).

You must remember this password to decrypt your BLS key in the future.

For non-interactive use the password can be specified with the --password flag.

Note:
As you can directly copy your encrypted BLS keys to another node,
this import mechanism is not needed when you transfer a key between nodes.
`,
			},
			{
				Name:      "decrypt",
				Usage:     "Decrypts a keystore file to get the secret key",
				Action:    blsDecrypt,
				ArgsUsage: "<pubkey>",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.BLSKeyStoreDirFlag,
					utils.PasswordFileFlag,
				},
				Description: `
    geth bls decrypt <pubkey>

Decrypts the keystore file for the given public key and prints the secret key.
`,
			},
			{
				Name:      "generate-proof",
				Aliases:   []string{"pop"},
				Usage:     "Generates a proof of possession for the given public key",
				Action:    blsGenerateProof,
				ArgsUsage: "<pubkey>",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.BLSKeyStoreDirFlag,
					utils.PasswordFileFlag,
					chainIDFlag,
				},
				Description: `
    geth bls generate-proof --chain-id <chain-id-or-name> <pubkey>
	
Generates a proof of possession for the given public key. The proof is used to prove the ownership of the BLS public key when registering a validator.
`,
			},
		},
	}
)

func getBLSParams(ctx *cli.Context) (string, int, int) {
	var dir string
	if ctx.IsSet(utils.BLSKeyStoreDirFlag.Name) {
		dir = ctx.String(utils.BLSKeyStoreDirFlag.Name)
	}

	cfg := loadBaseConfig(ctx)

	if dir == "" {
		if cfg.Node.DataDir == "" {
			utils.Fatalf("Cannot determine BLS keystore directory, please set --datadir or --bls.keystore")
		}
		dir = filepath.Join(cfg.Node.DataDir, "bls-keystore")
	}

	var scryptN, scryptP int
	if cfg.Node.UseLightweightKDF {
		scryptN, scryptP = keystore.LightScryptN, keystore.LightScryptP
	} else {
		scryptN, scryptP = keystore.StandardScryptN, keystore.StandardScryptP
	}
	return dir, scryptN, scryptP
}

func getChainID(ctx *cli.Context) uint64 {
	chainID := ctx.String(chainIDFlag.Name)
	if chainID == "" {
		return 0
	}
	switch chainID {
	case "cross":
		return 612055
	case "zonezero":
		return 612044
	case "crossdev3":
		return 612088
	case "crossdev":
		return 612066
	}
	id, _ := strconv.ParseUint(chainID, 0, 64)
	return id
}

func blsList(ctx *cli.Context) error {
	dir, _, _ := getBLSParams(ctx)
	keys, err := blskeystore.ListKeys(dir)
	if err != nil {
		utils.Fatalf("Failed to list BLS keys: %v", err)
	}
	if len(keys) == 0 {
		fmt.Println("No BLS keys found.")
		return nil
	}
	for i, k := range keys {
		fmt.Printf("BLS Key #%d: {0x%s} %s\n", i, k.PubKey, k.Path)
	}
	return nil
}

func blsNew(ctx *cli.Context) error {
	dir, scryptN, scryptP := getBLSParams(ctx)

	password, ok := readPasswordFromFile(ctx.Path(utils.PasswordFileFlag.Name))
	if !ok {
		password = utils.GetPassPhrase("Your new BLS key is locked with a password. Please give a password. Do not forget this password.", true)
	}

	secretKey, err := bls.RandKey()
	if err != nil {
		utils.Fatalf("Failed to generate BLS key: %v", err)
	}
	key, keyfilepath, err := blskeystore.StoreKey(dir, secretKey, password, scryptN, scryptP)
	if err != nil {
		utils.Fatalf("Failed to store BLS key: %v", err)
	}

	fmt.Printf("\nYour new BLS key was generated\n\n")
	fmt.Printf("Public key:             0x%s\n", hex.EncodeToString(key.PublicKey.Marshal()))
	if ctx.Bool(showSecretKeyFlag.Name) {
		fmt.Printf("Secret key:             0x%s\n", hex.EncodeToString(key.SecretKey.Marshal()))
	}
	fmt.Printf("Path of the key file:   %s\n\n", keyfilepath)
	fmt.Printf("- You can share your public key with anyone. It is needed for validator registration.\n")
	fmt.Printf("- You must NEVER share the secret key with anyone! The key controls your validator!\n")
	fmt.Printf("- You must BACKUP your key file! Without the key, it's impossible to sign blocks!\n")
	fmt.Printf("- You must REMEMBER your password! Without the password, it's impossible to decrypt the key!\n\n")
	return nil
}

func blsUpdate(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		utils.Fatalf("No BLS public key specified to update")
	}
	dir, scryptN, scryptP := getBLSParams(ctx)

	for _, pubKeyHex := range ctx.Args().Slice() {
		keyFile, err := blskeystore.FindKeyByPubKey(dir, pubKeyHex)
		if err != nil {
			utils.Fatalf("Failed to find BLS key: %v", err)
		}

		newPassword := utils.GetPassPhrase("Please give a NEW password. Do not forget this password.", true)
		updateFn := func(attempt int) error {
			prompt := fmt.Sprintf("Please provide the OLD password for BLS key %s | Attempt %d/%d", pubKeyHex, attempt+1, 3)
			password := utils.GetPassPhrase(prompt, false)
			return blskeystore.UpdateKey(keyFile, password, newPassword, scryptN, scryptP)
		}
		// Let user attempt unlock thrice.
		err = updateFn(0)
		for attempts := 1; attempts < 3 && errors.Is(err, keystore.ErrDecrypt); attempts++ {
			err = updateFn(attempts)
		}
		if err != nil {
			return fmt.Errorf("could not update BLS key: %w", err)
		}
	}
	return nil
}

func blsImport(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		utils.Fatalf("keyfile must be given as the only argument")
	}
	keyfile := ctx.Args().First()

	keyData, err := os.ReadFile(keyfile)
	if err != nil {
		utils.Fatalf("Failed to read key file: %v", err)
	}
	keyString := strings.TrimSpace(string(keyData))
	if keyString == "" {
		utils.Fatalf("Empty key file")
	}
	keyBytes := common.FromHex(keyString)
	if len(keyBytes) != 32 {
		utils.Fatalf("Invalid BLS key length: expected 32 bytes, got %d", len(keyBytes))
	}
	secretKey, err := bls.SecretKeyFromBytes(keyBytes)
	if err != nil {
		utils.Fatalf("Invalid BLS key: %v", err)
	}

	dir, scryptN, scryptP := getBLSParams(ctx)

	password, ok := readPasswordFromFile(ctx.Path(utils.PasswordFileFlag.Name))
	if !ok {
		password = utils.GetPassPhrase("Your BLS key will be locked with a password. Please give a password. Do not forget this password.", true)
	}

	key, keyfilepath, err := blskeystore.StoreKey(dir, secretKey, password, scryptN, scryptP)
	if err != nil {
		utils.Fatalf("Failed to import BLS key: %v", err)
	}

	fmt.Printf("BLS key imported successfully.\n")
	fmt.Printf("Public key:             0x%s\n", hex.EncodeToString(key.PublicKey.Marshal()))
	if ctx.Bool(showSecretKeyFlag.Name) {
		fmt.Printf("Secret key:             0x%s\n", hex.EncodeToString(key.SecretKey.Marshal()))
	}
	fmt.Printf("Path of the key file:   %s\n", keyfilepath)
	return nil
}

func blsDecrypt(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		utils.Fatalf("No BLS public key specified to decrypt")
	}
	dir, _, _ := getBLSParams(ctx)

	pubKeyHex := ctx.Args().First()
	keyFile, err := blskeystore.FindKeyByPubKey(dir, pubKeyHex)
	if err != nil {
		utils.Fatalf("Failed to find BLS key: %v", err)
	}

	password, ok := readPasswordFromFile(ctx.Path(utils.PasswordFileFlag.Name))
	if !ok {
		password = utils.GetPassPhrase(fmt.Sprintf("Please provide the password for BLS key %s", pubKeyHex), false)
	}

	keyjson, err := os.ReadFile(keyFile)
	if err != nil {
		utils.Fatalf("Failed to read BLS key file: %v", err)
	}
	key, err := blskeystore.DecryptBLSKey(keyjson, password)
	if err != nil {
		utils.Fatalf("Failed to decrypt BLS key: %v", err)
	}

	fmt.Printf("0x%s\n", hex.EncodeToString(key.SecretKey.Marshal()))
	return nil
}

func blsGenerateProof(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		utils.Fatalf("No BLS public key specified to generate proof")
	}

	chainID := getChainID(ctx)
	if chainID == 0 {
		utils.Fatalf("Chain ID is not set")
	}

	dir, _, _ := getBLSParams(ctx)

	pubKeyHex := ctx.Args().First()
	keyFile, err := blskeystore.FindKeyByPubKey(dir, pubKeyHex)
	if err != nil {
		utils.Fatalf("Failed to find BLS key: %v", err)
	}

	password, ok := readPasswordFromFile(ctx.Path(utils.PasswordFileFlag.Name))
	if !ok {
		password = utils.GetPassPhrase(fmt.Sprintf("Please provide the password for BLS key %s", pubKeyHex), false)
	}

	keyjson, err := os.ReadFile(keyFile)
	if err != nil {
		utils.Fatalf("Failed to read BLS key file: %v", err)
	}
	key, err := blskeystore.DecryptBLSKey(keyjson, password)
	if err != nil {
		utils.Fatalf("Failed to decrypt BLS key: %v", err)
	}

	publicKeyByte := key.PublicKey.Marshal()
	chainIDByte := new(big.Int).SetUint64(chainID).FillBytes(make([]byte, 32))
	msgByte := crypto.Keccak256(publicKeyByte, chainIDByte)
	proof := key.SecretKey.Sign(msgByte)
	if len(proof.Marshal()) == 0 {
		return fmt.Errorf("Failed to sign message")
	}
	fmt.Printf("Proof: 0x%x\n", proof.Marshal())
	return nil
}
