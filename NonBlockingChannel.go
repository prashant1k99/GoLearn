package main

import "fmt"

// Basic sends and receive on channels are blocking. However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.

func main() {
	messages := make(chan string)
	signals  := make(chan bool)

	// Here's a non-blocking reeive. If a value is available on messages then select will take the <- messages case with that value.
	// If not it will immediately take the defaul case.
	select {
	case msg := <- messages:
		fmt.Println("Received message", msg)
	default: fmt.Println("no message received")
	}
	// no message received

	// A non-blocking sends works similarly. Here msg cannt be sent to the messages channel, because the channel has no buffer and there is no reveiver. Therefore the default case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	// no message sent

	// We can use multiple cases above the default claude to implement a multi-way non-blocking select. Here we attempt non-blocking select. Here we attempt non-blocking receives on both messages and signals.
	select {
	case msg := <- messages:
		fmt.Println("recevied message", msg)
	case sig:= <- signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
	// no activity
}