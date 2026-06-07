// timeouts set limits on fetching external resources
// channels and select implement timeouts in go
package main

import (
	"fmt"
	"time"
)

const timeout = 2
const lrf = 3

func main() {
	ch := make(chan int, 1) // buffered, so the send is nonblocking, which means we don't need goroutine
	go func() {
		time.Sleep(time.Second * lrf) // long running function
		ch <- 1
	}()

	// select implements timeout
	select {
	case thingy := <-ch:
		fmt.Println(thingy)
	case <-time.After(time.Second * timeout):
		fmt.Println("timed out")
	}
}
