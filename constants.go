package main

import (
	"fmt"
	"math/big"
)

func constants(){
	fmt.Println("Learning about Constants.")
	var decimals = big.NewInt(1000000000000000000)
	var bigDecimalPlaceholder big.Int
	fmt.Println(decimals)

	var decimalsDivision = bigDecimalPlaceholder.Mul(decimals, decimals)
	fmt.Println("Multiplying the two big integers: ", decimalsDivision)

	const number = 2.0
	const secondnumber = 10 / number
	fmt.Println("constant operation: ", secondnumber)
}