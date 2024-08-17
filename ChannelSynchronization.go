package main

import (
	"fmt"
	"time"
)

// We can use channels to synchronize execution across goroutines. Here’s an example of using a blocking receive to wait for a goroutine to finish. When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.

// This si the function we'll run in a goroutine. The done channel will be used to notify another aoroutine that this functon work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done
	done <- true
}

func main() {
	// Start a worker goroutine, giving it the channel to notify on.
	done := make(chan bool, 1)
	go worker(done)
	// zation.go
	// working...done

	// Block until we receive a notification from the worker on the channel.
	<- done
}