package main

// by default channels are unbuffered, meaning they will only accept
// sends (chan <-) if there is a corresponding receive
// (<- chan) ready to receive the sent value.
// buffered channels acept a limited number of values without a correspodning
// receiver for those values

import "fmt"

func main() {

	//make a channel of strings buffering up to 2 values
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
