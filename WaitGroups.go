package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multipel goroutines to finish, we can use a wait group.

// This is the function we'll run in every goroutine.
func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
	// This WaitGroup is used to wait for all teh goroutines launched here to finish.
	// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each.
    for i := 1; i <= 3; i++ {
        wg.Add(1)

		// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
		// This wat the worker itself does not have to be aware of the concurrency primitives involved in its execution.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

    wg.Wait() // Block until the counter goes back to zero
    fmt.Println("All workers done")
}
// Worker 1 starting
// Worker 2 starting
// Worker 3 starting
// Worker 3 done
// Worker 2 done
// Worker 1 done
// All workers done