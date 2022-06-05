package main

import "fmt"

func variadicSumFunction(numberList ...int){
	sum := 0
	for _, value := range numberList{
		sum += value
	}
	fmt.Println(sum)
}