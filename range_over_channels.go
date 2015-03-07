package main

import "fmt"

// iterating over channel data

func main() {

	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

	// if we hadn't closed the channel
	// we'd block on a 3rd receive in the loop

}
