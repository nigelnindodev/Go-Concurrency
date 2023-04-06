package helpers

import (
	"math"
)

// Take returns a new channel taking n items from the provided valueStream, then closes./**
func Take(
	done <-chan interface{},
	valueStream <-chan interface{},
	num int,
) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

// RepeatFn takes a function and repeatedly runs it until done is called/**
func RepeatFn(
	done <-chan interface{},
	fn func() interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func ToInt(
	done <-chan interface{},
	valueStream <-chan interface{},
) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for value := range valueStream {
			select {
			case <-done:
				return
			case intStream <- value.(int):
			}
		}
	}()
	return intStream
}

func PrimeFinder(
	done <-chan interface{},
	intStream <-chan int,
) <-chan interface{} {
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
		for value := range intStream {
			isNumberPrime := isPrime(value)
			if isNumberPrime {
				select {
				case <-done:
					return
				case primeStream <- value:
				}
			}
			//TODO: Why does adding an else here to terminate early in case we have a done but no prime number lead to a deadlock?
		}
	}()
	return primeStream
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}