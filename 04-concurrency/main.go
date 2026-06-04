package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"sync"
)

const w = "generic"

func main() {
	urls := []string{
		"https://go.dev/blog/intro-generics",
		"https://go.dev/blog/gofix",
	}

	fmt.Println("sequential")
	occurencesSeq(urls)
	fmt.Println("concurrent")
	occurencesCnc(urls)
}

// sequential version
func occurencesSeq(urls []string) []int {
	result := make([]int, len(urls))
	for i, url := range urls {
		body := getBody(url)
		c := countWord("generic", body)
		result[i] = c
		fmt.Printf("count: %d, word: %s, url: %s\n", c, w, url)
	}

	return result
}

// concurrent version
func occurencesCnc(urls []string) []int {
	wg := sync.WaitGroup{}
	result := make([]int, len(urls))

	for i, url := range urls {
		wg.Add(1)
		go func() {
			body := getBody(url)
			c := countWord("generic", body)
			result[i] = c
			fmt.Printf("count: %d, word: %s, url: %s\n", c, w, url)
			wg.Done()
		}()
	}
	wg.Wait()
	return result
}

func countWord(word string, body []byte) int {
	re := regexp.MustCompile(word)
	matches := re.FindAll(body, -1)
	// fmt.Printf("%q\n", matches)
	return len(matches)
}

func getBody(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return body
}
