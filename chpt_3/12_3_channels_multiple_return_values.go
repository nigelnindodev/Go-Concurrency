package main

import "fmt"

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello Channels!"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v", ok, salutation)
	/**
	- The second return value is a way for a read operation to indicate whether the read off the channel was a value
	generated by a write elsewhere in the process, or a default value generated from a closed channel.
	*/
}
