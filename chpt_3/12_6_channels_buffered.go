package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4) // look at the result if channel is not buffered!
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done (intStream channel closed).")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}

	/**
	- If a goroutine making writes to a channel has knowledge of how many writes to be made,it can be useful to create a
	buffered channel whose capacity is the number of writes to be made, then make the writes as quickly as possible as
	the writes will not block (while waiting for a read off the channel).
	- In the case above, this allows us to quickly write 4 items to the channel before any reads can begin.
	*/
}
