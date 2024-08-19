package main

import (
	"fmt"
)

// Define commands that can be sent to the stateful goroutine
type CounterCommand struct {
	increment bool
	get       bool
	response  chan int
}

func main() {
	// Create a command channel
	cmdChan := make(chan CounterCommand)
	
	// Start the stateful goroutine
	go func() {
		counter := 0 // The state of the goroutine
		for cmd := range cmdChan {
			if cmd.increment {
				counter++
			}
			if cmd.get {
				cmd.response <- counter
			}
		}
	}()

	// Increment the counter
	cmdChan <- CounterCommand{increment: true}

	// Get the current counter value
	respChan := make(chan int)
	cmdChan <- CounterCommand{get: true, response: respChan}
	fmt.Println("Counter:", <-respChan)
	// Counter: 1

	// Increment the counter again
	cmdChan <- CounterCommand{increment: true}

	// Get the current counter value again
	cmdChan <- CounterCommand{get: true, response: respChan}
	fmt.Println("Counter:", <-respChan)
	// Counter: 2

	// Close the command channel to stop the goroutine
	close(cmdChan)
}

/*
Explanation:
Command Struct (CounterCommand): This struct is used to define the operations we want to perform on the stateful goroutine. It has flags for incrementing the counter, getting the current value, and a channel to send back the response.

Stateful Goroutine: The goroutine runs an infinite loop where it listens for commands on the cmdChan channel. It maintains its internal state (counter) and modifies or reports this state based on the commands it receives.

Increment Command: When an increment command is received, the goroutine increments its internal counter state.

Get Command: When a get command is received, the goroutine sends the current value of the counter back on the response channel.

Encapsulation: The state (counter) is fully encapsulated within the goroutine, meaning no other goroutine can directly access or modify it. All interaction with the state happens via messages passed through channels.

Benefits:
No Explicit Synchronization: Since the state is encapsulated within a single goroutine, there's no need for Mutexes, atomic operations, or other synchronization primitives.

Reduced Risk of Race Conditions: By isolating state within a single goroutine, you avoid the complexities and potential pitfalls of shared state in concurrent programming.

Clear Communication Patterns: Using channels to interact with the goroutine makes it clear how state is being accessed and modified, which can make the program easier to reason about.


*/