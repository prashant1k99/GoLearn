package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Decrement the counter when the goroutine completes
    fmt.Printf("Worker %d starting\n", id)
    
    // Simulate some work
    time.Sleep(time.Second)
    
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Increment the counter
        go worker(i, &wg)
    }

    wg.Wait() // Block until the counter goes back to zero
    fmt.Println("All workers done")
}

/*
Explanation:
wg.Add(1): This increments the WaitGroup counter by 1. We do this every time we start a new goroutine.

go worker(i, &wg): This starts a new goroutine, passing the WaitGroup by reference.

wg.Done(): This decrements the WaitGroup counter by 1. It’s called when the goroutine completes its work. We use defer to ensure it's called at the end of the function, even if the function exits early.

wg.Wait(): This blocks the main goroutine until the counter goes back to zero, meaning all worker goroutines have finished.

Worker 1 starting
Worker 2 starting
Worker 3 starting
Worker 1 done
Worker 2 done
Worker 3 done
All workers done

*/