package main

import (
	"fmt"
	h "goconcurrency/chpt_4/helpers"
	"math/rand"
)

func main() {
	tee := func(
		done <-chan interface{},
		in <-chan interface{},
	) (_, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})
		go func() {
			defer close(out1)
			defer close(out2)
			for val := range h.OrDone(done, in) {
				var out1, out2 = out1, out2
				for i := 0; i < 2; i++ {
					select {
					case <-done:
					case out1 <- val:
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2
	}

	done := make(chan interface{})
	defer close(done)

	out1, out2 := tee(
		done,
		h.Take(
			done,
			h.RepeatFn(
				done,
				func() interface{} { return rand.Int() }),
			6),
	)

	for val1 := range out1 {
		fmt.Printf("out1: %v, out2:%v\n", val1, <-out2)
	}

}
