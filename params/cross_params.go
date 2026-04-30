package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ##CROSS: config
var (
	CrossGenesisHash     = common.HexToHash("0x675a5d7c97cc1be3bf061635d31206966ef7ae0205a45890a4bab931456a8348")
	ZoneZeroGenesisHash  = common.HexToHash("0xcd8ce15999597691045f5f9675debfe4e0a7e5dca94b866e01df6b5b9d4263aa")
	CrossDev3GenesisHash = common.HexToHash("0xda75bb3ff72575b0fdc746a7fd23ebb733c8ce83cc6a6710470d1cf5528eab3f")
	CrossDevGenesisHash  = common.HexToHash("0x43edb9ebbb2f946b92268c6934a2f89e4f2dc4249223f3ed61a320cb8dcc1b19")

	FoundationCross     = common.HexToAddress("0xb9e345ba27826d71eb89a015852c752e341010ec")
	FoundationZoneZero  = common.HexToAddress("0x06Dc63E28d18172A689213071884c66c5281b493")
	FoundationCrossDev3 = common.HexToAddress("0xB9032595eC0465f43de9CF68c1E230888a5d16b6")
	FoundationCrossDev  = common.HexToAddress("0xB9032595eC0465f43de9CF68c1E230888a5d16b6")

	BeneficiaryCross     = common.HexToAddress("0xeb569ffda1a757938187320866959438aa61f4c6")
	BeneficiaryZoneZero  = common.HexToAddress("0x579c60A3176C5B588aeAD61a1F878a6A19CCc84E")
	BeneficiaryCrossDev3 = common.HexToAddress("0xB9032595eC0465f43de9CF68c1E230888a5d16b6")
	BeneficiaryCrossDev  = common.HexToAddress("0xB9032595eC0465f43de9CF68c1E230888a5d16b6")

	EcoCross = common.HexToAddress("0x0575a1b8e9e8950356b0c682bb270e16905eb108")
)

