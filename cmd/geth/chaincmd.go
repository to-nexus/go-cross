// Copyright 2015 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"runtime"
	"slices"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/history"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/internal/debug"
	"github.com/ethereum/go-ethereum/internal/era"
	"github.com/ethereum/go-ethereum/internal/flags"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls"
	"github.com/urfave/cli/v2"
)

var (
	initCommand = &cli.Command{
		Action:    initGenesis,
		Name:      "init",
		Usage:     "Bootstrap and initialize a new genesis block",
		ArgsUsage: "<genesisPath>",
		Flags: slices.Concat([]cli.Flag{
			utils.CachePreimagesFlag,
			utils.OverridePrague,
			utils.OverrideVerkle,
		}, utils.DatabaseFlags),
		Description: `
The init command initializes a new genesis block and definition for the network.
This is a destructive action and changes the network in which you will be
participating.

It expects the genesis file as argument.`,
	}
	dumpGenesisCommand = &cli.Command{
		Action:    dumpGenesis,
		Name:      "dumpgenesis",
		Usage:     "Dumps genesis block JSON configuration to stdout",
		ArgsUsage: "",
		Flags:     slices.Concat([]cli.Flag{utils.DataDirFlag}, utils.NetworkFlags),
		Description: `
The dumpgenesis command prints the genesis configuration of the network preset
if one is set.  Otherwise it prints the genesis from the datadir.`,
	}
	// ##CROSS: fork breakpoint
	makeBreakpointGenesisCommand = &cli.Command{
		Action:    makeBreakpointGenesis,
		Name:      "make-breakpoint-genesis",
		Usage:     "Builds a custom genesis with Breakpoint active from block one",
		ArgsUsage: "[<inputGenesis> <outputGenesis>]",
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:  "validators",
				Usage: "Number of validators to generate without input arguments",
				Value: 3,
			},
			&cli.BoolFlag{
				Name:  "json",
				Usage: "Print generated keys as JSON",
			},
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Overwrite the output file if it already exists",
			},
		},
		Description: `
The make-breakpoint-genesis command installs and initializes the Breakpoint
contracts using the Istanbul PoSA configuration in the input genesis. Without
arguments, it generates genesis.json and prints all keys.`,
	}
	// ##
	importCommand = &cli.Command{
		Action:    importChain,
		Name:      "import",
		Usage:     "Import a blockchain file",
		ArgsUsage: "<filename> (<filename 2> ... <filename N>) ",
		Flags: slices.Concat([]cli.Flag{
			utils.GCModeFlag,
			utils.SnapshotFlag,
			utils.CacheFlag,
			utils.CacheDatabaseFlag,
			utils.CacheTrieFlag,
			utils.CacheGCFlag,
			utils.CacheSnapshotFlag,
			utils.CacheNoPrefetchFlag,
			utils.CachePreimagesFlag,
			utils.NoCompactionFlag,
			utils.MetricsEnabledFlag,
			utils.MetricsEnabledExpensiveFlag,
			utils.MetricsHTTPFlag,
			utils.MetricsPortFlag,
			utils.MetricsEnableInfluxDBFlag,
			utils.MetricsEnableInfluxDBV2Flag,
			utils.MetricsInfluxDBEndpointFlag,
			utils.MetricsInfluxDBDatabaseFlag,
			utils.MetricsInfluxDBUsernameFlag,
			utils.MetricsInfluxDBPasswordFlag,
			utils.MetricsInfluxDBTagsFlag,
			utils.MetricsInfluxDBTokenFlag,
			utils.MetricsInfluxDBBucketFlag,
			utils.MetricsInfluxDBOrganizationFlag,
			utils.TxLookupLimitFlag,
			utils.VMTraceFlag,
			utils.VMTraceJsonConfigFlag,
			utils.TransactionHistoryFlag,
			utils.LogHistoryFlag,
			utils.LogNoHistoryFlag,
			utils.LogExportCheckpointsFlag,
			utils.StateHistoryFlag,
		}, utils.DatabaseFlags, debug.Flags),
		Before: func(ctx *cli.Context) error {
			flags.MigrateGlobalFlags(ctx)
			return debug.Setup(ctx)
		},
		Description: `
The import command allows the import of blocks from an RLP-encoded format. This format can be a single file
containing multiple RLP-encoded blocks, or multiple files can be given.

If only one file is used, an import error will result in the entire import process failing. If
multiple files are processed, the import process will continue even if an individual RLP file fails
to import successfully.`,
	}
	exportCommand = &cli.Command{
		Action:    exportChain,
		Name:      "export",
		Usage:     "Export blockchain into file",
		ArgsUsage: "<filename> [<blockNumFirst> <blockNumLast>]",
		Flags:     slices.Concat([]cli.Flag{utils.CacheFlag}, utils.DatabaseFlags),
		Description: `
Requires a first argument of the file to write to.
Optional second and third arguments control the first and
last block to write. In this mode, the file will be appended
if already existing. If the file ends with .gz, the output will
be gzipped.`,
	}
	importHistoryCommand = &cli.Command{
		Action:    importHistory,
		Name:      "import-history",
		Usage:     "Import an Era archive",
		ArgsUsage: "<dir>",
		Flags:     slices.Concat([]cli.Flag{utils.TxLookupLimitFlag, utils.TransactionHistoryFlag}, utils.DatabaseFlags, utils.NetworkFlags),
		Description: `
The import-history command will import blocks and their corresponding receipts
from Era archives.
`,
	}
	exportHistoryCommand = &cli.Command{
		Action:    exportHistory,
		Name:      "export-history",
		Usage:     "Export blockchain history to Era archives",
		ArgsUsage: "<dir> <first> <last>",
		Flags:     utils.DatabaseFlags,
		Description: `
The export-history command will export blocks and their corresponding receipts
into Era archives. Eras are typically packaged in steps of 8192 blocks.
`,
	}
	importPreimagesCommand = &cli.Command{
		Action:    importPreimages,
		Name:      "import-preimages",
		Usage:     "Import the preimage database from an RLP stream",
		ArgsUsage: "<datafile>",
		Flags:     slices.Concat([]cli.Flag{utils.CacheFlag}, utils.DatabaseFlags),
		Description: `
The import-preimages command imports hash preimages from an RLP encoded stream.
It's deprecated, please use "geth db import" instead.
`,
	}

	dumpCommand = &cli.Command{
		Action:    dump,
		Name:      "dump",
		Usage:     "Dump a specific block from storage",
		ArgsUsage: "[? <blockHash> | <blockNum>]",
		Flags: slices.Concat([]cli.Flag{
			utils.CacheFlag,
			utils.IterativeOutputFlag,
			utils.ExcludeCodeFlag,
			utils.ExcludeStorageFlag,
			utils.IncludeIncompletesFlag,
			utils.StartKeyFlag,
			utils.DumpLimitFlag,
		}, utils.DatabaseFlags),
		Description: `
This command dumps out the state for a given block (or latest, if none provided).
`,
	}

	pruneCommand = &cli.Command{
		Action:    pruneHistory,
		Name:      "prune-history",
		Usage:     "Prune blockchain history (block bodies and receipts) up to the merge block",
		ArgsUsage: "",
		Flags:     utils.DatabaseFlags,
		Description: `
The prune-history command removes historical block bodies and receipts from the
blockchain database up to the merge block, while preserving block headers. This
helps reduce storage requirements for nodes that don't need full historical data.`,
	}
)

