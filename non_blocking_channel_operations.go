// sends and receives on channels are blocking

// we can use select with a default clause to implement
// non-blocking sends, receives, and
// even non-blocking multi-way selects

package main

import "fmt"

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	// non-blocking receive
	// if a value is available on `message` then select will take the
	// <- messages case with that value. if not it will immediately
	// take the default case
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// a non-blocking send works similarly
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// we can use multiple cases above the default caluse to implement
	// multi-way non-blocking select.
	// here we attempt non-blocking receives on both `messages` and `signals`
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
