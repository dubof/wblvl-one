package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := new(big.Int)
	a.SetString("3000000", 10) 

	b := new(big.Int)
	b.SetString("5000000", 10) 

	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())
	fmt.Println("===========================")

	resultAdd := new(big.Int).Add(a, b)
	fmt.Printf("a + b = %s\n", resultAdd.String())

	resultSub := new(big.Int).Sub(a, b)
	fmt.Printf("a - b = %s\n", resultSub.String())

	resultMul := new(big.Int).Mul(a, b)
	fmt.Printf("a * b = %s\n", resultMul.String())

	resultDiv := new(big.Int).Div(a, b)
	fmt.Printf("a / b = %s\n", resultDiv.String())
}