package main

import (
	"fmt"
	"time"
)

// Function that sends data into a channel (send-only)
func sender(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i // Send data into the channel
        time.Sleep(time.Millisecond * 100)
    }
    close(ch) // Close the channel when done sending
}

// Function that receives data from a channel (receive-only)
func receiver(ch <-chan int) {
    for num := range ch {
        fmt.Println("Received:", num) // Receive data from the channel
    }
	// Received: 1
	// Received: 2
	// Received: 3
	// Received: 4
	// Received: 5
}

func main() {
    ch := make(chan int) // Create an integer channel

    // Start sender and receiver goroutines
    go sender(ch)
    go receiver(ch)

    // Wait for a while to let the goroutines finish
    time.Sleep(time.Second)
}
