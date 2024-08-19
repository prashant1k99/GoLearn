package main

import "os"

// A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn't occur during normal operation, or that we aren't prepared to handle gracefully.
func main() {
	// We'll use panic throughout this site to check for unexpected errors. This si sthe only program on teh site designed to panic.
	// panic("a problem")
	/*
	panic: a problem

	goroutine 1 [running]:
	main.main()
			/Users/prashantsingh/Desktop/Learn/GoLearn/Panic.go:6 +0x2c
	exit status 2
	*/

	// A common use of panic is to abort if a function returns an error value that we don't know how to (or want to) handle.
	// Here's an example of panicking if we get an unexpecetd error when creating a new file
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}