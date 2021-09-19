# ethunits

Ethereum unit conversion module for personal usage

## Usage

```
fromWei, err := ethunits.ConvertFromWei(big.NewInt(1500000000000000000), ethunits.Ether)
if err != nil {
	panic(err)
}
fmt.Println(fromWei)

toWei, err := ethunits.ConvertToWei(new(big.Float).SetFloat64(1.5), ethunits.Wei)
if err != nil {
	panic(err)
}
fmt.Println(toWei)
```

## Supported units

Supported units can be found in `constants.go`
