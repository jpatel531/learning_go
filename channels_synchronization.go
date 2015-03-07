// we use channels to synchronize execution across goroutines
// here's an example of using a blocking to receive to wait
// for a gorouting to finish

package main

import "fmt"
import "time"

// the done channel will be use to
// notify another goroutine that this
// function's work is done
func worker(done chan bool) {
	fmt.Print("working")
	time.Sleep(time.Second)
	fmt.Println("done")
	//send a value to notify when we're done
	done <- true
}

func main() {

	done := make(chan bool, 1)

	// start a worker routine, giving it the channel to notify on
	go worker(done)

	// block until we receive a notification from the channel
	<-done

	// if you remove the `<- done` line from this program
	// the program would exit before the `worker` even started

}
