// another alternative to mutexes is to
// use the built-in synchronization features
// of goroutines and channels to achieve the same result.
// this channel-based approach aligns with Go's ideas
// of sharing memory by communicating
// and having each piece of data owned by exactly
// 1 goroutine

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// our state will be owned by a single goroutine
// this will guarantee that data is never corrupted
// with concurrent access
// in order to read or write, other goroutines
// will send messages to the owning goroutine
// and receive corresponding replies.
// `readOp` and `writeOp` structs encapsulate those
// requests and a way for the owning goroutine to respond

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// counting operations
	var ops int64 = 0

	// reads and writes channels will be used
	// by other go routines to issue read and write
	// requests
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// this is the goroutine that owns state, which is a map that is private
	// to the stateful goroutine
	// this goroutine repeatedly selects on the reads and writes channels
	// repsonding to requests as they arrive.
	// a response is executed by first performing the requested operation
	// and then sending a value on the response channel `resp` to indicate
	// success and the desired value in the case of `reads`
	go func() {

		var state = make(map[int]int)

		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// this starts 100 goroutines to issue reads.
	// each read requires a constructing `readOp`, sending it over the reads channel
	// and receiving the result over the provided `resp` channel
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 10 writes follow a similar approach
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

}
