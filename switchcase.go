package main

import (
	"fmt"
	"time"
)

func switchCase(){
	i := 0
	switch i{
	case 0:
		fmt.Println("Switch case zero")
	case 1:
		fmt.Println("Swith case one")
	}
	day := time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("Not weekend")
	}

	checkValue := func(i interface{})  {
		switch i.(type) {
		case bool:
			fmt.Println("Boolean case")
		case int:
			fmt.Println("Integer case")
		case float32:
			fmt.Println("Float case")
		default:
			fmt.Println("None")
		}	
	}
	checkValue(1)
	checkValue("Swopnil")
	checkValue(true)
	var floatValue float32 = 3.2
	checkValue(floatValue)
}