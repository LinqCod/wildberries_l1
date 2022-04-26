package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("124500000000000000110000030010024", 10)
	b, _ := new(big.Int).SetString("234555168000000000000000000000000", 10)

	fmt.Printf("%v + %v = %v\n", a, b, Add(a, b))
	fmt.Printf("%v - %v = %v\n", a, b, Sub(a, b))
	fmt.Printf("%v * %v = %v\n", a, b, Mul(a, b))
	fmt.Printf("%v / %v = %v\n", a, b, Div(a, b))
}

func Add(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func Sub(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func Mul(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func Div(a *big.Int, b *big.Int) *big.Float {
	a1 := new(big.Float).SetInt(a)
	b1 := new(big.Float).SetInt(b)
	return new(big.Float).Quo(a1, b1)
}
