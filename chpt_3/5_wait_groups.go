package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}
	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)

	for i := 0; i < 5; i++ {
		go hello(&wg, i)
	}

	wg.Wait()
	/**
	Since we don't know which order goroutines will run, we should be getting different ordering of numbers for
	different invocations
	*/
}
