### 41) Atomic Counters:

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// The primary mechanism for managing state in Go is communication over channels.
// We saw this for example with worker pools. There are a few other options for managing state though.
// Here we'll lool at using the sync/await pacakge for atomic counters accessed by multiple goroutines.

func main() {
	// We'll use an atomic integer type to represent our (always positive) couter.
	var ops atomic.Uint64

	// A WaitGroup will help us wait for all goroutines to finish their work.
	var wg sync.WaitGroup

	// We'll start 50 goroutines that each increment the counter exactly 1000 times.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// To atomically increment the counter we use Add.
				ops.Add(1)
			}

			wg.Done()
		}()
	}

	// Wait until all teh goroutines are done.
	wg.Wait()

	// Here no goroutines are writing to 'ops', but using Load it's safe to atomically read a value even while other goroutines are (atomically) updating it.
	fmt.Println("ops:", ops.Load())
	// ops: 50000
}
```

### 42) Mutexes:

```go
package main

import (
	"fmt"
	"sync"
)

// In the previous example we saw how to manage simple counter state using atomic operations.
// For more complex state we can use a mutex to safely access data across multiple goroutines.

// Container holds a map of counter; since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronize access. Note that mutexes musst not be copied, so if this struct is passed around, it should be done by pointer.
type Container struct {
	mu sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Lock the mutex before accessing counters; unlock it at the end of the function using defer statement.
	c.mu.Lock()

	defer c.mu.Unlock()

	c.counters[name]++
}

func (c *Container) dec(name string) {
	// Lock the mutex before accessing counters; unlock it at the end of the function using defer statement.
	c.mu.Lock()

	defer c.mu.Unlock()

	c.counters[name] = c.counters[name] - 1
}



func main() {
	c := Container {
		// Note that the zero value of a mutex is usable as-is, so no initialization is required here.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doInc := func(name string, n int) {
		for i := 0; i <n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	doDec := func(name string, n int) {
		for i := 0; i <n; i++ {
			c.dec(name)
		}
		wg.Done()
	}

	// Run several goroutines concurrently; note that they all access the same Container , and two of them access the same Counter.
	wg.Add(5)
	go doInc("a", 10000)
	go doDec("a", 290)
	go doInc("a", 10000)
	go doDec("b", 201)
	go doInc("b", 10000)
	// map[a:19710 b:9799]

	// Wait for the goroutines to finish
	wg.Wait()
	fmt.Println(c.counters)
}
```

### 43) Stateful Goroutines:

```go
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// In the previous example we used explicit locking with mutexes to synchronize access to shared state across multiple goroutines.
// Another option is to use the built-in synchronization features of goroutines and channels to achieve the same result.
// This channel-based approach aligns with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.

// In this example our state will be owned by a single goroutine. This will guarantee that the data is never corrupted with concurrent access.
// In order to read or write that state, other goroutines will send messages to the owning goroutine and receive corresponding replies.
// These readOp and writeOp structs encapsulate those requests and a way for the owning goroutine to respond.
type readOp struct {
	key int
	resp chan int
}
type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	// As before we'll count how many operatrions we perform
	var readOps uint64
	var writeOps uint64

	// The reads and writes channel will be used by other goroutines to issue read and write requests, respectivley
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Here is the goroutine that owns the state, which is a map as in the previous example but now private to the stateful goroutine. This goroutine repeatedly selects on the reads and writes channels, responding to requests as they arrive.
	// A response is executed by first performing the requested operation and then sending a value on the response channel resp to indicate success (and the desired value in the case of reads).
	go func ()  {
		var state = make(map[int]int)

		for {
			select {
			case read := <- reads:
				read.resp <- state[read.key]
			case write := <- writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// This starts 100 goroutines to issue reads to the stateowning goroutine via the reads channel.
	// Each read requires constructing a readOp, sending it over the reads channel, and then receiving the result over the provided resp channel.
	for r := 0; r < 100; r++ {
		go func ()  {
			for {
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<- read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We start 10 writes as well, using a similar approach
	for w := 0; w < 10; w++ {
		go func ()  {
			for {
				write := writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<- write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the goroutine work for a second.
	time.Sleep(time.Second)

	// Finally, capture and report the op counts.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	// readOps: 85050
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
	// writeOps: 8538
}
```

### 44) Sorting:

```go
package main

import (
	"fmt"
	"slices"
)

// Go's slices pacakge implements sorting for builtins and user-defined types. We'll look at sorting for builtins first.

func main() {
	// Sorting functions are generic, and work for any ordered built-in type.
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("strings:", strs)
	// strings: [a b c]

	// An example of sorting ints
	ints := []int{7,5,2,9,100}
	slices.Sort(ints)
	fmt.Println("ints:", ints)
	// ints: [2 5 7 9 100]

	// We can also use the slices package to check if a slice is already in sorted order.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted:", s)
	// Sorted: true
}
```

### 45) Sorting by Functions:

```go
package main

import (
	"cmp"
	"fmt"
	"slices"
)

// Sometimes we'll want to sort a collection by something other than its natural order. For example, suppose we wanted to sort strings by their length instead of alphabetically. Here's an example of custom sorts in Go.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// We implement a comparison function for string lengths. cpm.Compare is helpful for this.
	lenCmp := func (a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Now we can call slices.SortFunc with this custom comparison function to sort fruits by name length
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)
	// [kiwi peach banana]

	// We can use the same technique to sort a slice of values that aren't built-in types.
	type Person struct {
		name string
		age int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 71},
	}

	// Sort people by age using slices.SortFunc.
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)
	// [{TJ 25} {Jax 37} {Alex 71}]
}
```
