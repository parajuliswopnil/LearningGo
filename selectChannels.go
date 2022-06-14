package main

import (
	"fmt"
	"time"
)

func selectChannelsAccordingToExecTime(){
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)
	startTime := time.Now().Second();
	go func(){
		time.Sleep(time.Second)
		channel1 <- "first message"
	}()

	go func ()  {
		time.Sleep(time.Second * time.Duration(5))
		channel2 <- "second messgae"
	}()

	go func ()  {
		time.Sleep(time.Second * time.Duration(10))	
		channel3 <- "third message"
	}()

	
	for i:=0; i< 3; i++{
		select{
		case msg1 := <- channel1:
			fmt.Println(msg1)
			sleep()
		case <- time.After(time.Second * time.Duration(5)):
			fmt.Println("second channel")
		case msg3 := <- channel3:
			fmt.Println("third channel", msg3)
		}
	}
	endTime := time.Now().Second()

	fmt.Println("total Execution time: ", endTime - startTime)

}

func sleep(){
	time.Sleep(time.Second * time.Duration(15))
}