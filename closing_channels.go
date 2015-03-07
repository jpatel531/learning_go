package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// 'more' value will be false if `jobs` has been closed
			j, more := <-jobs

			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	// sent over 3 jobs to the worker over the jobs channel
	// we'll close the jobs channel
	close(jobs)

	fmt.Println("sent all jobs")

	// waits for `done` to have a value, then closes the app
	<-done

}
