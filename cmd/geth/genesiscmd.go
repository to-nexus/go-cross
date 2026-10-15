// Copyright 2026 The go-ethereum Authors
// This file is part of go-ethereum.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/urfave/cli/v2"
)

// ##CROSS: make genesis

var (
	makeGenesisBreakpointFlag = &cli.BoolFlag{
		Name:  "breakpoint",
		Usage: "Add the Breakpoint system contracts, active from genesis",
	}
	makeGenesisBridgeFlag = &cli.BoolFlag{
		Name:  "bridge",
		Usage: "Add the CrossBridge proxy and implementation, initialized at genesis",
	}
	makeGenesisMulticallFlag = &cli.BoolFlag{
		Name:  "multicall",
		Usage: "Add Multicall3",
	}
	makeGenesisSingletonFactoryFlag = &cli.BoolFlag{
		Name:    "singleton-factory",
		Aliases: []string{"sf"},
		Usage:   "Add the deterministic deployment proxies",
	}
	makeGenesisAllFlag = &cli.BoolFlag{
		Name:  "all",
		Usage: "Add every contract above",
	}
	makeGenesisBridgeOwnerFlag = &cli.StringFlag{
		Name:  "bridge.owner",
		Usage: "Bridge admin address (default: the PoSA admin)",
	}
	makeGenesisBridgeDevFlag = &cli.StringFlag{
		Name:  "bridge.dev",
		Usage: "Bridge reward wallet address (default: the bridge owner)",
	}
	makeGenesisBridgeThresholdFlag = &cli.UintFlag{
		Name:  "bridge.threshold",
		Usage: "Number of bridge validator signatures required",
		Value: 1,
	}

	makeGenesisCommand = &cli.Command{
		Action:    makeGenesis,
		Name:      "make-genesis",
		Usage:     "Builds a custom genesis with initialized predeployed contracts",
		ArgsUsage: "[<inputGenesis> <outputGenesis>]",
		Flags: []cli.Flag{
			makeGenesisBreakpointFlag,
			makeGenesisBridgeFlag,
			makeGenesisMulticallFlag,
			makeGenesisSingletonFactoryFlag,
			makeGenesisAllFlag,
			makeGenesisBridgeOwnerFlag,
			makeGenesisBridgeDevFlag,
			makeGenesisBridgeThresholdFlag,
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
The make-genesis command adds the selected contracts to the input genesis,
and runs their initializers at genesis. CrossEx is always added.
Accounts in the input genesis keep their balances. Without arguments, it generates
the PoSA keys, builds the genesis from the CrossDev3 chain config, writes it to
genesis.json, and prints all keys. The generated genesis enables Cancun and
Prague from genesis, and Breakpoint and BreakpointBravo with --breakpoint.`,
	}
)

func makeGenesis(ctx *cli.Context) error {
	all := ctx.Bool(makeGenesisAllFlag.Name)
	opts := core.GenesisContracts{
		Breakpoint:       all || ctx.Bool(makeGenesisBreakpointFlag.Name),
		Bridge:           all || ctx.Bool(makeGenesisBridgeFlag.Name),
		Multicall:        all || ctx.Bool(makeGenesisMulticallFlag.Name),
		SingletonFactory: all || ctx.Bool(makeGenesisSingletonFactoryFlag.Name),
	}
	if opts.Bridge {
		var err error
		if opts.BridgeOwner, err = parseGenesisAddress(ctx, makeGenesisBridgeOwnerFlag); err != nil {
			return err
		}
		if opts.BridgeDev, err = parseGenesisAddress(ctx, makeGenesisBridgeDevFlag); err != nil {
			return err
		}
		threshold := ctx.Uint(makeGenesisBridgeThresholdFlag.Name)
		if threshold == 0 || threshold > math.MaxUint8 {
			return fmt.Errorf("--%s must be between 1 and %d", makeGenesisBridgeThresholdFlag.Name, math.MaxUint8)
		}
		opts.BridgeThreshold = uint8(threshold)
	} else {
		for _, flag := range []cli.Flag{makeGenesisBridgeOwnerFlag, makeGenesisBridgeDevFlag, makeGenesisBridgeThresholdFlag} {
			if ctx.IsSet(flag.Names()[0]) {
				return fmt.Errorf("--%s needs --%s", flag.Names()[0], makeGenesisBridgeFlag.Name)
			}
		}
	}

	var (
		genesis    *core.Genesis
		keys       *breakpointGenesisKeys
		outputPath string
	)
	switch ctx.Args().Len() {
	case 0:
		var err error
		if genesis, keys, err = makeDefaultGenesis(ctx.Uint("validators"), opts.Breakpoint); err != nil {
			return err
		}
		outputPath = "genesis.json"
	case 2:
		if ctx.IsSet("validators") || ctx.Bool("json") {
			return errors.New("--validators and --json are only available without input arguments")
		}
		input, err := os.Open(ctx.Args().Get(0))
		if err != nil {
			return fmt.Errorf("open input genesis: %w", err)
		}
		defer input.Close()
		genesis = new(core.Genesis)
		if err := json.NewDecoder(input).Decode(genesis); err != nil {
			return fmt.Errorf("decode input genesis: %w", err)
		}
		outputPath = ctx.Args().Get(1)
	default:
		return fmt.Errorf("usage: %s", ctx.Command.ArgsUsage)
	}
	if err := core.MakeGenesis(genesis, opts); err != nil {
		return fmt.Errorf("build genesis: %w", err)
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

// parseGenesisAddress reads an optional address flag. It returns the zero address if the flag is not set.
func parseGenesisAddress(ctx *cli.Context, flag *cli.StringFlag) (common.Address, error) {
	value := ctx.String(flag.Name)
	if value == "" {
		return common.Address{}, nil
	}
	if !common.IsHexAddress(value) {
		return common.Address{}, fmt.Errorf("invalid --%s address: %s", flag.Name, value)
	}
	return common.HexToAddress(value), nil
}