// ##CROSS: fork breakpoint
type (
	breakpointGenesisKeys struct {
		admin      *ecdsa.PrivateKey
		validators []breakpointValidatorKeys
	}
	breakpointValidatorKeys struct {
		validator *ecdsa.PrivateKey
		operator  *ecdsa.PrivateKey
		signer    bls.SecretKey
	}
)

func makeBreakpointGenesis(ctx *cli.Context) error {
	var (
		genesis    *core.Genesis
		keys       *breakpointGenesisKeys
		outputPath string
	)
	switch ctx.Args().Len() {
	case 0:
		var err error
		genesis, keys, err = makeDefaultBreakpointGenesis(ctx.Uint("validators"))
		if err != nil {
			return err
		}
		outputPath = "genesis.json"
	case 2:
		if ctx.IsSet("validators") || ctx.Bool("json") {
			return errors.New("--validators and --json are only available without input arguments")
		}
		inputPath := ctx.Args().Get(0)
		outputPath = ctx.Args().Get(1)
		input, err := os.Open(inputPath)
		if err != nil {
			return fmt.Errorf("open input genesis: %w", err)
		}
		defer input.Close()

		genesis = new(core.Genesis)
		if err := json.NewDecoder(input).Decode(genesis); err != nil {
			return fmt.Errorf("decode input genesis: %w", err)
		}
		if err := core.MakeBreakpointGenesis(genesis); err != nil {
			return fmt.Errorf("build Breakpoint genesis: %w", err)
		}
	default:
		return fmt.Errorf("usage: %s", ctx.Command.ArgsUsage)
	}

	flags := os.O_WRONLY | os.O_CREATE | os.O_EXCL
	if ctx.Bool("force") {
		flags = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}
	output, err := os.OpenFile(outputPath, flags, 0o644)
	if err != nil {
		return fmt.Errorf("create output genesis: %w", err)
	}
	defer output.Close()
	encoder := json.NewEncoder(output)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(genesis); err != nil {
		return fmt.Errorf("encode output genesis: %w", err)
	}
	if keys != nil {
		if err := printBreakpointGenesisKeys(os.Stdout, keys, ctx.Bool("json")); err != nil {
			return fmt.Errorf("print generated keys: %w", err)
		}
	}
	return nil
}

