package utils

import "github.com/ethereum/go-ethereum/core/types"

type ApplyExtra func(*types.IstanbulExtra) error

func Combine(applies ...ApplyExtra) ApplyExtra {
	return func(extra *types.IstanbulExtra) error {
		for _, apply := range applies {
			err := apply(extra)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func ApplyHeaderIstanbulExtra(header *types.Header, applies ...ApplyExtra) error {
	extra, err := GetExtra(header)
	if err != nil {
		return err
	}

	err = Combine(applies...)(extra)
	if err != nil {
		return err
	}

	return SetExtra(header, extra)
}
