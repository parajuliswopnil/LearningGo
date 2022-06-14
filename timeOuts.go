package main

import "fmt"
import "time"

func timeOutsImplementation(){
	channel1 := make(chan string)

	go func ()  {
		displayQuestions()
		channel1 <- "display question"
	}()

	select{
	case channel1msg := <- channel1:
		fmt.Println(channel1msg)
	case <- time.After(5 * time.Second):
		fmt.Println("time is up")
	}
}

func displayQuestions(){
	
	for i:=0; i<5; i++{
		fmt.Println("Question: ", i)
		time.Sleep(50 * time.Second)
	}
}