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

var (
	CrossGenesisValidators = []common.Address{
		common.HexToAddress("0x02f5F938B40A60B2345Ac0C9Fd9D909bD1F012eE"),
		common.HexToAddress("0x091E0be1D78511dE3b6D7022C99539671d235D73"),
		common.HexToAddress("0x0A62e6a8553A10B2031fD956df6bc28852F4d91d"),
		common.HexToAddress("0x49bcC861395C42cA60d773FB053F25DDaACbeA50"),
		common.HexToAddress("0x60Eef089B72da83eC98330296b1E1Ff37C118588"),
		common.HexToAddress("0x6C79321D421E9a37041e1477929d478492a04608"),
		common.HexToAddress("0x7De4fa95e742D34f0979875c060f7e1886948f03"),
		common.HexToAddress("0x87EB0d108594FFC6DeB4dcBdB43E3B0899B7a863"),
		common.HexToAddress("0x884DCf7bE65a5287cFb9648283B90C70D0D59c8A"),
		common.HexToAddress("0x88B9B24C441f8775c01443F6Ce244cEF5dA0364A"),
		common.HexToAddress("0x90e456Fe1f07A56e891BA11C8dE045091363527D"),
		common.HexToAddress("0x95fa026C756874f12a499ef0f95531f0b93629F5"),
		common.HexToAddress("0xAB128537bBC770C8f9E9225c854a2170B708E88A"),
		common.HexToAddress("0xBb2182CE952F3501325cd77A7327f8De091b394c"),
		common.HexToAddress("0xbCb02c35e753b25c450727C3cb61F6Aa4D2D545F"),
		common.HexToAddress("0xc193C8F676aF724B1d2cCF1236200044E8c6C4a6"),
		common.HexToAddress("0xD6EE737Ff9EABda5f9E9365AfE7D032e2aEb208C"),
		common.HexToAddress("0xE6927A1181bf5760fC7e25ABC96f7C025fF73802"),
		common.HexToAddress("0xF405a387AA84e58c1A5D2150305AD1b1788646Cf"),
		common.HexToAddress("0xfa884b0C778Fa24D015481a262F3efe924dc0D61"),
	}
	ZoneZeroGenesisValidators  = extractValidators(ZoneZeroValidators)
	CrossDev3GenesisValidators = extractValidators(CrossDev3Validators)
	CrossDevGenesisValidators  = extractValidators(CrossDevValidators)
)

// ##

