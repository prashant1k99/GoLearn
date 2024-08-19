package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 // The counter must be of an int64 type for atomic operations.
	var wg sync.WaitGroup

	// Number of goroutines
	numGoroutines := 10

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			fmt.Println("In goroutine:", i)
			// Atomically increment the counter
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	// Print the final value of the counter
	fmt.Println("Final Counter Value:", counter)
}

/*
Explanation:
int64 Counter: The counter is defined as int64 because atomic operations in Go work with specific types like int32 or int64.
sync/atomic.AddInt64: This function atomically adds 1 to the counter. It ensures that even if multiple goroutines try to update the counter simultaneously, the operations won't conflict, leading to a correct final value.
sync.WaitGroup: Used to wait for all the goroutines to finish before printing the final counter value.
In this example, if numGoroutines is set to 10, the final counter value should be 10, as each goroutine increments the counter by 1.

Output:
In goroutine: 0
In goroutine: 9
In goroutine: 5
In goroutine: 6
In goroutine: 7
In goroutine: 8
In goroutine: 3
In goroutine: 2
In goroutine: 4
In goroutine: 1
Final Counter Value: 10
*/