package main

import (
	"fmt"
)

// Go makes it possible to recover from a panic, by using the recover built-in function. A recover can stop a panic from aborting the program and let it continue with execution instead.

// An example of where this can be useful: a server wouldn’t want to crash if one of the client connections exhibits a critical error. Instead, the server would want to close that connection and continue serving other clients. In fact, this is what Go’s net/http does by default for HTTP servers.


func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Starting the program...")

	mayPanic()

	fmt.Println("This line won't execute due to panic.")
}

func mayPanic() {
	fmt.Println("About to cause a panic...")
	panic("Something went wrong!")
}

/*
Explanation:
Deferred Function with recover:

The defer statement in main() ensures that the anonymous function containing recover() is called after main() returns, or when a panic occurs.
recover() checks if there's a panic. If there is, it returns the panic value, which is then handled (in this case, just printed out). If there’s no panic, recover() returns nil.
Panic Triggering:

The mayPanic() function triggers a panic using panic("Something went wrong!").
Normally, this would crash the program and stop execution. But since recover() is present in a deferred function, it catches the panic, and the program continues running after handling it.
Output:

The output will show that the program starts, the panic occurs, and then the panic is recovered, allowing the program to finish gracefully.
Starting the program...
About to cause a panic...
Recovered from panic: Something went wrong!
*/