// ##CROSS: istanbul posa
var (
	CrossValidators = []PoSAValidator{
		{
			ID:        "cross01",
			Operator:  common.HexToAddress("0x94D0561Aed5F5A1a5a9341b8c4771990a7307DEe"),
			Validator: common.HexToAddress("0x49bcC861395C42cA60d773FB053F25DDaACbeA50"),
			Signer:    hexutil.MustDecode("0x80411d28c6afd2ba1fcb7ac5173753647a5bc8e9ade4c7f70793245f59b3499f08257c7c1e125724b7f1e36cd1f2419d"),
		},
		{
			ID:        "cross02",
			Operator:  common.HexToAddress("0x753EaeCB21280e21F66694876e1b7C3593F35BCB"),
			Validator: common.HexToAddress("0x95fa026C756874f12a499ef0f95531f0b93629F5"),
			Signer:    hexutil.MustDecode("0xb14a51f77d937f30031bbd3116790197bd4fe18dc1b86f62d74674a717d90b30219f9f09657cc3af330726832d4e0176"),
		},
		{
			ID:        "cross03",
			Operator:  common.HexToAddress("0x9df34662E8Edc7BfF0bB88e9c699c8380C400C13"),
			Validator: common.HexToAddress("0x6C79321D421E9a37041e1477929d478492a04608"),
			Signer:    hexutil.MustDecode("0x8c10ece6d8079f21b353e90523e89dcb648fc3bf8162423fec97051f409706309267bc526bdb9f818a08db0168c28658"),
		},
		{
			ID:        "cross04",
			Operator:  common.HexToAddress("0xa25F3f9fF1C2ACfbB517B162C0E19657f01D8935"),
			Validator: common.HexToAddress("0x0A62e6a8553A10B2031fD956df6bc28852F4d91d"),
			Signer:    hexutil.MustDecode("0x9403e85207956f27d5cc16889ef01dd9f85082a55ad39dad00a5301756feaa2e20f728f53db1277f8441aefeb28ca104"),
		},
		{
			ID:        "cross05",
			Operator:  common.HexToAddress("0x92634AE3b938A81F2683468B5F7036e7a059B268"),
			Validator: common.HexToAddress("0xBb2182CE952F3501325cd77A7327f8De091b394c"),
			Signer:    hexutil.MustDecode("0x81500101de03678a922aa0fb8006bd63143507d629eda1194fe84760a62ebc70d932e16972ee716ad34022af3e07ab21"),
		},
		{
			ID:        "cross06",
			Operator:  common.HexToAddress("0xAD3eA7358fdb635A31e323860540eE8CD576bAE1"),
			Validator: common.HexToAddress("0x376602e1d2b06a114593ef6f1c68ab9d4324829f"),
			Signer:    hexutil.MustDecode("0x93fda7124fb8831107461c932b650ea45413b0b3edf8a4e9766e9710e45f6fc3ee7e27bccaf26fa1b059663424c13fd7"),
		},
		{
			ID:        "cross07",
			Operator:  common.HexToAddress("0x54EF0f813b6420591580067Ea8CBF03D5fd3a1d7"),
			Validator: common.HexToAddress("0x884DCf7bE65a5287cFb9648283B90C70D0D59c8A"),
			Signer:    hexutil.MustDecode("0xb95155f06f8ab7285cde401267bb1e88aa953deedef8e1eb75ea37b7bd3caede4e9518897a0ea59b7a3c921ae20ebeff"),
		},
		{
			ID:        "cross08",
			Operator:  common.HexToAddress("0x5641d5E41BaDB37ba578CF4292f5146a42cF07c2"),
			Validator: common.HexToAddress("0xbCb02c35e753b25c450727C3cb61F6Aa4D2D545F"),
			Signer:    hexutil.MustDecode("0x814ceeaea4d63775e4f6ac201a4a23dc3fd3da09e91c36a1d9afaf94708db6355763fa9a7a90cdacf4cb12342b992dc5"),
		},
		{
			ID:        "cross09",
			Operator:  common.HexToAddress("0xaEE23649489146b2042009fEdB5329490C5392c1"),
			Validator: common.HexToAddress("0xc193C8F676aF724B1d2cCF1236200044E8c6C4a6"),
			Signer:    hexutil.MustDecode("0x890c1887d6898c87c50ce0f36e7b5252286a7ac62923ce31cd86b34681db86416c35d3c948ebfee435ecf4908ed8ecdf"),
		},
		{
			ID:        "cross10",
			Operator:  common.HexToAddress("0xc8dC3e55a8921773620681151Fa96E5049b06f91"),
			Validator: common.HexToAddress("0x87EB0d108594FFC6DeB4dcBdB43E3B0899B7a863"),
			Signer:    hexutil.MustDecode("0xb1b38bf2d0cf6149651c3203493ccfac0e4c275dfc023ef6341543f277b5b5e27290a9a0f9dc18131a5b4be7239bb60a"),
		},
		{
			ID:        "cross11",
			Operator:  common.HexToAddress("0x21EBA729DA7e564af67a6e8a61cD16b868Fd9690"),
			Validator: common.HexToAddress("0xF405a387AA84e58c1A5D2150305AD1b1788646Cf"),
			Signer:    hexutil.MustDecode("0xacabc7ca29904c339256d2aab2b322649b28d0d4f33052942d708dc43cc03ec4cdfb1a8603b94a494aed44f463bf6015"),
		},
		{
			ID:        "cross12",
			Operator:  common.HexToAddress("0xE686B56ac554d0446642BC692e5DBb500c2eA5Cc"),
			Validator: common.HexToAddress("0xD6EE737Ff9EABda5f9E9365AfE7D032e2aEb208C"),
			Signer:    hexutil.MustDecode("0x85a25fbb840f3903bfa79c837847102e6fc7e85e8f093cea4023ea993558f82340a3d3dbeb243fc061a777f3362d2d4b"),
		},
		{
			ID:        "cross13",
			Operator:  common.HexToAddress("0x88D203C9deD3bDe6087f4fC698d812fA54938F04"),
			Validator: common.HexToAddress("0x091E0be1D78511dE3b6D7022C99539671d235D73"),
			Signer:    hexutil.MustDecode("0x85e48e28c844e175c9fee0982a7cf6f68bbd4bd270249ea10fa2bbb02e0fa088d30b70dc4a7d631cd02e426672b831a8"),
		},
		{
			ID:        "cross14",
			Operator:  common.HexToAddress("0xfCd954658958e3CE264f9E87107AEC85DD3902b2"),
			Validator: common.HexToAddress("0xE6927A1181bf5760fC7e25ABC96f7C025fF73802"),
			Signer:    hexutil.MustDecode("0xab0bfcb37374b2bb272570c8bbe2350cd41f814f7fa79d6430a6b1be4e93dd9d49dd3bbb7ca710fd500bf31dfbf3e5a5"),
		},
		{
			ID:        "cross15",
			Operator:  common.HexToAddress("0x52604a5A595b2518feA4442A91C636d25AbA1AeD"),
			Validator: common.HexToAddress("0xfa884b0C778Fa24D015481a262F3efe924dc0D61"),
			Signer:    hexutil.MustDecode("0x83ccf5d515984e027aca01dfa85a3c727a2c3f46920b2c4ea52fc508447b090f424ab7961ba7c987d2b3028cdbef243e"),
		},
		{
			ID:        "cross16",
			Operator:  common.HexToAddress("0x510C45cD12d19C88814B618500a99f3f22101238"),
			Validator: common.HexToAddress("0xAB128537bBC770C8f9E9225c854a2170B708E88A"),
			Signer:    hexutil.MustDecode("0x82738a341e809e4e87ae2911af688b2eb04e1b200b59669fcbe508b5b7b587982e299341084fb8b610001f9731c15456"),
		},
		{
			ID:        "cross17",
			Operator:  common.HexToAddress("0x5fEDce337f71ef36bb7fC586993352B7592aA901"),
			Validator: common.HexToAddress("0x90e456Fe1f07A56e891BA11C8dE045091363527D"),
			Signer:    hexutil.MustDecode("0xa771b18628384b6af6b9dd9a9190d89ed7a78a7e86bb832ed2a589e95b7904121d2b89410c8f96a2637ba34f9c1257f1"),
		},
		{
			ID:        "cross18",
			Operator:  common.HexToAddress("0x8e7c0B9D506d082D78CaFAF3EFaE7DD41e7EFD22"),
			Validator: common.HexToAddress("0x02f5F938B40A60B2345Ac0C9Fd9D909bD1F012eE"),
			Signer:    hexutil.MustDecode("0xa0369af2b624df7a7a06e55e1b6546c3d4ac00ba82e1533cfea9cacecaf0e75cec2766956d3b032f1bd4848f5302fed8"),
		},
		{
			ID:        "cross19",
			Operator:  common.HexToAddress("0x7A2114A62078d8807a60451B1cCe9245e7F741C5"),
			Validator: common.HexToAddress("0x60Eef089B72da83eC98330296b1E1Ff37C118588"),
			Signer:    hexutil.MustDecode("0x948dbd32473391c59e80f060a792a046f4f1b03f4f8599f8f3412ba550780e8e221dcd6ced9b1a828195fb572eae79f7"),
		},
		{
			ID:        "cross20",
			Operator:  common.HexToAddress("0xEe83F8b215496d250eaFB9782a7f5FEe4019c88e"),
			Validator: common.HexToAddress("0x88B9B24C441f8775c01443F6Ce244cEF5dA0364A"),
			Signer:    hexutil.MustDecode("0xb8d588db8c6e893e0b2a5d1553817a4448dbcd0668942e4304310244ddcfad625a8f504e9a5f0e407ddd20e30fcf8951"),
		},
		{
			ID:        "cross21",
			Operator:  common.HexToAddress("0x9E1BD983F492D184e8a5cA5171967A8150D5826c"),
			Validator: common.HexToAddress("0x7De4fa95e742D34f0979875c060f7e1886948f03"),
			Signer:    hexutil.MustDecode("0xa4ba952751bcd3045688172ae4cb867f2037d7f8bf487f3f11f4166fc2fb2231c773c5806235a98ccae3cf35724d626f"),
		},
	}
	ZoneZeroValidators = []PoSAValidator{
		{
			ID:        "cross01",
			Operator:  common.HexToAddress("0x83dd61e67ebfb60bb4bc3f814d8095cb50217aaa"),
			Validator: common.HexToAddress("0x3AeE6025948c380cD0E4e71cBB041337cc0E2C4E"),
			Signer:    hexutil.MustDecode("0xa64af411e912aed7f7060fe83f666d31a01bffbb65ffc23ed3869c3f65fccd02a2b8ce1249f9494ad6f73dadc4ccaa86"),
		},
		{
			ID:        "cross02",
			Operator:  common.HexToAddress("0x51e9fadb56701d4822a662b84b12261752bbbd0a"),
			Validator: common.HexToAddress("0x2bCFC7F77555B585941Ddb04d9B7977f1Ae612CB"),
			Signer:    hexutil.MustDecode("0xa1bdecc1744d0130b43dcc7d3120de493727e4ca040ebc984baf305e689e47e93cbea42a842349e0501e11b8c3992b23"),
		},
		{
			ID:        "cross03",
			Operator:  common.HexToAddress("0x5044693656f8b4c0c613e475f24a3ac4e9f95b4d"),
			Validator: common.HexToAddress("0xfb2Ce154CF90F35851D3D3c2C378D9820dEd4eb5"),
			Signer:    hexutil.MustDecode("0x9205dcff62f1df29e2b4efd8a69d60125629fc686608e7e0d4e80da26b15e806a636ea9ce0714bb5569db611f3f18506"),
		},
		{
			ID:        "cross04",
			Operator:  common.HexToAddress("0x76760a878f7000485943d626384ce8bdb64c91cc"),
			Validator: common.HexToAddress("0xd3f69551F006075aA71351a387eE23DdeaA5DbF1"),
			Signer:    hexutil.MustDecode("0xb297b9ba8fecffff5e5399db7a49e6ae0e6f9f18e8718b5337db9db5c13f8b21f75ec88afba6f40b16127242a3ee44e1"),
		},
		{
			ID:        "cross05",
			Operator:  common.HexToAddress("0x1ff80eab19c61915ef0ec7752163d72e7274416b"),
			Validator: common.HexToAddress("0x53bDC4196454c15af773BF00D3b4bC1BC913B751"),
			Signer:    hexutil.MustDecode("0xb65cd9b1401a72ff382a5740a4171dcab7f74417497c4c5c5c627d7aef9489d16e655823a49a93a04be8c7a0f2961105"),
		},
		{
			ID:        "cross06",
			Operator:  common.HexToAddress("0xddd4f04dd932f4c2ba7b8d12e8b1136d2c7f850b"),
			Validator: common.HexToAddress("0x28a0Ef4a0cb8204a810615C9b01475006dA35ECD"),
			Signer:    hexutil.MustDecode("0x919ab85dd0f4e5a871f1ae922240b8ae81ca2fe0a5d51592833b23727b001c6de28fdb3c452af9e85a40b8af06b0fa5a"),
		},
		{
			ID:        "cross07",
			Operator:  common.HexToAddress("0xebb17c6cbe05507eb358c4be1305e179a721bfd0"),
			Validator: common.HexToAddress("0x88501713D1Bb34d81c4857C8F32905574D67Cd57"),
			Signer:    hexutil.MustDecode("0xb42ae214f042f3eaa1092378c3cff63d2fb412775f03f08e394e52712cd9077133d712bc47f665f13e6554c6e02f3d8b"),
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
	CrossPoSAConfig = &PoSAConfig{
		CouncilPeriod:        86400,
		ValidatorEpochLength: 300,
		DelegationPool:       common.HexToAddress("0x000000068122d9bB43B4BCF5497A10EdfA9F5E93"),
		Admin:                common.HexToAddress("0x22C1522276855B028c31a731BA10D125811Af37c"),
		RewardStartBlock:     big.NewInt(30748187),
		Validators:           CrossValidators,
	}

	ZoneZeroPoSAConfig = &PoSAConfig{
		CouncilPeriod:        86400,
		ValidatorEpochLength: 300,
		DelegationPool:       common.HexToAddress("0x0000000c5BB5C7B081237B5efaFC753D61201b94"),
		Admin:                common.HexToAddress("0x1b1154b25435bB1dE06C6FC8CA8Df79050ca33c1"),
		RewardStartBlock:     big.NewInt(26182221),
		Validators:           ZoneZeroValidators,
	}

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
