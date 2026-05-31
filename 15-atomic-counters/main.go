// Demonstrates that atomic counters provide a way to
// update the same counter with many concurrent goroutines
// without data race failures.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const routines = 50
const increments = 1000

func main() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for range routines { // start 50 goroutines
		wg.Go(func() {
			for range increments { // each increments the ops counter 1000x
				ops.Add(1)
			}
		})
	}

	wg.Wait()
	fmt.Println(ops.Load())
}
