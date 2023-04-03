package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
	/**
	- Fix the issue mentioned by passing a copy of salutation into the closure
	- This ensures that when the goroutine is run, we refer to the correct string
	*/
}
