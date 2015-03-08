package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {

	// to use our pool of workers we need to send them work
	// and collect their results. we make 2 channels for this.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// sets up 3 workers
	// initially they are blocked as there are no `jobs` yet
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// we send 9 jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	// we close that channel to indicate
	// that's all the work we have
	close(jobs)

	// we collect the results of the work
	for a := 1; a <= 9; a++ {
		<-results
	}

	// all the jobs are being divided and shared among workers. 9 / 3 = 3
}
