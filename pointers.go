package main

import "fmt"

func passByValue(value int) {
	value = 0
}

func passByReference(value *int) {
	*value = 0
}

func changeValue() {
	i := 1
	passByValue(i)
	fmt.Println("Does not change the value as the method passByValue makes another copy of i. Value of i is: ", i)

	passByReference(&i)
	fmt.Println("Changes the value of the variable i as we have passed its pointer and the value is changed at the memory location: Value is: ", i)
}
