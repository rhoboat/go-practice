// to get total runtime measurements, run with:
// # time go run main.go
// output:
// go run main.go  0.21s user 0.17s system 14% cpu 2.538 total
package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		chan1 <- 1
	}()
	go func() {
		time.Sleep(time.Second * 2)
		chan2 <- 2
	}()

	// even if the following is not run,
	for range 2 { // without range 2, this will never end, leading to panic when goroutines are asleep (channels have no messages)
		select {
		case thingy := <-chan1:
			fmt.Println(thingy)
		case otherthingy := <-chan2:
			fmt.Println(otherthingy)
		}
	}
}
