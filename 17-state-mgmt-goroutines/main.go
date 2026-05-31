// Implement the previous example without wait groups or mutexes.
// Use only goroutines and channels.
// Each piece of data is owned by exactly one goroutine. (1:1 data:goroutine)
// Share memory with communication
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int // each read op needs its own channel for responses
}

type writeOp struct {
	key  int
	val  int
	resp chan bool // each write op needs its own channel for confirmation
}

func main() {
	// these are going to be updated by concurrent things
	var readCounter int64
	var writeCounter int64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	// stateful goroutine that manages state by listening to channels (like a database engine)
	go func() {
		state := map[int]int{} // close over some private state
		for {
			select {
			case r := <-reads:
				r.resp <- state[r.key]
			case w := <-writes:
				state[w.key] = w.val
				w.resp <- true
			}
		}
	}()

	for range 50 {
		go func() {
			for { // need to loop indefinitely to run as many times as possible until stopped
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp                      // read it (throw away the value read)
				atomic.AddInt64(&readCounter, 1) // atomically update the counter
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for range 10 {
		go func() {
			for { // need to loop indefinitely to run as many times as possible until stopped
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp                      // again throw away the value
				atomic.AddInt64(&writeCounter, 1) // atomically update the counter
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("reads:", readCounter)
	fmt.Println("writes:", writeCounter)
}
