package main

import "fmt"

type Car struct{
	companyName string
	price int
}

func makeCar(companyName string, price int) Car{
	car := Car{}
	car.companyName = companyName
	car.price = price
	return car
}

func structs(){
	fmt.Println(makeCar("Ferrari", 1000000))

	mercedes := Car{companyName: "Mercedes", price: 1000000}
	fmt.Println(mercedes)

	toyota := Car{"mclaren", 10000}
	fmt.Println(toyota)
}