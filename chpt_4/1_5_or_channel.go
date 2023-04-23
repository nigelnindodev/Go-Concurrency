package main

import (
	"fmt"
	"time"
)

/*
*
- The or channel takes in n channels and terminates after the first channel completes, while also closing other non-finished channels
- Terminating conditions explained:
  - len(channels) case 0:
  - We have no channels to wait for, so return nil
  - len(channels) case 1:
  - We have only one channel, so we don't have any other channel to compare to that will run the fastest
*/
func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		fmt.Printf("Running or function for %v channel(s)\n", len(channels))
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
			- Better not have an extremely long-running channel here!
			*/
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channels) {
			/**
			- Case 2 is a quick optimization. We only have two channels, so we can skip the recursive call for n channels.
			- TODO: You could potentially do away with the recursive call altogether if you knew how many channels you will have
			*/
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			/**
			- Recursively create a select for n channels
			- Another interesting thing here is if you arrange your channels according to the ones you think will
				terminate first, there is a performance benefit. An example:
					- When running the select below, case <-channels[0] hase already finished. No need to undergo
						recursion in such as case. Take this with a grain of salt though, because as seen before the
						select function tries to "balance out" invocations of channels.
			*/
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				/**
				- Also pass in the orDone channel during the recursive call so that when goroutines up the tree (recursive call stack) finish
					up (in this case the ones associated with channels[0..2]), it will also exit the goroutines down the
					tree (the ones added in channels[3:])
				- This will in essence lead to a creation of a new orDone channel for every iteration of the or function
					to handle closing of channels within a specific recursive step/stack
				*/
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()

		/**
		Don't forget about this single but important line!
			- This is what closes the or function returns, And is tied to the `defer close (orDone)`. In the main goroutine,
				once the close(orDone) is called, that's what will cause the termination of the main goroutine in this example.
			- The termination is caused by blocking the main goroutine by waiting for a read on the channel the or function returns until we have a result
		*/
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
	fmt.Printf("done after %v\n", time.Since(start))

	/**
	- This reminds me of some patterns in Scala (Akka), although much simpler to write really. An example:
		- You can use the or channel to call different competing services for information (the famous weather service
			example from Akka in Action) and close the slower call once the first one completes
	*/
}
