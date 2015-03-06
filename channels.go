package main

import "fmt"

// channels are pipes that connect concurrent goroutings
// you can send values into channels from one and receive in another

func main() {

	//channels are typed by the values they convey
	messages := make(chan string)

	// we'll send a value into a channel using the channel <- syntax
	go func() { messages <- "ping" }()

	// <- channel syntax receives a value from the channel
	msg := <-messages
	fmt.Println(msg)

	// by default sends and receives block until both sender and receiver are ready
	// this property allowed us to wait at the end of pur program for the ping message
	// without having to use any other synchronization
}
