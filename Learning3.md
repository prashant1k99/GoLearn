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

### 28) Channel Buffering:

```go
package main

import "fmt"

// By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.

func main() {
	// Here we make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<- messages)
	// buffered
	fmt.Println(<- messages)
	// channel
}
```

### 29) Channel Synchronization:

```go
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
```

### 30) Channel Direction:

```go
package main

import "fmt"

// When using channels as function parameters, you can specify if a channel is meant to only send or receive values. This specificity increases the type-safety of the program.

// This ping function only accepts a channel for sending values. It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan <- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
```

### 31) Select:

```go
package main

import (
	"fmt"
	"time"
)

// Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.

func main() {
	// For our example we’ll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We’ll use select to await both of these values simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recevied", msg1)
		case msg2 := <-c2:
			fmt.Println("recevied", msg2)
		}
	}
}
```
