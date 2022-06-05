package main

import "fmt"

func forloop(){
	fmt.Println("Learning for loop")

	// different methods to use for loop
	// method 1
	i := 0
	for i < 7 {
		fmt.Println(i)
		i += 1 
	}

	// method 2
	for i := 0; i < 7; i++{
		fmt.Println("method 2: ", i)
	}

	// method 3 
	for {
		fmt.Println("infinite loop?")
		break
	}

	// method 4: with conditions
	for i := 0; i < 7; i++{
		if i == 3{
			continue
		}
		if i==6 {
			break
		}
		fmt.Println("How many time did it execute?: ", i)
	}
}