package main

import "fmt"

// Closing a channel indicates that no more values will be sent on it.
// This can be useful to communicate completion to the channel's receivers.

func main() {
	// In this example we'll use a jobs channel to communicate work to be done from the main() goroutine to a worker goroutine.
	// When we have no more jobs for the worker we'll close the jobs channel.
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Here's the worker goroutine. It repeatedly receives from teh jobs with j, more := <-jobs.
	// In this special 2-value from receive, the more value will be false if jobs has been closed and all values in the channel have alreay been received.
	// We use this to notify on done when we've worked all our jobs.
	go func() {
		for {
			j, more := <- jobs
			if more {
				fmt.Println("Received job,", j)
			} else {
				fmt.Println("Recevied all jobs")
				done <- true
				return
			}
		}
	}()

	// This sends 3 jobs to the worker over the jobs channel, then closes it.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the synchronization approach we saw earlier.
	<- done

	// Reading from a closed channel succeeds immediately, returning the zero value of teh underlying tyoe.
	// The optional second return value is true if the value received was delivered by a successful send operation to the channel, or false if it was a zero
	_, ok := <-jobs
	fmt.Println("recevied more jobs:", ok)

	// sent job 1
	// sent job 2
	// sent job 3
	// sent all jobs
	// Received job, 1
	// Received job, 2
	// Received job, 3
	// Recevied all jobs
	// recevied more jobs: false
}