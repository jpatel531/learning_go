// when using channels as function parameters
// you can specify if a channel is only meant to send or receive
// data. this increases the type-safety of the program

package main

import "fmt"

func ping(pings chan<- string, msg string) {
	// will only accept a channel for sending values
	pings <- msg

}

func pong(pings <-chan string, pongs chan<- string) {

	// accepts one channel for receives (pings)
	// a second for sends (pongs)
	msg := <-pings
	pongs <- msg

}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
