package ethunits

import (
	"errors"
	"math/big"
)

func ConvertFromWei(value *big.Int, unit Unit) (*big.Float, error) {
	floatValue := new(big.Float).SetInt(value)

	intValue, err := stringToInt(unit.Value())
	if err != nil {
		return nil, err
	}

	return floatValue.Quo(floatValue, new(big.Float).SetInt(intValue)), nil
}

func ConvertToWei(value *big.Float, unit Unit) (*big.Int, error) {
	if unit == Wei {
		return floatToInt(value), nil
	}

	intValue, err := stringToInt(unit.Value())
	if err != nil {
		return nil, err
	}

	result := new(big.Int)
	value.Mul(value, new(big.Float).SetInt(intValue)).Int(result)
	return result, nil
}

func floatToInt(value *big.Float) *big.Int {
	intVal := new(big.Float).SetInt(big.NewInt(1000000000000000000))
	result := new(big.Int)
	value.Mul(value, intVal).Int(result)
	return result
}

func stringToInt(value string) (*big.Int, error) {
	n := new(big.Int)
	n, ok := n.SetString(value, 10)
	if !ok {
		return nil, errors.New("SetString error")
	}

	return n, nil
}
