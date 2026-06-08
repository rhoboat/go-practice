package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	go func() {
		for {
			select {
			case m := <-messages:
				fmt.Println()
				fmt.Println("message", m)
			default:
				fmt.Print(". ")
			}
		}
	}()

	ticker := time.Tick(time.Millisecond)
	for i := range 10 {
		t := <-ticker
		messages <- fmt.Sprint(i, " ", t.String())
	}
}
