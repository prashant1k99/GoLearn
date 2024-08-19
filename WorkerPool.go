package main

import (
	"fmt"
	"time"
)

// Worker function that processes jobs from the jobs channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2 // Example of processing
	}
}

func main() {
	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	fmt.Println("Added jobs list to worker")

	// Send 5 jobs to the jobs channel
	for j := 1; j <= numJobs; j++ {
		fmt.Println("Adding job:", j)
		jobs <- j
	}
	close(jobs) // Close the jobs channel to indicate no more jobs will be sent

	// Collect results
	for a := 1; a <= numJobs; a++ {
		fmt.Printf("Result from Job: %d\n", <-results)
	}
}

/* 
Explanation:

Worker Function:

The worker function continuously receives jobs from the jobs channel.
Each job is processed (in this case, a simple operation of multiplying by 2).
The result is sent to the results channel.
Main Function:

A job channel jobs is created to hold the tasks, and a results channel results is created to collect the outputs.
Three worker goroutines are started. Each worker will pick up jobs from the jobs channel as they become available.
Five jobs are sent to the jobs channel, and the channel is closed to indicate no more jobs will be added.
The main goroutine waits for results from all jobs and prints them.
Output
The output will show that jobs are being processed concurrently by different workers:
Added jobs list to worker
Adding job: 1
Adding job: 2
Adding job: 3
Adding job: 4
Adding job: 5
Adding job: 6
Adding job: 7
Adding job: 8
Adding job: 9
Adding job: 10
Worker 3 started job 3
Worker 1 started job 1
Worker 2 started job 2
Worker 2 finished job 2
Worker 2 started job 4
Result from Job: 4
Worker 1 finished job 1
Worker 1 started job 5
Result from Job: 2
Worker 3 finished job 3
Worker 3 started job 6
Result from Job: 6
Worker 1 finished job 5
Worker 1 started job 7
Result from Job: 10
Worker 3 finished job 6
Worker 3 started job 8
Result from Job: 12
Worker 2 finished job 4
Worker 2 started job 9
Result from Job: 8
Worker 2 finished job 9
Worker 3 finished job 8
Worker 1 finished job 7
Result from Job: 18
Result from Job: 16
Worker 2 started job 10
Result from Job: 14
Worker 2 finished job 10
Result from Job: 20

========================
Key Points:

Job Queue: The jobs channel is where tasks are queued up.
Concurrency: Multiple workers process jobs concurrently, which can improve performance by taking advantage of multiple CPU cores.
Channel Closing: Closing the jobs channel is important to signal to the workers that there are no more jobs to process.
Advantages of a Worker Pool
Resource Management: Limits the number of active goroutines, avoiding excessive resource usage.
Scalability: Can efficiently handle a large number of tasks.
Separation of Concerns: Workers focus on processing, while the main routine can handle task distribution and result collection.
Use Cases
Web Servers: Handling multiple incoming requests concurrently.
Task Queues: Processing background jobs like image processing, data aggregation, etc.
Parallel Processing: Distributing heavy computations across multiple goroutines.
*/