package main

import (
	"fmt"
	"sync"
)

type Counters struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Counters) inc(key string) {
	c.mu.Lock() // lock it to _others_
	defer c.mu.Unlock()
	c.counters[key]++
}

func main() {
	c := Counters{
		counters: map[string]int{"a": 0, "b": 0},
	}

	wg := sync.WaitGroup{}

	// define the function here so it closes over c
	increment := func(key string, amount int) {
		for range amount {
			// c.counters[key] += 1 // panics on concurrent map writes
			c.inc(key)
		}
	}

	wg.Go(func() {
		increment("a", 1000)
	})

	wg.Go(func() {
		increment("a", 999)
	})

	wg.Go(func() {
		increment("b", 1000)
	})

	wg.Wait()
	fmt.Println(c.counters)
}