func makeDefaultBreakpointGenesis(validatorCount uint) (*core.Genesis, *breakpointGenesisKeys, error) {
	if validatorCount == 0 {
		return nil, nil, errors.New("validator count must be greater than zero")
	}
	admin, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, fmt.Errorf("generate PoSA admin key: %w", err)
	}
	keys := &breakpointGenesisKeys{admin: admin, validators: make([]breakpointValidatorKeys, validatorCount)}
	validators := make([]params.PoSAValidator, len(keys.validators))
	validatorAddrs := make([]common.Address, len(keys.validators))
	for i := range keys.validators {
		validator, err := crypto.GenerateKey()
		if err != nil {
			return nil, nil, fmt.Errorf("generate validator %d key: %w", i+1, err)
		}
		operator, err := crypto.GenerateKey()
		if err != nil {
			return nil, nil, fmt.Errorf("generate operator %d key: %w", i+1, err)
		}
		signer, err := bls.RandKey()
		if err != nil {
			return nil, nil, fmt.Errorf("generate signer %d key: %w", i+1, err)
		}
		keys.validators[i] = breakpointValidatorKeys{validator: validator, operator: operator, signer: signer}
		validatorAddrs[i] = crypto.PubkeyToAddress(validator.PublicKey)
		validators[i] = params.PoSAValidator{
			ID:        fmt.Sprintf("validator%d", i+1),
			Operator:  crypto.PubkeyToAddress(operator.PublicKey),
			Validator: validatorAddrs[i],
			Signer:    signer.PublicKey().Marshal(),
		}
	}

	config := *params.CrossDev3ChainConfig
	istanbulConfig := *config.Istanbul
	posaConfig := *istanbulConfig.PoSA
	config.Istanbul = &istanbulConfig
	istanbulConfig.PoSA = &posaConfig
	istanbulConfig.Validators = validatorAddrs
	posaConfig.Admin = crypto.PubkeyToAddress(admin.PublicKey)
	posaConfig.Validators = validators

	adminBalance := new(big.Int).Mul(big.NewInt(100_000_000_000), big.NewInt(params.Ether))
	genesis := &core.Genesis{
		Config:     &config,
		Nonce:      0x90aa,
		Timestamp:  0x5f1663fc,
		ExtraData:  hexutil.MustDecode("0xc680c0c080c080"),
		GasLimit:   105000000,
		Difficulty: istanbul.DefaultDifficulty,
		Mixhash:    types.IstanbulDigest,
		Alloc: types.GenesisAlloc{
			posaConfig.Admin: {Balance: adminBalance},
		},
	}
	if err := core.MakeBreakpointGenesis(genesis); err != nil {
		return nil, nil, fmt.Errorf("build default Breakpoint genesis: %w", err)
	}
	return genesis, keys, nil
}

