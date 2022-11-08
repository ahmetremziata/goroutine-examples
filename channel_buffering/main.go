package main

import "fmt"

func main() {
	// Create a channel with 10 length
	channel := make(chan int, 10)

	for i := 0; i < 10; i++ {
		channel <- i
		fmt.Println(<-channel)
	}

	close(channel)
}
