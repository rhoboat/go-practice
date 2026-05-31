package main

import (
	"fmt"
	"time"
)

const numJobs = 5

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start workers
	// for id := range 3 { // if it's set lower, time go run main.go shows it takes longer
	for id := range numJobs {
		go worker(id, jobs, results)
	}

	// send jobs
	for i := range numJobs {
		jobs <- i
	}
	// close(jobs) // don't actually need to close

	// receive results
	for range numJobs {
		fmt.Print(<-results, " ")
	}
	// NOTE: Why doesn't this work?
	// Range over channels only works if the channel is terminated
	// We can't call close(results) because workers are still working at this moment.
	// So the best way to receive results is <-results
	// for r := range results {
	// 	fmt.Print(r, " ")
	// }
}

// receive jobs, send results
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "job", j)
		time.Sleep(time.Second)
		results <- j * 3
	}
}
