package main

import (
	"fmt"
	"time"
)

func doTask(done chan bool) {
	fmt.Println("Start Task...")
	time.Sleep(2 * time.Second)
	fmt.Println("Finish task...")
	done <- true
}

func main() {
	done := make(chan bool)
	go doTask(done)
	<-done
}