func printBreakpointGenesisKeys(w io.Writer, keys *breakpointGenesisKeys, jsonOutput bool) error {
	if jsonOutput {
		validators := make([]map[string]any, len(keys.validators))
		for i, validatorKeys := range keys.validators {
			validators[i] = map[string]any{
				"id": fmt.Sprintf("validator%d", i+1),
				"validator": map[string]string{
					"address":    crypto.PubkeyToAddress(validatorKeys.validator.PublicKey).Hex(),
					"privateKey": "0x" + hex.EncodeToString(crypto.FromECDSA(validatorKeys.validator)),
				},
				"operator": map[string]string{
					"address":    crypto.PubkeyToAddress(validatorKeys.operator.PublicKey).Hex(),
					"privateKey": "0x" + hex.EncodeToString(crypto.FromECDSA(validatorKeys.operator)),
				},
				"signer": map[string]string{
					"publicKey": "0x" + hex.EncodeToString(validatorKeys.signer.PublicKey().Marshal()),
					"secretKey": "0x" + hex.EncodeToString(validatorKeys.signer.Marshal()),
				},
			}
		}
		output := map[string]any{
			"admin": map[string]string{
				"address":    crypto.PubkeyToAddress(keys.admin.PublicKey).Hex(),
				"privateKey": "0x" + hex.EncodeToString(crypto.FromECDSA(keys.admin)),
			},
			"validators": validators,
		}
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		return encoder.Encode(output)
	}
	fmt.Fprintf(w, "PoSA admin\n  address: %s\n  private key: 0x%s\n", crypto.PubkeyToAddress(keys.admin.PublicKey), hex.EncodeToString(crypto.FromECDSA(keys.admin)))
	for i, validatorKeys := range keys.validators {
		fmt.Fprintf(w, "Validator %d\n", i+1)
		fmt.Fprintf(w, "  validator address: %s\n", crypto.PubkeyToAddress(validatorKeys.validator.PublicKey))
		fmt.Fprintf(w, "  validator private key: 0x%s\n", hex.EncodeToString(crypto.FromECDSA(validatorKeys.validator)))
		fmt.Fprintf(w, "  operator address: %s\n", crypto.PubkeyToAddress(validatorKeys.operator.PublicKey))
		fmt.Fprintf(w, "  operator private key: 0x%s\n", hex.EncodeToString(crypto.FromECDSA(validatorKeys.operator)))
		fmt.Fprintf(w, "  signer public key: 0x%s\n", hex.EncodeToString(validatorKeys.signer.PublicKey().Marshal()))
		fmt.Fprintf(w, "  signer secret key: 0x%s\n", hex.EncodeToString(validatorKeys.signer.Marshal()))
	}
	return nil
}

// ##

