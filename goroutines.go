package main

import (
	"fmt"
	"time"
)

func printSomethingInLoop(value string){
	for i:= 0; i < 5; i++{
		fmt.Println(value,": ", i)
	}
}

func makeThreads(){
	go printSomethingInLoop("With go routine")
	time.Sleep(time.Second)
	printSomethingInLoop("something in loop")
}