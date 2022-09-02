package main

import (
	"fmt"
	"math/big"
)

func main() {
	first := big.NewInt(1 << 21)
	second := big.NewInt(1 << 22)
	res := new(big.Int)

	res.Add(first, second)
	fmt.Printf("%s + %s = %s\n", first, second, res)

	res.Sub(first, second)
	fmt.Printf("%s - %s = %s\n", first, second, res)

	res.Div(first, second)
	fmt.Printf("%s / %s = %s\n", first, second, res)

	res.Mul(first, second)
	fmt.Printf("%s * %s = %s\n", first, second, res)
}
