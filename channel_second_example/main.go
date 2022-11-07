package main

import "fmt"

func printSomething(mychnl chan string) {

	for v := 0; v < 3; v++ {
		mychnl <- "Trendyol"
	}
	close(mychnl)
}

func main() {

	// Creating a channel
	channel := make(chan string)

	// Calling Goroutine
	go printSomething(channel)

	// When the value of ok is set to true means the channel is open and it can send or receive data
	// When the value of ok is set to false means the channel is closed
	for {
		res, ok := <-channel
		if ok == false {
			fmt.Println("Channel is closed! ", ok)
			break
		}
		fmt.Println("Channel is already opened. ", res, ok)
	}
}
