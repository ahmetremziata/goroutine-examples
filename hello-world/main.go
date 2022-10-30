package main

import (
	"fmt"
	"time"
)

func printHelloWorld(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func main() {
	go printHelloWorld("Hello, world!")
	go printHelloWorld("Merhaba dünya!")
	go printHelloWorld("Hallo welt!")
	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine finished")
}
