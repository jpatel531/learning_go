// select lets you wait on multiple channel operations

package main

import "time"
import "fmt"

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	// each channel will receive a value after some
	// amount of time
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {

		// we'll use selec tto wait both of these values
		// simultaneously, printing each one as it arrives

		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// Note that the total execution time is only ~2 seconds
	// since both the 1 and 2 second Sleeps execute concurrently.
}
