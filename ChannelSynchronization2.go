package main

import (
	"fmt"
	"time"
)

func worker(id int, done chan bool) {
    fmt.Printf("Worker %d: Starting work...\n", id)
    time.Sleep(time.Second * time.Duration(id)) // Simulate work with a sleep
    fmt.Printf("Worker %d: Finished work.\n", id)
    done <- true // Signal that this worker is done
}

func main() {
    done := make(chan bool) // Create a channel to synchronize

    // Start 3 worker goroutines
    for i := 1; i <= 3; i++ {
        go worker(i, done)
    }

    // Wait for all workers to finish
    for i := 1; i <= 3; i++ {
        <-done // Receive a signal from each worker
    }

    fmt.Println("All workers are done. Main function can continue now.")
	// Worker 3: Starting work...
	// Worker 2: Starting work...
	// Worker 1: Starting work...
	// Worker 1: Finished work.
	// Worker 2: Finished work.
	// Worker 3: Finished work.
	// All workers are done. Main function can continue now.
}
