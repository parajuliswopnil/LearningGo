package main

import ("fmt"
		"time"
)

func channelSync(done chan bool){
	fmt.Println("Starting >>>>>>")
	time.Sleep(time.Second)
	fmt.Println("Done")
	done <- true
}

func callChannelSync(){
	done := make(chan bool)
	go channelSync(done)

	<- done
}