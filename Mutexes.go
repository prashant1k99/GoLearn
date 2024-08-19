package main

import (
	"fmt"
	"sync"
)

// In the previous example we saw how to manage simple counter state using atomic operations.
// For more complex state we can use a mutex to safely access data across multiple goroutines.

// Container holds a map of counter; since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronize access. Note that mutexes musst not be copied, so if this struct is passed around, it should be done by pointer.
type Container struct {
	mu sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Lock the mutex before accessing counters; unlock it at the end of the function using defer statement.
	c.mu.Lock()
	
	defer c.mu.Unlock()

	c.counters[name]++
}

func (c *Container) dec(name string) {
	// Lock the mutex before accessing counters; unlock it at the end of the function using defer statement.
	c.mu.Lock()
	
	defer c.mu.Unlock()

	c.counters[name] = c.counters[name] - 1
}



func main() {
	c := Container {
		// Note that the zero value of a mutex is usable as-is, so no initialization is required here.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doInc := func(name string, n int) {
		for i := 0; i <n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	doDec := func(name string, n int) {
		for i := 0; i <n; i++ {
			c.dec(name)
		}
		wg.Done()
	}

	// Run several goroutines concurrently; note that they all access the same Container , and two of them access the same Counter.
	wg.Add(5)
	go doInc("a", 10000)
	go doDec("a", 290)
	go doInc("a", 10000)
	go doDec("b", 201)
	go doInc("b", 10000)
	// map[a:19710 b:9799]

	// Wait for the goroutines to finish
	wg.Wait()
	fmt.Println(c.counters)
}