// initGenesis will initialise the given JSON format genesis file and writes it as
// the zero'd block (i.e. genesis) or will fail hard if it can't succeed.
func initGenesis(ctx *cli.Context) error {
	var genesis *core.Genesis

	if ctx.Args().Len() != 1 {
		utils.Fatalf("need genesis.json file as the only argument")
	}
	genesisPath := ctx.Args().First()
	if len(genesisPath) == 0 {
		utils.Fatalf("invalid path to genesis file")
	}

	switch genesisPath {
	// ##CROSS: config
	case "cross":
		genesis = core.DefaultCrossGenesisBlock()
	case "zonezero":
		genesis = core.DefaultZoneZeroGenesisBlock()
	case "crossdev3", "onedev3":
		genesis = core.DefaultCrossDev3GenesisBlock()
	case "crossdev", "onedev":
		genesis = core.DefaultCrossDevGenesisBlock()
	// ##
	default:
		file, err := os.Open(genesisPath)
		if err != nil {
			utils.Fatalf("Failed to read genesis file: %v", err)
		}
		defer file.Close()

		genesis = new(core.Genesis)
		if err := json.NewDecoder(file).Decode(genesis); err != nil {
			utils.Fatalf("invalid genesis file: %v", err)
		}
	}
	// Open and initialise both full and light databases
	stack, _ := makeConfigNode(ctx)
	defer stack.Close()

	var overrides core.ChainOverrides
	if ctx.IsSet(utils.OverridePrague.Name) {
		v := ctx.Uint64(utils.OverridePrague.Name)
		overrides.OverridePrague = &v
	}
	if ctx.IsSet(utils.OverrideVerkle.Name) {
		v := ctx.Uint64(utils.OverrideVerkle.Name)
		overrides.OverrideVerkle = &v
	}

	chaindb, err := stack.OpenDatabaseWithFreezer("chaindata", 0, 0, ctx.String(utils.AncientFlag.Name), "", false)
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	defer chaindb.Close()

	triedb := utils.MakeTrieDatabase(ctx, chaindb, ctx.Bool(utils.CachePreimagesFlag.Name), false, genesis.IsVerkle())
	defer triedb.Close()

	_, hash, compatErr, err := core.SetupGenesisBlockWithOverride(chaindb, triedb, genesis, &overrides)
	if err != nil {
		utils.Fatalf("Failed to write genesis block: %v", err)
	}
	if compatErr != nil {
		utils.Fatalf("Failed to write chain config: %v", compatErr)
	}
	log.Info("Successfully wrote genesis state", "database", "chaindata", "hash", hash)

	return nil
}

