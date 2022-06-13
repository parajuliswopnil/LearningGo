package main

import "fmt"

func channelBuffering(){
	message := make(chan string, 2)

	message <- "buffering"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}