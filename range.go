package main

import "fmt"

func rangeLearn(){
	numberList := make([]int, 3)
	numberList[0] = 1
	numberList[1] = 2
	numberList[2] = 3

	for i, value := range numberList{
		if value == 5{
			fmt.Println("Found in index: ", i)
			break
		} 
		if i == len(numberList) - 1{
			fmt.Println("Not found")
		}
	}

	customerId := make(map[string]string)
	customerId["swopnil"] = "parajuli"
	customerId["randomGuy"] = "tooRandom"

	for key, value := range customerId{
		fmt.Println(key)
		fmt.Println(value)
		fmt.Printf("%s -> %s", key, value)
	}
}