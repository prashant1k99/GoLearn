package main

import (
	"fmt"
	"time"
)

func sendToChannel(ch chan<- string, message string, delay time.Duration) {
    time.Sleep(delay) // Simulate some work with a delay
    ch <- message     // Send the message to the channel
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    // Start two goroutines to send data to the channels
    go sendToChannel(ch1, "Message from Channel 1", 2*time.Second)
    go sendToChannel(ch2, "Message from Channel 2", 1*time.Second)

    // Use select to wait on both channels
    select {
    case msg1 := <-ch1:
        fmt.Println("Received:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received:", msg2)
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout: No messages received")
    }
}

/* 
Explanation:
sendToChannel Function:

This function simulates doing some work by sleeping for a specified duration before sending a message to the channel.
Main Function:

Two channels, ch1 and ch2, are created to carry strings.
Two goroutines are started to send messages to these channels after a delay (2 seconds for ch1 and 1 second for ch2).
The select statement then waits for one of the channels to send a message.
The select has three cases:
If ch1 sends a message first, it prints that message.
If ch2 sends a message first, it prints that message.
If neither channel sends a message within 3 seconds, a timeout message is printed using time.After, which creates a channel that sends a message after the specified duration.
*/