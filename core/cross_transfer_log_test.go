package core

// ##CROSS: transfer log

import (
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/params/predeploy"
	"github.com/kylelemons/godebug/diff"
)

var (
	testUnit            = int64(1_000_000_000_000_000_000)
	testBigUnit         = big.NewInt(testUnit)
	testKey1, _         = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	testKey2, _         = crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	testAddr1           = crypto.PubkeyToAddress(testKey1.PublicKey)
	testAddr2           = crypto.PubkeyToAddress(testKey2.PublicKey)
	testChainID         = big.NewInt(612044)
	testAdventureConfig = &params.ChainConfig{ // config with Adventure fork
		ChainID:                 testChainID,
		HomesteadBlock:          big.NewInt(0),
		DAOForkBlock:            big.NewInt(0),
		DAOForkSupport:          false,
		EIP150Block:             big.NewInt(0),
		EIP155Block:             big.NewInt(0),
		EIP158Block:             big.NewInt(0),
		ByzantiumBlock:          big.NewInt(0),
		ConstantinopleBlock:     big.NewInt(0),
		PetersburgBlock:         big.NewInt(0),
		IstanbulBlock:           big.NewInt(0),
		MuirGlacierBlock:        big.NewInt(0),
		BerlinBlock:             big.NewInt(0),
		LondonBlock:             big.NewInt(0),
		ArrowGlacierBlock:       big.NewInt(0),
		GrayGlacierBlock:        big.NewInt(0),
		TerminalTotalDifficulty: common.Big1,
		ShanghaiTime:            u64(0),
		AdventureTime:           u64(0),
	}
)

func TestCrossTransferSig(t *testing.T) {
	if predeploy.CrossExAddr.String() != "0xfeed00000000000000000000000000000000C0DE" {
		t.Fatalf("CrossExAddr: got %s, want %s", predeploy.CrossExAddr.String(), "0xfeed00000000000000000000000000000000C0DE")
		t.Fatalf("CrossExAddr: got %s, want %s", predeploy.CrossExAddr.String(), "0xfeed00000000000000000000000000000000C0DE")
	}
	if transferLogSig.String() != "0x1b49dfd76419ac50d37de77c8afd5e57b6472c9ddd8399f88dc61343356a462b" {
		t.Fatalf("transferLogSig: got %s, want %s", transferLogSig.String(), "0x1b49dfd76419ac50d37de77c8afd5e57b6472c9ddd8399f88dc61343356a462b")
	}
}

func TestAddTransferLog_transferTx(t *testing.T) {
	var (
		gspec = &Genesis{
			Config: testAdventureConfig,
			Alloc: types.GenesisAlloc{
				testAddr1: types.Account{
					Balance: big.NewInt(2 * testUnit),
					Nonce:   0,
				},
			},
		}
		gspecWithCrossEx = &Genesis{
			Config: testAdventureConfig,
			Alloc: types.GenesisAlloc{
				predeploy.CrossExAddr: types.Account{
					Code: common.Hex2Bytes(predeploy.CrossExBinRuntime),
				},
				testAddr1: types.Account{
					Balance: big.NewInt(2 * testUnit),
					Nonce:   0,
				},
			},
		}
		genFunc = func(at int, value *big.Int) func(int, *BlockGen) {
			// block generator which will add a tx - will transfer value - at block(at)
			return func(i int, gen *BlockGen) {
				if i == at {
					addTx(gen, testChainID, testKey1, testAddr2, value, nil)
				}
			}
		}
	)

	configNoCross := *testAdventureConfig
	configNoCross.AdventureTime = nil

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
			gspec:   gspecWithCrossEx,
			atBlock: 0,
		},
		{
			name:    "at block 3",
			gspec:   gspecWithCrossEx,
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
			name: "no Adventure",
			gspec: &Genesis{
				Config: &configNoCross,
				Alloc: types.GenesisAlloc{
					testAddr1: types.Account{
						Balance: big.NewInt(2 * testUnit),
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
				gen = genFunc(tt.atBlock, testBigUnit)
			}
			db, chain, receipts := GenerateChainWithGenesis(tt.gspec, beacon.NewFaker(), 5, gen)
			blockchain, _ := NewBlockChain(db, nil, tt.gspec, nil, beacon.NewFaker(), vm.Config{}, nil)
			defer blockchain.Stop()

			if i, err := blockchain.InsertChain(chain); err != nil {
				t.Fatalf("insert error (block %d): %v", chain[i].NumberU64(), err)
				return
			}

			if tt.noReceipt {
				// recipient's balance is zero
				state, _ := blockchain.State()
				bal2 := state.GetBalance(testAddr2)
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
				bal2 := state.GetBalance(testAddr2)
				if bal2.CmpBig(testBigUnit) != 0 {
					t.Errorf("unexpeted balance: got %s, want %s", bal2.String(), testBigUnit.String())
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
							Address: predeploy.CrossExAddr,
							Topics: []common.Hash{
								transferLogSig,
								common.BytesToHash(testAddr1.Bytes()),
								common.BytesToHash(testAddr2.Bytes()),
							},
						},
					}
				}
				crossCheckReceipts(t, "transfer", types.Receipts{receipt}, types.Receipts{&want})
			}
		})
	}
}

