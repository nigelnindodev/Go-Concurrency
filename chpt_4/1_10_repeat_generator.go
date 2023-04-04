package main

import (
	"fmt"
	"time"
)

func main() {
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			// for ... for loop generates values into the value stream until the done signal is sent
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	/**
	Finish up example in book by writing a function that prints 1,2 repeatedly for 5 seconds and then terminates
	*/
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		defer close(done)
	}()

	repeatStream := repeat(done, 1, 2)
	for v := range repeatStream {
		fmt.Println(v)
	}
	/**
	NB: A use case for this would be to make an API request to an idempotent server x number of times within a given duration.
	*/
}
