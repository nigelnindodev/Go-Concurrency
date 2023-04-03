package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
	/**
	- do only count the number of times the function passed to Do is called, not the number of invocations of different
	functions
	- so the result will be 1
	*/
}
