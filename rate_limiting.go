// rate limiting controls resource utilization
// and maintains quality of service
// go elegantly supports rate limiting with goroutines
// channels
// and tickers

package main

import "time"
import "fmt"

func main() {

	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	// the limiter channel will receive a value every 200 ms.
	// this is the regulator in our rate limiting scheme
	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		// by blocking on a receive from the limiter channel before
		// serving each request, we limit ourselves to
		// 1 request every 200ms
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// to allow short bursts of requests in our rate limit while
	// preserving the overall rate limit,
	// we buffer our limiter channel. this will allow bursts
	// of up to 3 events
	burstyLimiter := make(chan time.Time, 3)

	//fill up the channel
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// every 200 ms we'll try to add a value to burstyLimiter
	// up to its limit of 3
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// simulate 5 more incoming requests
	// the first 3 will benefit from the burst capacity of burstyLimiter
	burstyRequests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	// Running our program we see the first batch of requests
	// handled once every ~200 milliseconds as desired.
	// For the second batch of requests we serve the first 3
	// immediately because of the burstable rate limiting,
	// then serve the remaining 2 with ~200ms delays each.

}
