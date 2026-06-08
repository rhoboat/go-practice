// Implement two kinds of rate limiting:
// 1. Use a regular limiter: ticker as a blocker
// 2. Use a bursty limiter: ticker with a buffer
package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	limiter := time.Tick(500 * time.Millisecond)

	for i := range 5 {
		requests <- i
	}
	close(requests)

	for req := range requests {
		<-limiter
		fmt.Print(req, " ")
	}
	fmt.Println()

	burstyRequests := make(chan int, 5)
	burstyLimiter := make(chan time.Time, 3)

	for i := range 5 {
		burstyRequests <- i
	}
	close(burstyRequests)

	for range 3 {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(500 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	for range burstyRequests {
		<-burstyLimiter
		fmt.Println(time.Now())
	}
}
