package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(terminated)
			for {
				/**
				- The select helps read values from strings channel. But since it's nil, it shouldn't print anythin. However,
				the done channel helps us close it. In this case, after one second and clean up resources.
				*/
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{}) // create a done channel to the parent to signal cancellation
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling doWork goroutine...")
		close(done) // send signal after on second of sleep
	}()

	// line below unblocks after the channel being read from closes
	<-terminated // block and wait for terminated channel created inside doWork function
	fmt.Println("Done")
}
