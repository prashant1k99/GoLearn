### 26) GoRoutines:

```go
package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution

func f1(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Suppose we have a function call f(s). Here's how we'd call that in the usual way, running it synchronously.
	f1("direct")

	// To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
	go f1("goroutine")

	// You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in seperate goroutines now. Wait for them to finish (for a more robust approach, use a WaitGroup)
	time.Sleep(time.Second)
	fmt.Println("done")

	// direct : 0
	// direct : 1
	// direct : 2
	// going
	// goroutine : 0
	// goroutine : 1
	// goroutine : 2
	// done
}
```

### 27) Channels:

```go
package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

func main() {
	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	messages := make(chan string)

	go func() {
		// Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
		messages <- "ping"
		messages <- "pong"
	}()


	// The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
	msg := <- messages
	fmt.Println(msg)
	// ping
	msg = <- messages
	fmt.Println(msg)
	// pong
}
```
