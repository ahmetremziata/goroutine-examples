package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	//No goroutine cannot access this resource
	m.Lock()
	msg = s
	//After that we set free this resource
	m.Unlock()
}

func main() {
	msg = "Hello, world!"

	var mutex sync.Mutex

	wg.Add(2)
	// Race condition when two goroutines access same resources
	go updateMessage("Hello universe!", &mutex)
	go updateMessage("Hello cosmos!", &mutex)

	wg.Wait()
	fmt.Println(msg)
	//It can be cosmos or universe as we do not know which goroutine access source last
}