// ##CROSS: istanbul posa
var (
	CrossValidators = []PoSAValidator{
		{
			Validator: common.HexToAddress("0x02f5F938B40A60B2345Ac0C9Fd9D909bD1F012eE"),
		},
		{
			Validator: common.HexToAddress("0x091E0be1D78511dE3b6D7022C99539671d235D73"),
		},
		{
			Validator: common.HexToAddress("0x0A62e6a8553A10B2031fD956df6bc28852F4d91d"),
		},
		{
			Validator: common.HexToAddress("0x49bcC861395C42cA60d773FB053F25DDaACbeA50"),
		},
		{
			Validator: common.HexToAddress("0x60Eef089B72da83eC98330296b1E1Ff37C118588"),
		},
		{
			Validator: common.HexToAddress("0x6C79321D421E9a37041e1477929d478492a04608"),
		},
		{
			Validator: common.HexToAddress("0x7De4fa95e742D34f0979875c060f7e1886948f03"),
		},
		{
			Validator: common.HexToAddress("0x87EB0d108594FFC6DeB4dcBdB43E3B0899B7a863"),
		},
		{
			Validator: common.HexToAddress("0x884DCf7bE65a5287cFb9648283B90C70D0D59c8A"),
		},
		{
			Validator: common.HexToAddress("0x88B9B24C441f8775c01443F6Ce244cEF5dA0364A"),
		},
		{
			Validator: common.HexToAddress("0x90e456Fe1f07A56e891BA11C8dE045091363527D"),
		},
		{
			Validator: common.HexToAddress("0x95fa026C756874f12a499ef0f95531f0b93629F5"),
		},
		{
			Validator: common.HexToAddress("0xAB128537bBC770C8f9E9225c854a2170B708E88A"),
		},
		{
			Validator: common.HexToAddress("0xBb2182CE952F3501325cd77A7327f8De091b394c"),
		},
		{
			Validator: common.HexToAddress("0xbCb02c35e753b25c450727C3cb61F6Aa4D2D545F"),
		},
		{
			Validator: common.HexToAddress("0xc193C8F676aF724B1d2cCF1236200044E8c6C4a6"),
		},
		{
			Validator: common.HexToAddress("0xD6EE737Ff9EABda5f9E9365AfE7D032e2aEb208C"),
		},
		{
			Validator: common.HexToAddress("0xE6927A1181bf5760fC7e25ABC96f7C025fF73802"),
		},
		{
			Validator: common.HexToAddress("0xF405a387AA84e58c1A5D2150305AD1b1788646Cf"),
		},
		{
			Validator: common.HexToAddress("0xfa884b0C778Fa24D015481a262F3efe924dc0D61"),
		},
	}
	ZoneZeroValidators = []PoSAValidator{
		{
			Validator: common.HexToAddress("0x3AeE6025948c380cD0E4e71cBB041337cc0E2C4E"),
		},
		{
			Validator: common.HexToAddress("0x2bCFC7F77555B585941Ddb04d9B7977f1Ae612CB"),
		},
		{
			Validator: common.HexToAddress("0xfb2Ce154CF90F35851D3D3c2C378D9820dEd4eb5"),
		},
		{
			Validator: common.HexToAddress("0xd3f69551F006075aA71351a387eE23DdeaA5DbF1"),
		},
		{
			Validator: common.HexToAddress("0x28a0Ef4a0cb8204a810615C9b01475006dA35ECD"),
		},
		{
			Validator: common.HexToAddress("0x53bDC4196454c15af773BF00D3b4bC1BC913B751"),
		},
		{
			Validator: common.HexToAddress("0x88501713D1Bb34d81c4857C8F32905574D67Cd57"),
		},
	}
	CrossDev3Validators = []PoSAValidator{
		{
			ID:        "apne2cn01",
			Operator:  common.HexToAddress("0xCBCA857E24C6E1141c1039d54Bd395ec84dB65d7"),
			Validator: common.HexToAddress("0x415B1312a4aDC370eb791Fd0dB6086D5059b746A"),
			Signer:    hexutil.MustDecode("0xa75481e47f8daf60860ad87590f51575507422744c6f5e136df72ebb8cf5084afbe59b8cd4d09f71bc2f04a3fd678b66"),
		},
		{
			ID:        "apne2cn02",
			Operator:  common.HexToAddress("0x2BeA13D18188976e5e2B3F324AB0F928e3eA0C58"),
			Validator: common.HexToAddress("0x8c04752f2b5b3A541b5709A095887ecB2a815F85"),
			Signer:    hexutil.MustDecode("0xae82f20b81364837a6f4f86791c792c712f289d01cb956599c50dbd0e047161c9de33bdcbf1fa68ca5b5b30f0b7f7f62"),
		},
		{
			ID:        "apne2cn03",
			Operator:  common.HexToAddress("0xbEeCBA9C3A86F934b42FE0E8DFFeA90211eae953"),
			Validator: common.HexToAddress("0x17afdD710Ecd39435eFc693c8Fadc9B8411b8a23"),
			Signer:    hexutil.MustDecode("0x89e6ad30b4ce516057f9c7718d2dcc81742a4a4ab8425b8f8cdb7f80a088f7829aff1622b2f1d1a5bd4d3948d65a6387"),
		},
	}
	CrossDevValidators = CrossDev3Validators[:1]
)

func extractValidators(validators []PoSAValidator) []common.Address {
	addresses := make([]common.Address, 0, len(validators))
	for _, validator := range validators {
		addresses = append(addresses, validator.Validator)
	}
	return addresses
}

var (
	CrossDev3PoSAConfig = &PoSAConfig{
		CouncilPeriod:        86400,
		ValidatorEpochLength: 300,
		DelegationPool:       common.HexToAddress("0x000000af71b874fa018c65A89C9cBe16C5831253"),
		Admin:                common.HexToAddress("0x901eb0353ca4E815e94EeF2a66F12325ea3Ac859"),
		RewardStartBlock:     big.NewInt(6690300),
		Validators:           CrossDev3Validators,
	}

	CrossDevPoSAConfig = &PoSAConfig{
		CouncilPeriod:        86400,
		ValidatorEpochLength: 300,
		DelegationPool:       common.HexToAddress("0x000000af71b874fa018c65A89C9cBe16C5831253"),
		Admin:                common.HexToAddress("0x901eb0353ca4E815e94EeF2a66F12325ea3Ac859"),
		RewardStartBlock:     big.NewInt(6690300),
		Validators:           CrossDevValidators,
	}
)

// ##