func TestAddTransferLog_callContract(t *testing.T) {
	var (
		contract1 = common.HexToAddress("0x000000000000000000000000000000000000c002")
		contract2 = common.HexToAddress("0x000000000000000000000000000000000000C003")

		// function deposit(uint256 value) external payable
		// function transfer(address to, uint256 value) external payable
		abi1 = `[{"type":"function","name":"deposit","inputs":[{"name":"value","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"transfer","inputs":[{"name":"to","type":"address","internalType":"address"},{"name":"value","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"payable"}]`

		// contract1: event NativeTransfer(address indexed from, address indexed to, uint256 amount)
		nativeTransferSig = common.HexToHash("0xce8688f853ffa65c042b72302433c25d7a230c322caba0901587534b6551091d")
		// contract2: event Received(address indexed from, uint256 amount)
		receivedSig = common.HexToHash("0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874")

		gspec = &Genesis{
			Config: testAdventureConfig,
			Alloc: types.GenesisAlloc{
				predeploy.CrossExAddr: types.Account{
					Code: common.Hex2Bytes(predeploy.CrossExBinRuntime),
				},
				contract1: types.Account{
					Code: common.Hex2Bytes("608060405260043610610028575f3560e01c8063a9059cbb1461002c578063b6b55f2514610041575b5f80fd5b61003f61003a36600461024c565b610054565b005b61003f61004f366004610281565b61018b565b8034146100995760405162461bcd60e51b815260206004820152600e60248201526d0ecc2d8eaca40dad2e6dac2e8c6d60931b60448201526064015b60405180910390fd5b306001600160a01b038316036100e45760405162461bcd60e51b815260206004820152601060248201526f63616e27742073656e6420746f206d6560801b6044820152606401610090565b5f80836001600160a01b0316836040515f6040518083038185875af1925050503d805f811461012e576040519150601f19603f3d011682016040523d82523d5f602084013e610133565b606091505b50915091508161014557805160208201fd5b6040518381526001600160a01b0385169033907fce8688f853ffa65c042b72302433c25d7a230c322caba0901587534b6551091d9060200160405180910390a350505050565b8034146101cb5760405162461bcd60e51b815260206004820152600e60248201526d0ecc2d8eaca40dad2e6dac2e8c6d60931b6044820152606401610090565b6040516311f9fbc960e21b8152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000c0036001600160a01b0316906347e7ef249083906044015f604051808303818588803b158015610232575f80fd5b505af1158015610244573d5f803e3d5ffd5b505050505050565b5f806040838503121561025d575f80fd5b82356001600160a01b0381168114610273575f80fd5b946020939093013593505050565b5f60208284031215610291575f80fd5b503591905056fea26469706673582212208c7371b1cae9f00af71d95e8e4ad01ea51e74047cfaecfbc9f41d1b8b526e37164736f6c63430008180033"),
				},
				contract2: types.Account{
					Code: common.Hex2Bytes("608060405260043610601b575f3560e01c806347e7ef2414601f575b5f80fd5b602e602a36600460b9565b6030565b005b80341460735760405162461bcd60e51b815260206004820152600e60248201526d0ecc2d8eaca40dad2e6dac2e8c6d60931b604482015260640160405180910390fd5b816001600160a01b03167f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f885258748260405160ad91815260200190565b60405180910390a25050565b5f806040838503121560c9575f80fd5b82356001600160a01b038116811460de575f80fd5b94602093909301359350505056fea2646970667358221220a02e9b2fbb39fb4a0aaf8fd6827ca1cded4e491b8c59fcfb8efee60598cdc10064736f6c63430008180033"),
				},
				testAddr1: types.Account{
					Balance: big.NewInt(5 * testUnit),
					Nonce:   0,
				},
			},
		}
	)

	atBlock1 := 3
	atBlock2 := 4
	ab, err := abi.JSON(strings.NewReader(abi1))
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	db, chain, receipts := GenerateChainWithGenesis(gspec, beacon.NewFaker(), 5, func(i int, gen *BlockGen) {
		switch i {
		case atBlock1:
			// transfer(addr2, 1 ether)
			input, err := ab.Pack("transfer", testAddr2, testBigUnit)
			if err != nil {
				t.Fatalf("failed to pack: %v", err)
			}
			addTx(gen, testChainID, testKey1, contract1, testBigUnit, input)
		case atBlock2:
			// deposit(1 ether)
			input, err := ab.Pack("deposit", testBigUnit)
			if err != nil {
				t.Fatalf("failed to pack: %v", err)
			}
			addTx(gen, testChainID, testKey1, contract1, testBigUnit, input)
		}
	})
	blockchain, _ := NewBlockChain(db, nil, gspec, nil, beacon.NewFaker(), vm.Config{}, nil)
	defer blockchain.Stop()

	if i, err := blockchain.InsertChain(chain); err != nil {
		t.Fatalf("insert error (block %d): %v", chain[i].NumberU64(), err)
		return
	}

	// check balances
	state, _ := blockchain.State()
	{
		bal := state.GetBalance(testAddr2)
		if bal.CmpBig(testBigUnit) != 0 {
			t.Errorf("(%s) unexpeted balance: got %s, want %s", testAddr2.String(), bal.String(), testBigUnit.String())
		}
	}
	{
		bal := state.GetBalance(contract2)
		if bal.CmpBig(testBigUnit) != 0 {
			t.Errorf("(%s) unexpeted balance: got %s, want %s", contract2.String(), bal.String(), testBigUnit.String())
		}
	}

	// check the receipts at block 3
	{
		want := append(types.Receipts{}, &types.Receipt{
			Status: types.ReceiptStatusSuccessful,
			Logs: []*types.Log{
				{
					// transfer(addr2, 1 ether): CrossTransfer(CrossEx, addr1, contract1, ...)
					Address: predeploy.CrossExAddr,
					Topics: []common.Hash{
						transferLogSig,
						common.BytesToHash(testAddr1.Bytes()),
						common.BytesToHash(contract1.Bytes()),
					},
				},
				{
					// contract1 -> addr2: CrossTransfer(CrossEx, contract1, addr2, ...)
					Address: predeploy.CrossExAddr,
					Topics: []common.Hash{
						transferLogSig,
						common.BytesToHash(contract1.Bytes()),
						common.BytesToHash(testAddr2.Bytes()),
					},
				},
				{
					// in contract1: NativeTransfer(addr1, addr2, 1 ether)
					Address: contract1,
					Topics: []common.Hash{
						nativeTransferSig,
						common.BytesToHash(testAddr1.Bytes()),
						common.BytesToHash(testAddr2.Bytes()),
					},
				},
			},
		})
		crossCheckReceipts(t, "block 3", receipts[atBlock1], want)
	}
	// check the receipts at block 4
	{
		var want types.Receipts
		want = append(want, &types.Receipt{
			Status: types.ReceiptStatusSuccessful,
			Logs: []*types.Log{
				{
					// deposit(1 ether): CrossTransfer(CrossEx, addr1, contract1, ...)
					Address: predeploy.CrossExAddr,
					Topics: []common.Hash{
						transferLogSig,
						common.BytesToHash(testAddr1.Bytes()),
						common.BytesToHash(contract1.Bytes()),
					},
				},
				{
					// contract2.deposit(addr1, 1 ether): CrossTransfer(CrossEx, contract1, contract2, ...)
					Address: predeploy.CrossExAddr,
					Topics: []common.Hash{
						transferLogSig,
						common.BytesToHash(contract1.Bytes()),
						common.BytesToHash(contract2.Bytes()),
					},
				},
				{
					// in contract2: Received(addr1, 1 ether)
					Address: contract2,
					Topics: []common.Hash{
						receivedSig,
						common.BytesToHash(testAddr1.Bytes()),
					},
				},
			},
		})

		crossCheckReceipts(t, "block 4", receipts[atBlock2], want)
	}
}

