package main

import "fmt"

// In closing channel example we saw how for and range provide iteration over basic data structure.
// We can also use this syntax to iterate over values received from a channel.

func main() {
	// we'll iterate over 2 values in the queue channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This range iterated over each element as it's received from queue. Because we closed the cahnnela above, the iteration terminated after receiving teh 2 elements.
	for elem := range queue {
		fmt.Println(elem)
		// one
		// two
	}
}