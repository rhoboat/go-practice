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

// Implement a read/write state manager
type readop struct {
	key  int
	resp chan int // to send a result response
}

type riteop struct {
	key  int
	val  int
	resp chan bool // to send a success response
}

func main() {
	// main owns these data for counting operations
	// var readops, riteops int64
	var readops, riteops atomic.Int64

	// channels communicate across goroutines
	reads, rites := make(chan readop), make(chan riteop)

	// goroutine owns records
	go func() {
		records := map[int]int{}
		for {
			select {
			case r := <-reads:
				r.resp <- records[r.key]
			case w := <-rites:
				records[w.key] = w.val
				w.resp <- true
			}
		}
	}()

	// do concurrent reads
	for range 50 {
		go func() {
			for {
				read := readop{key: rand.Intn(5), resp: make(chan int)}
				reads <- read
				<-read.resp
				// atomic.AddInt64(&readops, 1)
				readops.Add(1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// do concurrent writes
	for range 10 {
		go func() {
			for {
				rite := riteop{key: rand.Intn(5), val: rand.Intn(100), resp: make(chan bool)}
				rites <- rite
				<-rite.resp
				// atomic.AddInt64(&riteops, 1)
				riteops.Add(1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	// show the counts
	// fmt.Println("reads", readops)
	// fmt.Println("writes", riteops)
	fmt.Println("reads", readops.Load())
	fmt.Println("writes", riteops.Load())
}
