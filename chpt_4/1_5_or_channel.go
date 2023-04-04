package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			/**
			- first recursive terminating criteria
			- return a nil channel if the slice is empty in the first place
			*/
			return nil
		case 1:
			/**
			- second recursive terminating criteria
			- if slice only has one element, return that element
			*/
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channels) {
			/**
			- Due to the recursion termination criteria, every recursive call to the or function will have at least two
				channels.Case 2 is an optimization to keep the number of goroutines constrained where there are calls to
				the function or with only two channels.
			*/
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			/**
			- Recursively creates an or channel from all channels in the slice after the third index, and then select
				from it.
			- TODO: learn destructuring syntax in Go
			*/
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()

		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))

	/**
	- This reminds me of some patterns in Scala (Akka), although much simpler to write really. An example:
		- You can use the or channel to call different competing services for information (the famous weather service
			example from Akka in Action) and close the slower call once the first one completes
	*/
}
