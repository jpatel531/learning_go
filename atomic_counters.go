package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {

	var ops uint64 = 0 // unsigned int to represent our always-positive counter

	// we'll start 50 go routines that increment the counter once a milisecond
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1) //atomically incrementing the counter, giving it the ops memory address
				runtime.Gosched()         //allow other goroutines to proceed

				// each routine is waiting for others to have incremented the count,
				// before it starts incrementing
			}
		}()
	}

	time.Sleep(time.Second) // wait a second to allow some ops to accumulate

	opsFinal := atomic.LoadUint64(&ops) //fetch the value from the memory address

	fmt.Println("ops:", opsFinal)

	// it's doing a lock, yeah?

}
