package main

import "fmt"

/*
*
  - We have an intStream <- chan int read channel in this example. A data pipeline input point (in this case (multiple, add))
    should be able to define different types as well, and as long as the data output of a channel conforms to its input, we
    can join them together.
*/
func main() {
	/**
	Think is a generator as the data initialization point to convert no channel data into channels os that it can be transformed.
		- This could be:
			- A database call
			- API requests
	*/
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multiplyStream := make(chan int)
		go func() {
			defer close(multiplyStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multiplyStream <- i * multiplier:
				}
			}
		}()
		return multiplyStream
	}

	add := func(
		done <-chan interface{},
		intStream <-chan int,
		additive int,
	) <-chan int {
		additiveStream := make(chan int)
		go func() {
			defer close(additiveStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case additiveStream <- i + additive:
				}
			}
		}()
		return additiveStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := add(done, multiply(done, intStream, 2), 1)
	for v := range pipeline {
		fmt.Println(v)
	}
}
