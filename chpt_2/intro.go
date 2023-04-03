package main

import (
	"fmt"
	"sync"
)

func main() {
	var data int
	var memoryAccess sync.Mutex

	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if data == 0 {
		fmt.Printf("the value is %v /n", data)
	}
	memoryAccess.Unlock()
}