func dumpGenesis(ctx *cli.Context) error {
	// check if there is a testnet preset enabled
	var genesis *core.Genesis
	if utils.IsNetworkPreset(ctx) {
		genesis = utils.MakeGenesis(ctx)
	} else if ctx.IsSet(utils.DeveloperFlag.Name) && !ctx.IsSet(utils.DataDirFlag.Name) {
		genesis = core.DeveloperGenesisBlock(11_500_000, nil)
	}

	if genesis != nil {
		if err := json.NewEncoder(os.Stdout).Encode(genesis); err != nil {
			utils.Fatalf("could not encode genesis: %s", err)
		}
		return nil
	}

	// dump whatever already exists in the datadir
	stack, _ := makeConfigNode(ctx)

	db, err := stack.OpenDatabase("chaindata", 0, 0, "", true)
	if err != nil {
		return err
	}
	defer db.Close()

	genesis, err = core.ReadGenesis(db)
	if err != nil {
		utils.Fatalf("failed to read genesis: %s", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(*genesis); err != nil {
		utils.Fatalf("could not encode stored genesis: %s", err)
	}

	return nil
}

func importChain(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		utils.Fatalf("This command requires an argument.")
	}
	stack, cfg := makeConfigNode(ctx)
	defer stack.Close()

	// Start metrics export if enabled
	utils.SetupMetrics(&cfg.Metrics)

	chain, db := utils.MakeChain(ctx, stack, false)
	defer db.Close()

	// Start periodically gathering memory profiles
	var peakMemAlloc, peakMemSys atomic.Uint64
	go func() {
		stats := new(runtime.MemStats)
		for {
			runtime.ReadMemStats(stats)
			if peakMemAlloc.Load() < stats.Alloc {
				peakMemAlloc.Store(stats.Alloc)
			}
			if peakMemSys.Load() < stats.Sys {
				peakMemSys.Store(stats.Sys)
			}
			time.Sleep(5 * time.Second)
		}
	}()
	// Import the chain
	start := time.Now()

	var importErr error

	if ctx.Args().Len() == 1 {
		if err := utils.ImportChain(chain, ctx.Args().First()); err != nil {
			importErr = err
			log.Error("Import error", "err", err)
		}
	} else {
		for _, arg := range ctx.Args().Slice() {
			if err := utils.ImportChain(chain, arg); err != nil {
				importErr = err
				log.Error("Import error", "file", arg, "err", err)
				if err == utils.ErrImportInterrupted {
					break
				}
			}
		}
	}
	chain.Stop()
	fmt.Printf("Import done in %v.\n\n", time.Since(start))

	// Output pre-compaction stats mostly to see the import trashing
	showDBStats(db)

	// Print the memory statistics used by the importing
	mem := new(runtime.MemStats)
	runtime.ReadMemStats(mem)

	fmt.Printf("Object memory: %.3f MB current, %.3f MB peak\n", float64(mem.Alloc)/1024/1024, float64(peakMemAlloc.Load())/1024/1024)
	fmt.Printf("System memory: %.3f MB current, %.3f MB peak\n", float64(mem.Sys)/1024/1024, float64(peakMemSys.Load())/1024/1024)
	fmt.Printf("Allocations:   %.3f million\n", float64(mem.Mallocs)/1000000)
	fmt.Printf("GC pause:      %v\n\n", time.Duration(mem.PauseTotalNs))

	if ctx.Bool(utils.NoCompactionFlag.Name) {
		return nil
	}

	// Compact the entire database to more accurately measure disk io and print the stats
	start = time.Now()
	fmt.Println("Compacting entire database...")
	if err := db.Compact(nil, nil); err != nil {
		utils.Fatalf("Compaction failed: %v", err)
	}
	fmt.Printf("Compaction done in %v.\n\n", time.Since(start))

	showDBStats(db)
	return importErr
}

func exportChain(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		utils.Fatalf("This command requires an argument.")
	}

	stack, _ := makeConfigNode(ctx)
	defer stack.Close()

	chain, db := utils.MakeChain(ctx, stack, true)
	defer db.Close()
	start := time.Now()

	var err error
	fp := ctx.Args().First()
	if ctx.Args().Len() < 3 {
		err = utils.ExportChain(chain, fp)
	} else {
		// This can be improved to allow for numbers larger than 9223372036854775807
		first, ferr := strconv.ParseInt(ctx.Args().Get(1), 10, 64)
		last, lerr := strconv.ParseInt(ctx.Args().Get(2), 10, 64)
		if ferr != nil || lerr != nil {
			utils.Fatalf("Export error in parsing parameters: block number not an integer\n")
		}
		if first < 0 || last < 0 {
			utils.Fatalf("Export error: block number must be greater than 0\n")
		}
		if head := chain.CurrentSnapBlock(); uint64(last) > head.Number.Uint64() {
			utils.Fatalf("Export error: block number %d larger than head block %d\n", uint64(last), head.Number.Uint64())
		}
		err = utils.ExportAppendChain(chain, fp, uint64(first), uint64(last))
	}
	if err != nil {
		utils.Fatalf("Export error: %v\n", err)
	}
	fmt.Printf("Export done in %v\n", time.Since(start))
	return nil
}

func importHistory(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		utils.Fatalf("usage: %s", ctx.Command.ArgsUsage)
	}

	stack, _ := makeConfigNode(ctx)
	defer stack.Close()

	chain, db := utils.MakeChain(ctx, stack, false)
	defer db.Close()

	var (
		start   = time.Now()
		dir     = ctx.Args().Get(0)
		network string
	)

	// Determine network.
	if utils.IsNetworkPreset(ctx) {
		switch {
		// ##CROSS: config
		case ctx.Bool(utils.OneFlag.Name):
			network = "cross"
		case ctx.Bool(utils.ZoneZeroFlag.Name):
			network = "zonezero"
		case ctx.Bool(utils.OneDev3Flag.Name):
			network = "crossdev3"
		case ctx.Bool(utils.OneDevFlag.Name):
			network = "crossdev"
		// ##
		case ctx.Bool(utils.MainnetFlag.Name):
			network = "mainnet"
		case ctx.Bool(utils.SepoliaFlag.Name):
			network = "sepolia"
		case ctx.Bool(utils.HoleskyFlag.Name):
			network = "holesky"
		case ctx.Bool(utils.HoodiFlag.Name):
			network = "hoodi"
		}
	} else {
		// No network flag set, try to determine network based on files
		// present in directory.
		var networks []string
		for _, n := range params.NetworkNames {
			entries, err := era.ReadDir(dir, n)
			if err != nil {
				return fmt.Errorf("error reading %s: %w", dir, err)
			}
			if len(entries) > 0 {
				networks = append(networks, n)
			}
		}
		if len(networks) == 0 {
			return fmt.Errorf("no era1 files found in %s", dir)
		}
		if len(networks) > 1 {
			return errors.New("multiple networks found, use a network flag to specify desired network")
		}
		network = networks[0]
	}

	if err := utils.ImportHistory(chain, dir, network); err != nil {
		return err
	}
	fmt.Printf("Import done in %v\n", time.Since(start))
	return nil
}

