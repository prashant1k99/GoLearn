package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Reading and writing files are basic tasks needed for many Go programs. First we'll look at some examples of reading files.

// Reading files requires checking most calls for errors. This helper will streamline our errors check below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Perhaps the most basic file reading task is slurping a file's entire contents into memory.
	dat, err := os.ReadFile("./README.md")
	check(err)
	fmt.Print(string(dat))
	// ### Learning Go by Practice

	// You'll often want more control over how and what parts of a file are read. For these tasks, start by Opening a file to obtain an os.File value.
	f, err := os.Open("./README.md")
	check(err)
	fmt.Println(f)
	// &{0x14000128180}

	// Read some bytes from the beginnign of the file. Allow up to 5 to be read but also note how many actually were read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	// ### Learning Go by Practice

	// This is the practice ground for Prashant to test his GoLang code.

	// You can also Seek to a known location in the file and Read from there.
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	// 2 bytes @ 6: ar
	fmt.Printf("%v\n", string(b2[:n2]))

	// Other methods of seeking are relative to teh current cursor position,
	_, err = f.Seek(4, io.SeekCurrent)
	check(err)

	// and relative to teh end of the file.
	_, err = f.Seek(-10, io.SeekEnd)
	check(err)

	// The io package provides some functions that may be helpful for file reading.
	// For example, reads like the ones above can be more robustly implemented with ReadAtLeast.
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s \n", n3, o3, string(b3))
	// 2 bytes @ 6: ar

	// There is no built-in rewind, but Seek(0, io.SeekStart) accomplishes this.
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// The bufio package implements a buffered reader that may be useful both for its efficiency with many small reads and because of the additional reading methods it provides.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	// 5 bytes: ### L

	// Close the file when you’re done (usually this would be scheduled immediately after Opening with defer).
	defer f.Close()
}
