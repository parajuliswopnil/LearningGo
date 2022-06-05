package main

import "fmt"

func variables(){
	fmt.Println("Learning about variables in go.")
	var name string = "Swopnil";

	var lastName string
	lastName = "Parajuli"

	age := 22

	var boolValue bool = false && true
	
	fmt.Println("", name, " ", lastName, "age: ", age);
	fmt.Println("Boolean operators result: ", boolValue)

}