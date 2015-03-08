package main

import "time"
import "fmt"

func main() {

	// uses a similar mechanism to timers:
	// a channel that is sent values

	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1500)

	// tickers can be stopped like timers
	ticker.Stop()

	fmt.Println("Ticker stopped")
}
