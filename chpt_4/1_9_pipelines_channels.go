package main

import "fmt"

func main() {
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
