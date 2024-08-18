package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// Sender Goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i // Send numbers to the channel
		}
		close(ch) // Close the channel after sending all values
	}()

	// Receiver Goroutine
	for val := range ch { // Receive values from the channel
		fmt.Println(val)
	}

	fmt.Println("Channel closed, no more data to receive")
}

/*
Explanation:

Channel Creation: We create an unbuffered channel ch of type int.

Sender Goroutine:

A goroutine is started that sends the numbers 1, 2, and 3 to the channel.
After sending all the values, the close(ch) function is called to close the channel. This signals that no more values will be sent on the channel.
Receiver Goroutine:

The main goroutine reads from the channel using a for range loop. This loop continues to receive values from the channel until the channel is closed.
Once the channel is closed and all values are received, the loop exits.
Final Output:

The program prints the received values (1, 2, 3) and then prints "Channel closed, no more data to receive" once the loop exits.
Important Points:
Receiving After Channel Close: Once the channel is closed and all values are consumed, any further attempts to receive from the channel will yield the zero value for the channel’s type (e.g., 0 for int, "" for string, etc.).

Detecting Channel Closure: The range loop over a channel automatically stops when the channel is closed.
*/