package main

import "time"

import "fmt"

func main() {

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// select is implementing the timeout

	// select proceeds with the first receive that's ready
	// so it will take the timeout case if the first
	// operation takes more than the allowed 1s

	select {
	case res := <-c1: //awaits the result
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	//  since <- c2 is the first value to surface, it will print the res first
	select {
	case res := <-c2:
		fmt.Print(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

}
