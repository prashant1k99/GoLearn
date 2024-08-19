package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Number of goroutines
	numGoroutines := 10

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			// Lock the mutex before accessing the shared variable
			mu.Lock()
			counter++
			// Unlock the mutex after the shared variable is updated
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	// Print the final value of the counter
	fmt.Println("Final Counter Value:", counter)
	// Final Counter Value: 10
}

/*
Explanation:
sync.Mutex: The Mutex is a locking mechanism used to ensure that only one goroutine can access the critical section of code (in this case, the increment of counter) at a time.

mu.Lock(): This locks the Mutex, preventing any other goroutine from entering the critical section until the Mutex is unlocked.

Critical Section: The code between mu.Lock() and mu.Unlock() is the critical section where shared data (counter) is being modified.

mu.Unlock(): This unlocks the Mutex, allowing other goroutines to enter the critical section.

Race Condition Prevention: Without the mutex, if multiple goroutines tried to increment the counter simultaneously, it could lead to a race condition where the counter value becomes unpredictable. The mutex ensures that only one goroutine can modify the counter at any given time.
*/