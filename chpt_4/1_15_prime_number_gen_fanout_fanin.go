package main

import (
	"fmt"
	h "goconcurrency/chpt_4/helpers"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	randGen := func() interface{} { return rand.Intn(50000000) }
	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := h.ToInt(done, h.RepeatFn(done, randGen))
	fmt.Println("Primes")

	/**
	- Fan out the number of prime finders according to the number of CPU's available
	- All fanned out channels will read from the same random source
	- NB: once a number is read by worker off the channel it is pushed out, so the other workers will not process it.
	- NB: This is much faster than the previous example as we are using n workers to go through the numbers trying to find primes
	*/
	numFinders := runtime.NumCPU()
	finders := make([]<-chan interface{}, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = h.PrimeFinder(done, randIntStream)
	}

	for prime := range h.Take(done, h.FanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
