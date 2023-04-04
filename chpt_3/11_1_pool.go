package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}
	myPool.Get()             // creates a new instance and invokes the new function
	instance := myPool.Get() // creates a new instance as well and invokes the new function
	myPool.Put(instance)     // puts the previously created instance back into the pool
	myPool.Get()             // will not call the New function as it will use the instance returned into the pool
}
