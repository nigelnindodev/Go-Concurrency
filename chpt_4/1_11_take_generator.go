package main

import "fmt"

func main() {
	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				/**
				- Remember takeStream can block valuesStream. One instance is when valueStream is a channel that never stops
					producing values. In such a case it might actually be beneficial as the takeStream will prevent
					creation of more values after it takes x elements required.
				- TODO: How would this behave for a buffered channel? I have a hunch, try it out yourself if reading this.
				*/
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})

	valueStream := make(chan interface{})
	go func() {
		values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for _, v := range values {
			fmt.Printf("Adding %d to valueStream\n", v)
			valueStream <- v
		}
	}()

	takenValues := take(done, valueStream, 4)
	for v := range takenValues {
		fmt.Println(v)
	}
}
