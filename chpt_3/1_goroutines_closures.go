package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
	/**
	- The goroutine closes over the value salutation
	- The loop will probably exit before any of the goroutines start running
	- This will cause the value good day to be printed three times
	*/
}
