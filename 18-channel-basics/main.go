package main

import (
	"fmt"
)

func main() {

	// Nil channel
	var nilChan chan *int
	fmt.Println(nilChan) // prints <nil>
	// Receiving from a nil channel will block indefinitely
	// Uncomment the following line to see this (it will cause a deadlock)
	// fmt.Println(<-nilChan)

	// Empty, but unbuffered channel
	unbufChan := make(chan int)
	// Sending to an unbuffered channel without a receiver will block
	go func() {
		unbufChan <- 1 // Blocks until someone receives from the channel
	}()
	fmt.Println(<-unbufChan) // Receives

	// Empty, buffered channel
	bufChan := make(chan int, 1)
	// Sending to buffered channel without receiver won't block
	// unless limits are exceeded
	bufChan <- 2
	fmt.Println(<-bufChan) // Receives

	// Empty, but with pointer and value nil
	pointerChan := make(chan *int)
	fmt.Println(pointerChan) // print an address(the channel itself)
	go func() {
		var val *int
		pointerChan <- val //  you can send nil to a channel
	}()
	// Receiving a nil value from a non-nil channel is okay
	fmt.Println(<-pointerChan) // Will print "<nil>"

	// Closed channel
	closedChan := make(chan int, 1)
	closedChan <- 1
	close(closedChan)
	// Attempting to send to a closed channel will cause a panic
	// closedChan <- 1 // Uncomment to see panic

	// Receiving from an empty closed channel returns immediately with the zero value
	val, ok := <-closedChan
	if !ok {
		fmt.Println("Channel closed, received:", val)
	} else {
		fmt.Println("Received from closed channel:", val)
	}
}
