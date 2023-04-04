package main

import (
	"fmt"
	"net/http"
)

// Result specifies the entire expected result set from the check status goroutine. This allows the main goroutine to
// define logic about how it would like to handle errors/**
type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case results <- result:
				case <-done:
					return
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	errCount := 0
	urls := []string{"https://www.google.com", "https://badhost", "https://badhost", "https://badhost", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("Too many Errors, breaking!")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
