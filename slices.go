package main

import "fmt"

func slices(){
	slice1 := make([]int, 3)
	fmt.Println(len(slice1))
	
	slice1[0] = 0
	slice1[1] = 1
	slice1[2] = 2
	
	fmt.Println(slice1)
	slice1 = append(slice1, 3)
	fmt.Println(slice1)

	copyOfSlice1 := make([]int, len(slice1))
	copy(copyOfSlice1, slice1)
	fmt.Println(copyOfSlice1)
	fmt.Println(slice1)
}