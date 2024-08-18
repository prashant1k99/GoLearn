package main

import (
	"fmt"
	"time"
)

// Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.
// Implementing timeouts in Go is easy and elegant thanks to channels and select

func main() {
	// For our example, suppose we're executing an exteral call that returns on a channel c1 after 2s. Note that the channel is buffered, so the send in teh goroutine is nonlocking.
	// This is a common pattern to prevent goroutine leaks in case the channel is never read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result"
	}()

	// Here's the select implementing a timeout. res := <- c1 awaits the result and <-time.After awaits a value to be sent after the timeout of 1s.
	// Since select proceeds with teh first receive that's ready, we'll take the timeout case if the operation takes more than the allowed 1s.
	select {
	case res := <- c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}
	// Timeout 1

	// If we allow a longer timeout of 3s, then the receive from c2 will success and we'll pring the result.
	select {
	case res := <- c1:
		fmt.Println(res)
		// result
	case <- time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}