// exportHistory exports chain history in Era archives at a specified
// directory.
func exportHistory(ctx *cli.Context) error {
	if ctx.Args().Len() != 3 {
		utils.Fatalf("usage: %s", ctx.Command.ArgsUsage)
	}

	stack, _ := makeConfigNode(ctx)
	defer stack.Close()

	chain, _ := utils.MakeChain(ctx, stack, true)
	start := time.Now()

	var (
		dir         = ctx.Args().Get(0)
		first, ferr = strconv.ParseInt(ctx.Args().Get(1), 10, 64)
		last, lerr  = strconv.ParseInt(ctx.Args().Get(2), 10, 64)
	)
	if ferr != nil || lerr != nil {
		utils.Fatalf("Export error in parsing parameters: block number not an integer\n")
	}
	if first < 0 || last < 0 {
		utils.Fatalf("Export error: block number must be greater than 0\n")
	}
	if head := chain.CurrentSnapBlock(); uint64(last) > head.Number.Uint64() {
		utils.Fatalf("Export error: block number %d larger than head block %d\n", uint64(last), head.Number.Uint64())
	}
	err := utils.ExportHistory(chain, dir, uint64(first), uint64(last), uint64(era.MaxEra1Size))
	if err != nil {
		utils.Fatalf("Export error: %v\n", err)
	}
	fmt.Printf("Export done in %v\n", time.Since(start))
	return nil
}

// importPreimages imports preimage data from the specified file.
// it is deprecated, and the export function has been removed, but
// the import function is kept around for the time being so that
// older file formats can still be imported.
func importPreimages(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		utils.Fatalf("This command requires an argument.")
	}

	stack, _ := makeConfigNode(ctx)
	defer stack.Close()

	db := utils.MakeChainDatabase(ctx, stack, false)
	defer db.Close()
	start := time.Now()

	if err := utils.ImportPreimages(db, ctx.Args().First()); err != nil {
		utils.Fatalf("Import error: %v\n", err)
	}
	fmt.Printf("Import done in %v\n", time.Since(start))
	return nil
}

func parseDumpConfig(ctx *cli.Context, db ethdb.Database) (*state.DumpConfig, common.Hash, error) {
	var header *types.Header
	if ctx.NArg() > 1 {
		return nil, common.Hash{}, fmt.Errorf("expected 1 argument (number or hash), got %d", ctx.NArg())
	}
	if ctx.NArg() == 1 {
		arg := ctx.Args().First()
		if hashish(arg) {
			hash := common.HexToHash(arg)
			if number := rawdb.ReadHeaderNumber(db, hash); number != nil {
				header = rawdb.ReadHeader(db, hash, *number)
			} else {
				return nil, common.Hash{}, fmt.Errorf("block %x not found", hash)
			}
		} else {
			number, err := strconv.ParseUint(arg, 10, 64)
			if err != nil {
				return nil, common.Hash{}, err
			}
			if hash := rawdb.ReadCanonicalHash(db, number); hash != (common.Hash{}) {
				header = rawdb.ReadHeader(db, hash, number)
			} else {
				return nil, common.Hash{}, fmt.Errorf("header for block %d not found", number)
			}
		}
	} else {
		// Use latest
		header = rawdb.ReadHeadHeader(db)
	}
	if header == nil {
		return nil, common.Hash{}, errors.New("no head block found")
	}
	startArg := common.FromHex(ctx.String(utils.StartKeyFlag.Name))
	var start common.Hash
	switch len(startArg) {
	case 0: // common.Hash
	case 32:
		start = common.BytesToHash(startArg)
	case 20:
		start = crypto.Keccak256Hash(startArg)
		log.Info("Converting start-address to hash", "address", common.BytesToAddress(startArg), "hash", start.Hex())
	default:
		return nil, common.Hash{}, fmt.Errorf("invalid start argument: %x. 20 or 32 hex-encoded bytes required", startArg)
	}
	conf := &state.DumpConfig{
		SkipCode:          ctx.Bool(utils.ExcludeCodeFlag.Name),
		SkipStorage:       ctx.Bool(utils.ExcludeStorageFlag.Name),
		OnlyWithAddresses: !ctx.Bool(utils.IncludeIncompletesFlag.Name),
		Start:             start.Bytes(),
		Max:               ctx.Uint64(utils.DumpLimitFlag.Name),
	}
	log.Info("State dump configured", "block", header.Number, "hash", header.Hash().Hex(),
		"skipcode", conf.SkipCode, "skipstorage", conf.SkipStorage,
		"start", hexutil.Encode(conf.Start), "limit", conf.Max)
	return conf, header.Root, nil
}

