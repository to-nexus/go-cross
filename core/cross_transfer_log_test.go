package core

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/kylelemons/godebug/diff"
)

func TestAddTransferLog(t *testing.T) {
	var (
		unit    = int64(1_000_000_000_000_000_000) // 1 ether
		bigUnit = big.NewInt(unit)
		key1, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
		key2, _ = crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
		addr1   = crypto.PubkeyToAddress(key1.PublicKey)
		addr2   = crypto.PubkeyToAddress(key2.PublicKey)
		chainID = big.NewInt(612044)
		config  = &params.ChainConfig{ // config with crossway fork
			ChainID:                       chainID,
			HomesteadBlock:                big.NewInt(0),
			DAOForkBlock:                  big.NewInt(0),
			DAOForkSupport:                false,
			EIP150Block:                   big.NewInt(0),
			EIP155Block:                   big.NewInt(0),
			EIP158Block:                   big.NewInt(0),
			ByzantiumBlock:                big.NewInt(0),
			ConstantinopleBlock:           big.NewInt(0),
			PetersburgBlock:               big.NewInt(0),
			IstanbulBlock:                 big.NewInt(0),
			MuirGlacierBlock:              big.NewInt(0),
			BerlinBlock:                   big.NewInt(0),
			LondonBlock:                   big.NewInt(0),
			ArrowGlacierBlock:             big.NewInt(0),
			GrayGlacierBlock:              big.NewInt(0),
			TerminalTotalDifficulty:       common.Big0,
			TerminalTotalDifficultyPassed: true,
			ShanghaiTime:                  u64(0),
			CrosswayTime:                  u64(0),
		}
		gspec = &Genesis{
			Config: config,
			Alloc: types.GenesisAlloc{
				addr1: types.Account{
					Balance: big.NewInt(2 * unit),
					Nonce:   0,
				},
			},
		}
		gspecWithCrosstoken = &Genesis{
			Config: config,
			Alloc: types.GenesisAlloc{
				crosstokenAddress: types.Account{
					Code: crosstokenCode,
				},
				addr1: types.Account{
					Balance: big.NewInt(2 * unit),
					Nonce:   0,
				},
			},
		}
		genFunc = func(at int, value *big.Int) func(int, *BlockGen) {
			// block generator which will add a tx - will transfer value - at block(at)
			return func(i int, gen *BlockGen) {
				if i == at {
					gen.AddTx(
						types.MustSignNewTx(key1, gen.Signer(), &types.DynamicFeeTx{
							ChainID:   chainID,
							Nonce:     gen.TxNonce(addr1),
							To:        &addr2,
							Value:     value,
							GasTipCap: common.Big1,
							GasFeeCap: gen.BaseFee(),
							Gas:       100000,
						}),
					)
				}
			}
		}
	)

	configNoCross := *config
	configNoCross.CrosswayTime = nil

	for _, tt := range []struct {
		name      string
		gspec     *Genesis
		genFunc   func(int, *BlockGen)
		atBlock   int
		noReceipt bool
		noLog     bool
	}{
		{
			name:    "at genesis",
			gspec:   gspecWithCrosstoken,
			atBlock: 0,
		},
		{
			name:    "at block 3",
			gspec:   gspecWithCrosstoken,
			atBlock: 3,
		},
		{
			name:    "without contract",
			gspec:   gspec,
			atBlock: 1,
		},
		{
			name:      "zero value",
			gspec:     gspec,
			genFunc:   genFunc(2, common.Big0),
			noReceipt: true,
		},
		{
			name: "no crossway",
			gspec: &Genesis{
				Config: &configNoCross,
				Alloc: types.GenesisAlloc{
					addr1: types.Account{
						Balance: big.NewInt(2 * unit),
						Nonce:   0,
					},
				},
			},
			atBlock: 2,
			noLog:   true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			gen := tt.genFunc
			if gen == nil {
				gen = genFunc(tt.atBlock, bigUnit)
			}
			db, chain, receipts := GenerateChainWithGenesis(tt.gspec, beacon.NewFaker(), 5, gen)
			blockchain, _ := NewBlockChain(db, nil, tt.gspec, nil, beacon.NewFaker(), vm.Config{}, nil, nil)
			defer blockchain.Stop()

			if i, err := blockchain.InsertChain(chain); err != nil {
				t.Fatalf("insert error (block %d): %v", chain[i].NumberU64(), err)
				return
			}

			if tt.noReceipt {
				// recipient's balance is zero
				state, _ := blockchain.State()
				bal2 := state.GetBalance(addr2)
				if bal2.CmpBig(common.Big0) != 0 {
					t.Errorf("unexpeted balance: got %s, want %v", bal2.String(), 0)
				}

				// no receipt
				receipts0 := receipts[tt.atBlock]
				if len(receipts0) != 0 {
					t.Fatalf("invalid receipts count: got %v, want %v", len(receipts0), 0)
				}
			} else {
				// check recipient's balance
				state, _ := blockchain.State()
				bal2 := state.GetBalance(addr2)
				if bal2.CmpBig(bigUnit) != 0 {
					t.Errorf("unexpeted balance: got %s, want %s", bal2.String(), bigUnit.String())
				}

				// check the receipt
				receipts0 := receipts[tt.atBlock]
				if len(receipts0) != 1 {
					t.Fatalf("invalid receipts count: got %v, want %v", len(receipts0), 1)
				}
				receipt := receipts0[0]

				want := types.Receipt{
					Status: types.ReceiptStatusSuccessful,
				}
				if !tt.noLog {
					// check CrossTransfer event
					want.Logs = []*types.Log{
						{
							Address: crosstokenAddress,
							Topics: []common.Hash{
								transferLogSig,
								common.BytesToHash(addr1.Bytes()),
								common.BytesToHash(addr2.Bytes()),
							},
							Data: crossTransferData(bigUnit, nil, bigUnit),
						},
					}
				}
				crossCheckReceipt(t, receipt, &want)
			}
		})
	}
}

func crossTransferData(amount, balFrom, balTo *big.Int) (data []byte) {
	for _, v := range []*big.Int{amount, balFrom, balTo} {
		if v == nil {
			v = common.Big0
		}
		data = append(data, common.LeftPadBytes(v.Bytes(), 32)...)
	}
	return
}

func crossCheckReceipt(t *testing.T, got, want *types.Receipt) {
	// only check status and logs
	got1 := &types.Receipt{
		Status: got.Status,
	}
	for _, l := range got.Logs {
		got1.Logs = append(got1.Logs, &types.Log{
			Address: l.Address,
			Topics:  l.Topics,
			Data:    append(append(l.Data[:32], common.LeftPadBytes([]byte{}, 32)...), l.Data[64:]...), // ignore data[32:64]
		})
	}
	r1, err := json.MarshalIndent(got1, "", "  ")
	if err != nil {
		t.Fatal("error marshaling input receipts:", err)
	}

	r2, err := json.MarshalIndent(want, "", "  ")
	if err != nil {
		t.Fatal("error marshaling input receipts:", err)
	}

	d := diff.Diff(string(r1), string(r2))
	if d != "" {
		t.Error("unexpected receipt:", d)
	}
}
