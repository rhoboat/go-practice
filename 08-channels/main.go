package main

import (
	"fmt"
	"time"
)

func main() {
	// both sends and receives are blocking.
	// the following works only if the send is concurrent with the receive
	unbuffered := make(chan string)
	go func() { unbuffered <- "teengs" }()
	fmt.Println(<-unbuffered)

	// buffered channel because it has a size. this can be 1, as well
	messages := make(chan string, 2)
	// this doesn't have to be concurrent because of buffering
	messages <- "hajjo"
	messages <- "hew"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// Sync main() with the worker channel
	done := make(chan bool)
	go worker(done) // because unbuffered, need goroutine
	// buffered example:
	// done := make(chan bool, 1)
	// worker(done) // goroutine not required
	<-done // block main() until we receive a message from the done channel
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
