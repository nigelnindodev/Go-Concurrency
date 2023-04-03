package main

import (
	"sync"
	"testing"
)

func benchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}

	sender := func() {
		defer wg.Done()
		/**
		- Line below waits until we are told to begin
		- This eliminates the time factor of creating a goroutine affecting measurement of context switching
		*/
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token // send an empty struct to the receiver goroutine
		}
	}

	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c // receives a message from the channel, but does nothing with it
		}
	}

	wg.Add(2)

	go sender()
	go receiver()

	b.StartTimer() // begin the performance timer
	close(begin)   // start the goroutines
	wg.Wait()
}
