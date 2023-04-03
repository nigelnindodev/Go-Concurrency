package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})    //create condition with sync.Mutex as the locker
	queue := make([]interface{}, 0, 10) // create a slice with length of 0 and capacity 10

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:] // simulate a dequeue
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal() // let a goroutine waiting on a condition know that something has happened
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		/**
		- Check if the queue is filled up in a loop
		- Predicate is important to figure out if queue is accepting items or not
		NB: Think len(queue) > 1 would be a better predicate as it covers an unusual condition where the queue is somehow larger than 2
		*/
		for len(queue) == 2 {
			c.Wait() // suspend the main goroutine until a signal on the condition is sent
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}

	/**
	- The lock and unlock in the for loop is for enqueuing and the locks in the removeFromQueue function are self-explanatory.
	- The Cond type has two methods for notifying goroutines blocked on a wait call
		Signal: finds the goroutine waiting the longest and notifies that
		Broadcast: sends a signal to all goroutines that are waiting
	*/
}
