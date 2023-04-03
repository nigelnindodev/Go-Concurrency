package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello Channels!"
	}()
	fmt.Println(<-stringStream)
	/**
	- We will see output to stdout because channels are blocking
	- This also means that a goroutine that attempts to write to channel that is full will wait until the channel is
	empty.Any that attempts to read from a channel that is empty will wait until at least one item is placed into it.
	*/
}
