// run with:
// # time go run main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

const numJobs = 5

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start workers
	wg := sync.WaitGroup{} // or: var wg sync.WaitGroup
	for id := range numJobs {
		wg.Go(func() {
			worker(id, jobs, results)
		})
	}

	// send jobs
	for i := range numJobs {
		jobs <- i
		fmt.Println("sent", i)
	}
	close(jobs) // NOTE: need to close the sender!

	wg.Wait()

	// receive results
	for range numJobs {
		fmt.Print(<-results, " ")
	}
}

// receive jobs, send results
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "job", j)
		time.Sleep(time.Second)
		results <- j * 3
	}
}