func addTx(gen *BlockGen, chainID *big.Int, key1 *ecdsa.PrivateKey, to common.Address, value *big.Int, input []byte) {
	gen.AddTx(
		types.MustSignNewTx(key1, gen.Signer(), &types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     gen.TxNonce(crypto.PubkeyToAddress(key1.PublicKey)),
			To:        &to,
			Value:     value,
			Data:      input,
			GasTipCap: common.Big1,
			GasFeeCap: gen.BaseFee(),
			Gas:       100000,
		}),
	)
}

func crossCheckReceipts(t *testing.T, prefix string, got, want types.Receipts) {
	if got.Len() != want.Len() {
		t.Fatalf("[%v] invalid receipts count: got %v, want %v", prefix, got.Len(), want.Len())
	}

	for i := range got {
		// only check status and logs (address and topics)
		got1 := &types.Receipt{
			Status: got[i].Status,
		}
		for _, l := range got[i].Logs {
			got1.Logs = append(got1.Logs, &types.Log{
				Address: l.Address,
				Topics:  l.Topics,
			})
		}

		r1, err := json.MarshalIndent(got1, "", "  ")
		if err != nil {
			t.Fatalf("[%v] error marshaling input receipts: %s", prefix, err)
		}
		r2, err := json.MarshalIndent(want[i], "", "  ")
		if err != nil {
			t.Fatalf("[%v] error marshaling input receipts: %s", prefix, err)
		}

		d := diff.Diff(string(r1), string(r2))
		if d != "" {
			t.Fatalf("[%v] unexpected receipt: %s", prefix, d)
		}
	}
}
