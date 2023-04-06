package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	repeatFn := func(
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

	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		defer close(done)
	}()

	randGen := func() interface{} { return rand.Int() }
	randomNums := repeatFn(done, randGen)

	for v := range randomNums {
		fmt.Println(v)
	}
}
