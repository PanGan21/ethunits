package main

import (
	"ethunits"
	"fmt"
	"math/big"
)

func main() {
	fromWei, err := ethunits.ConvertFromWei(big.NewInt(1500000000000000000), ethunits.Ether)
	if err != nil {
		panic(err)
	}
	fmt.Println(fromWei)

	toWei, err := ethunits.ConvertToWei(new(big.Float).SetFloat64(0.01), ethunits.Ether)
	if err != nil {
		panic(err)
	}
	fmt.Println(toWei)

	toWei2, err := ethunits.ConvertToWei(new(big.Float).SetFloat64(1.5), ethunits.Wei)
	if err != nil {
		panic(err)
	}

	fmt.Println(toWei2)
}
