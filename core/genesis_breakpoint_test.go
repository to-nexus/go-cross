package core_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMakeBreakpointGenesis(t *testing.T) {
	config := *params.CrossDev3ChainConfig
	istanbul := *config.Istanbul
	posa := *istanbul.PoSA
	config.Istanbul = &istanbul
	config.Istanbul.PoSA = &posa
	posa.Validators = append([]params.PoSAValidator(nil), posa.Validators[:2]...)
	posa.DelegationPool = common.Address{}
	posa.RewardStartBlock = nil
	config.BreakpointTime = nil

	coinbase := common.HexToAddress("0x1234")
	reserveBalance := big.NewInt(12345)
	systemBalance := big.NewInt(7)
	genesis := &core.Genesis{
		Config:     &config,
		Coinbase:   coinbase,
		GasLimit:   30_000_000,
		Difficulty: big.NewInt(1),
		Alloc: types.GenesisAlloc{
			contracts.CrossReserveAddr: {Balance: reserveBalance},
			params.SystemAddress:       {Balance: systemBalance},
		},
	}
	require.NoError(t, core.MakeBreakpointGenesis(genesis))
	assert.Equal(t, coinbase, genesis.Coinbase)
	assert.Zero(t, *genesis.Config.BreakpointTime)
	assert.Equal(t, contracts.DelegationPoolAddr, posa.DelegationPool)
	assert.Equal(t, big.NewInt(1), posa.RewardStartBlock)
	assert.Equal(t, reserveBalance, genesis.Alloc[contracts.CrossReserveAddr].Balance)
	assert.Equal(t, systemBalance, genesis.Alloc[params.SystemAddress].Balance)

	backend := simulated.NewBackend(genesis.Alloc)
	defer backend.Close()
	validatorSet := breakpoint.NewValidatorSet()
	active, err := bind.Call(
		validatorSet.Instance(backend.Client(), contracts.ValidatorSetAddr),
		&bind.CallOpts{},
		validatorSet.PackGetActiveValidators(),
		validatorSet.UnpackGetActiveValidators,
	)
	require.NoError(t, err)
	assert.Len(t, active.ValidatorAddrs, 2)

	stakeHub := breakpoint.NewStakeHub()
	initialStake := new(big.Int).Mul(big.NewInt(1_000_000), big.NewInt(params.Ether))
	for _, validator := range posa.Validators {
		staked, err := bind.Call(
			stakeHub.Instance(backend.Client(), contracts.StakeHubAddr),
			&bind.CallOpts{},
			stakeHub.PackGetStakedAmount(validator.Operator),
			stakeHub.UnpackGetStakedAmount,
		)
		require.NoError(t, err)
		assert.Equal(t, initialStake, staked)
	}
	assert.Equal(t, new(big.Int).Mul(initialStake, big.NewInt(2)), genesis.Alloc[contracts.DelegationPoolAddr].Balance)
}
