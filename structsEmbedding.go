package main

import "fmt"

type CarInformation struct{
	companyName string
	manufacturedYear int
}

func (car CarInformation) getCarInformation() map[string] string{
	carInfo := make(map[string]string)
	carInfo["companyName"] = car.companyName
	carInfo["manufacturedYear"] = fmt.Sprint(car.manufacturedYear)
	return carInfo
}

type CarOwnerInfo struct{
	CarInformation
	licenceNumber int
}	

func makeCarOwner(){
	carOwner := CarOwnerInfo{CarInformation: CarInformation{companyName: "Mercedes", manufacturedYear: 2009}, licenceNumber: 2211}
	fmt.Println(carOwner.getCarInformation())

	type owner interface{
		getCarInformation() map[string]string
	}

	var carOwner2 owner = carOwner
	fmt.Println(carOwner2.getCarInformation())
}

