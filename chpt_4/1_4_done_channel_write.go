package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			/**
			Wouldn't it be better to add the <- done select code block? That was we don't have to add a time.Sleep at
			the end of the code
			*/
			defer fmt.Println("newRandStream closure exited")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

	// simulate some work
	time.Sleep(1 * time.Second)
}
