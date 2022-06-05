package main

import "fmt"

func maps(){
	userInfo := make(map[string]int)
	userInfo["swopnil"] = 34
	userInfo["randomGuy"] = 35

	fmt.Println(userInfo)
	fmt.Println(userInfo["randomGuys"])

	value, ptrs := userInfo["swopnil"]
	fmt.Println(value)
	fmt.Println(ptrs)

	customerInfo := map[string]int{"swopnil": 1, "randomGuy": 2}
	fmt.Println(customerInfo)

	delete(customerInfo, "randomGuy")
	value, isPresent := customerInfo["randomGuy"]
	fmt.Println("value is: ", value, "is present?: ", isPresent)
}