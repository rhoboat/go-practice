package main

import (
	"fmt"
)

func main() {
	msgs := make(chan int, 5)
	done := make(chan bool)

	// concurrent receiver
	go func() {
		for {
			m, more := <-msgs
			if more {
				fmt.Println("received", m)
			} else {
				fmt.Println("received all")
				done <- true
				return
			}
		}
	}()

	// concurrent sends
	for i := range 5 {
		fmt.Println("sent", i)
		msgs <- i
	}
	fmt.Println("sent all")
	close(msgs) // must close the channel or it panics
	// sent exactly 5, received all 5

	<-done // need this to sync main with goroutine

	// saves the zero value of the channel type
	// same syntax can inspect closed channels
	zero, more := <-msgs
	fmt.Println("has more?", more, zero)
}
