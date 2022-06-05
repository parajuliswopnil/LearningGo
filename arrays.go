package main

import (
	"fmt"
)

func array(){
	fmt.Println("Learning about arrays")

	var numberList[5] int
	fmt.Println(numberList)

	for i := 0; i < len(numberList); i++{
		numberList[i] = i
	}
	fmt.Println(numberList)

	numberlist2 := [6]int{1,2,3,4,5}
	fmt.Println(numberlist2)
}