package main

import (
	"fmt"
	"time"
)

func main() {

	// you tell the timer how long to wait and it provides a channel
	// that will be notified at that time
	// this timer will wait 2 seconds
	timer1 := time.NewTimer(time.Second * 2)

	// timer1.C blocks on the timer's channel
	// until it sends a value indicating that the timer has expired
	<-timer1.C

	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	// cancelling the timer
	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("timer 2 stopped")
	}
}
