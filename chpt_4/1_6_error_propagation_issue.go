package main

import (
	"fmt"
	"net/http"
)

func main() {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err) // printing because we cannot pass back errors!
					continue
				}
				select {
				case responses <- resp:
				case <-done:
					return
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done) // close after sending http requests

	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}
