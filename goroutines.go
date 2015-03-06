// go routine is a lightweight thread of execution

package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// usual, synchronous way of calling a function
	f("direct")

	// this new routine will execute concurrently with the calling one
	go f("goroutine")

	// goroutine from anonymous function

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// our two function calls are running
	// async in separate goroutines. so execution falls through to here:
	// Scanln requires we press a key before the program exits

	var input string
	fmt.Scanln(&input)

	fmt.Println("done")

	// we see the output of the blocking calls first, then the interweaved
	// output of the two goroutines, which are concurrent in the Go runtime.
}