func dump(ctx *cli.Context) error {
	stack, _ := makeConfigNode(ctx)
	defer stack.Close()

	db := utils.MakeChainDatabase(ctx, stack, true)
	defer db.Close()

	conf, root, err := parseDumpConfig(ctx, db)
	if err != nil {
		return err
	}
	triedb := utils.MakeTrieDatabase(ctx, db, true, true, false) // always enable preimage lookup
	defer triedb.Close()

	state, err := state.New(root, state.NewDatabase(triedb, nil))
	if err != nil {
		return err
	}
	if ctx.Bool(utils.IterativeOutputFlag.Name) {
		state.IterativeDump(conf, json.NewEncoder(os.Stdout))
	} else {
		fmt.Println(string(state.Dump(conf)))
	}
	return nil
}

// hashish returns true for strings that look like hashes.
func hashish(x string) bool {
	_, err := strconv.Atoi(x)
	return err != nil
}

func pruneHistory(ctx *cli.Context) error {
	stack, _ := makeConfigNode(ctx)
	defer stack.Close()

	// Open the chain database
	chain, chaindb := utils.MakeChain(ctx, stack, false)
	defer chaindb.Close()
	defer chain.Stop()

	// Determine the prune point. This will be the first PoS block.
	prunePoint, ok := history.PrunePoints[chain.Genesis().Hash()]
	if !ok || prunePoint == nil {
		return errors.New("prune point not found")
	}
	var (
		mergeBlock     = prunePoint.BlockNumber
		mergeBlockHash = prunePoint.BlockHash.Hex()
	)

	// Check we're far enough past merge to ensure all data is in freezer
	currentHeader := chain.CurrentHeader()
	if currentHeader == nil {
		return errors.New("current header not found")
	}
	if currentHeader.Number.Uint64() < mergeBlock+params.FullImmutabilityThreshold {
		return fmt.Errorf("chain not far enough past merge block, need %d more blocks",
			mergeBlock+params.FullImmutabilityThreshold-currentHeader.Number.Uint64())
	}

	// Double-check the prune block in db has the expected hash.
	hash := rawdb.ReadCanonicalHash(chaindb, mergeBlock)
	if hash != common.HexToHash(mergeBlockHash) {
		return fmt.Errorf("merge block hash mismatch: got %s, want %s", hash.Hex(), mergeBlockHash)
	}

	log.Info("Starting history pruning", "head", currentHeader.Number, "tail", mergeBlock, "tailHash", mergeBlockHash)
	start := time.Now()
	rawdb.PruneTransactionIndex(chaindb, mergeBlock)
	if _, err := chaindb.TruncateTail(mergeBlock); err != nil {
		return fmt.Errorf("failed to truncate ancient data: %v", err)
	}
	log.Info("History pruning completed", "tail", mergeBlock, "elapsed", common.PrettyDuration(time.Since(start)))

	// TODO(s1na): what if there is a crash between the two prune operations?

	return nil
}
