package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(completed)
			// This for range makes the defer statement to close the channel not run, because we have a nil channel!
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	fmt.Println("Done")

	/**
	- We leak here because we have passed in a nil channel, so no strings ever get written to the channel.
	- doWork never exits, also the completed channel never closes.
	*/
}
