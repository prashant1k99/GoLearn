package main

import (
	"bufio"
	"fmt"
	"os"
)

// Writing files in Go follows similar patterns to the ones we saw earlier for reading.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// To start, here's how to dump a string (or just bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create("/tmp/dat2")
	check(err)

	// It's idiomatic to defer a Close immediately after opening a file.
	defer f.Close()

	// You can Write byte slices as you's expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
	// wrote 5 bytes

	// A WriteString is also available.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	// Wrote 7 bytes

	// Issue a sync to flush writes to stable storage.
	f.Sync()

	// bufio provides buffered writes in addition to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffer\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	// wrote 7 bytes

	// Use Flush to ensure all buffered operations have been applied to the underlying writer.
	w.Flush()
}
