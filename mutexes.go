// for more complex state, we can use a mutex
// to safely access data across multiple
// goroutines

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// state will be a map
	var state = make(map[int]int)

	// mutex will sync access to state
	var mutex = &sync.Mutex{}

	// ops will count the number of operations performed
	var ops int64 = 0

	// 100 goroutines that execute repeated reads against state
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)      // pick a key to access
				mutex.Lock()             // lock to ensure exclusive access
				total += state[key]      // read the value at key
				mutex.Unlock()           // unlock it
				atomic.AddInt64(&ops, 1) // increment the atomic counter
				runtime.Gosched()        // yield after each operation. the yielding is handled automatically
				// e.g. with every channel operationa nd for blocking calls like time.Sleep
				// but in this case we need to do it manually.
			}
		}()
	}

	// 10 goroutines to simulate writes, using same pattern as for reads
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()

			}
		}()
	}

	time.Sleep(time.Second)

	// take final operations count
	opsFinal := atomic.LoadInt64(&ops)

	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	// see how state ended up
	fmt.Println("state:", state)
	mutex.Unlock()

}
