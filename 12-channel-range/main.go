package main

import (
	"fmt"
)

func main() {
	msgs := make(chan string, 2)

	// non-concurrent sends to buffered channel
	msgs <- "thingy"
	msgs <- "teengy"
	close(msgs) // must close the channel because we're using range next, or it panics

	// range over channels only works when the channel has been terminated
	for m := range msgs {
		fmt.Println("received", m)
	}

	zero, more := <-msgs
	fmt.Println("has more?", more, zero)
}
