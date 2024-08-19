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

### 32) Timeouts:

```go
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
```

### 33) Non Blocking Channel:

```go
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
```

### 34) Closing Channel:

```go
package main

import "fmt"

// Closing a channel indicates that no more values will be sent on it.
// This can be useful to communicate completion to the channel's receivers.

func main() {
	// In this example we'll use a jobs channel to communicate work to be done from the main() goroutine to a worker goroutine.
	// When we have no more jobs for the worker we'll close the jobs channel.
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Here's the worker goroutine. It repeatedly receives from teh jobs with j, more := <-jobs.
	// In this special 2-value from receive, the more value will be false if jobs has been closed and all values in the channel have alreay been received.
	// We use this to notify on done when we've worked all our jobs.
	go func() {
		for {
			j, more := <- jobs
			if more {
				fmt.Println("Received job,", j)
			} else {
				fmt.Println("Recevied all jobs")
				done <- true
				return
			}
		}
	}()

	// This sends 3 jobs to the worker over the jobs channel, then closes it.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the synchronization approach we saw earlier.
	<- done

	// Reading from a closed channel succeeds immediately, returning the zero value of teh underlying tyoe.
	// The optional second return value is true if the value received was delivered by a successful send operation to the channel, or false if it was a zero
	_, ok := <-jobs
	fmt.Println("recevied more jobs:", ok)

	// sent job 1
	// sent job 2
	// sent job 3
	// sent all jobs
	// Received job, 1
	// Received job, 2
	// Received job, 3
	// Recevied all jobs
	// recevied more jobs: false
}
```

### 35) Range over Channles:

```go
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
```

### 36) Timers:

```go
package main

import (
	"fmt"
	"time"
)

// We ofter want to execute Go code at some point in the future, or repeatedly at some interval.
// Go's built-in timer and ticker features make both of these tasks easy. We'll look first at timers and then at tickers.

func main() {
	// Timers represent a single event in the future. You tell teh timer how long you want to waut , and it provides a channel that will be notified at that time.
	// This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer's channel C until it sends a value indicating that timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")
	// Timer 1 fired

	// If you just wanted to wait, you could have used time.Sleep. One reason a timer may be useful is that you can cancel the timer before it fires.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<- timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	// Timer 2 stopped

	// Give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped
	time.Sleep(2 * time.Second)
}
```

### 37) Tickers:

```go
package main

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals.
// here's an example of a ticker taht ticks periodically utill we stop it.

func main() {
	// Tickers use a similar mechanism to timers: a channel taht is sent values.
	// Here we'll use the select builtin on the channel to await the values as they arrive every 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	// Tick at 2024-08-18 14:52:23.562127 +0530 IST m=+0.501268168
	// Tick at 2024-08-18 14:52:24.062116 +0530 IST m=+1.001253376
	// Tick at 2024-08-18 14:52:24.562107 +0530 IST m=+1.501241584
	// Ticker stopped

	// Tickers can be stopped like timers. Once a ticker is stopped it won;t receive any more values on its channel.
	// We'll stop ours after 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
```

### 38) Worker Pool:

```go
package main

import (
	"fmt"
	"time"
)

// Worker function that processes jobs from the jobs channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2 // Example of processing
	}
}

func main() {
	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	fmt.Println("Added jobs list to worker")

	// Send 5 jobs to the jobs channel
	for j := 1; j <= numJobs; j++ {
		fmt.Println("Adding job:", j)
		jobs <- j
	}
	close(jobs) // Close the jobs channel to indicate no more jobs will be sent

	// Collect results
	for a := 1; a <= numJobs; a++ {
		fmt.Printf("Result from Job: %d\n", <-results)
	}
}
```

### 39) WaitGroup:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multipel goroutines to finish, we can use a wait group.

// This is the function we'll run in every goroutine.
func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
	// This WaitGroup is used to wait for all teh goroutines launched here to finish.
	// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each.
    for i := 1; i <= 3; i++ {
        wg.Add(1)

		// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
		// This wat the worker itself does not have to be aware of the concurrency primitives involved in its execution.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

    wg.Wait() // Block until the counter goes back to zero
    fmt.Println("All workers done")
}
// Worker 1 starting
// Worker 2 starting
// Worker 3 starting
// Worker 3 done
// Worker 2 done
// Worker 1 done
// All workers done
```

### 40) Rate Limiting:

```go
package main

import (
	"fmt"
	"time"
)

// Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of Service.
// Go elegantly supports rate limiting with goroutines, channels and tickers.

func main() {
	// First we'll look at basic rate limiting. Suppose we want to limit our handling of incoming requests. We'll serve these requests off a channel of the same name.
	requests := make(chan int, 5)
	for i := 1; i < 5; i++ {
		requests <- i
	}
	close(requests)

	// This limitier channel will receive a value every 200 ms. This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receiver from the limiter channel before serving each request, we limit ourselves to 1 request every 200 ms.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	// We may want to allow short burts of requests in our rate  limiting scheme while preserving the overall rate limit.
	// We can accomplish this by buffering our limiter channel. This burstyLimiter channel will aloow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200ms we'll try to add a new value to burstylimiter, up to its limit of 3.
	go func ()  {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incming requests. The first 3 of these will benefit from teh burst capability of burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 1; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
```
