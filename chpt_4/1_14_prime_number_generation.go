package main

import (
	"fmt"
	h "goconcurrency/chpt_4/helpers"
	"math/rand"
	"time"
)

func main() {
	randGen := func() interface{} { return rand.Intn(50000000) }
	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := h.ToInt(done, h.RepeatFn(done, randGen))
	fmt.Println("Primes")
	for prime := range h.Take(done, h.PrimeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
