### 61) SHA256 Hashes:
```go
package main

// Go implements several hash functions in various crypto/* packages.
import (
	"crypto/sha256"
	"fmt"
)

// SHA256 hashes are frequently used to compute short identities for binary or text blobs. For example, TLS/SSL certificates use SHA256 to compute a certificate's signature. Here's how to compute SHA256 hashes in Go.

func main() {
	s := "sha256 this is sgtring"

	// Here we start with a new hash
	h := sha256.New()

	// Write expects bytes. If you hace a string s, use []byte(s) to coerce it to bytes
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice. THe argument to Sum can be used to append to an exisitng byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	fmt.Println(s)
	// sha256 this is sgtring
	fmt.Printf("%x\n", bs)
	// 3a97e70165a808c6d867ecb3d250de8712822c619aeedd6dd7f7794117b37a16
}
```

### 62) Base64 Encoding:
```go
package main

// This syntax imports the encoding/base64 package with the b64 name instead of the default base64. It’ll save us some space below.
import (
	b64 "encoding/base64"
	"fmt"
)

// Go provides built-in support for base64 encodign/decoding

func main() {
	// Here's the string we'll encode/decode
	data := "prashant1234567890"

	// Go supports both standard and URL-compatible base64. Here’s how to encode using the standard encoder. The encoder requires a []byte so we convert our string to that type.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	// cHJhc2hhbnQxMjM0NTY3ODkw

	// Decoding may return an error, which you can check if you don’t already know the input to be well-formed.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	// prashant1234567890

	// This encodes/decodes using a URL-compatible base64 format
	// uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	// cHJhc2hhbnQxMjM0NTY3ODkw
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
	// prashant1234567890
}
```

### 63) Reading Files:
```go
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
	// The total number of bytes read is stores in n1
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
```

### 64) Writing Files:
```go
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
```

### 65) Line Filters:
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A line filter is a common type of program that reads input on stdin, process it, and then prints some derived result to stdout. grep and sed are common line filters.
// Here's an example line filter in Go that writes a capitalized version of all input text. You can use this pattern to write your own Go line filters.

func main() {
	// Wrapping the unbuffered os.Stdin with a buffered scanner giver us a convernient Scan method that advances the scanner to the next token; which is the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)

	// Text returns the current token, here the next line, from the input.
	for scanner.Scan() {
		// Write out the uppercased line.
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// Check for errors during Scan. End of file is expected and not reported by Scan as an error.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "errors:", err)
		os.Exit(1)
	}
}
```

### 66) File Paths:
```go
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// The filepath package provides functions to parse and construct file paths in a way that is portable between operating systems; dir/file on Linux vs dir\file on Windows, for example>

func main() {
	// Join should be used to construct paths in a portable way. It takes any number of arguments and constructs a hierarchical path from them.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)
	// p: dir1/dir2/filename

	// You should always use Join instead of concatenating /s or \s manually. In addition to providing portability, Join will also normalize paths by removing superfluous separators and directory changes.
	fmt.Println(filepath.Join("dir1//", "filename"))
	// dir1/filename
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	// dir1/filename

	// Dir and Base can be used to split a path to the directory and the file. Alternatively, Split will return both in the same call.
	fmt.Println("Dir(p):", filepath.Dir(p))
	// Dir(p): dir1/dir2
	fmt.Println("Base(p):", filepath.Base(p))
	// Base(p): filename

	// We can check whether a path is absolute.
	fmt.Println(filepath.IsAbs("dir/file"))
	// false
	fmt.Println(filepath.IsAbs("/dir/file"))
	// true

	filename := "config.json"
	// Some file names have extensions following a dot. We can split the extension out of such names with Ext.
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	// .json

	// To find the file’s name with the extension removed, use strings.TrimSuffix.
	fmt.Println(strings.TrimSuffix(filename, ext))
	// config

	// Rel finds a relative path between a base and a target. It returns an error if the target cannot be made relative to base.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	// t/file
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	// ../c/t/file
}
```

### 67) Direcotries:
```go
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Go has several useful functions for working with directories in the file system.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Check a new sub-directory in the current working directory.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// When creating temp directories, it's good practice to defer their removal. os.RemoveAll will delete a whole directory tree (similarly to rm -rf)
	defer os.RemoveAll("subdir")

	// Helper function to create a new empty file
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// We can create a heirarchy of directories, including parent with MkdirAll. This is similar to the command-line mkdir -p.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir lists directory contents, returning a slice of os.DirEntry objects.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	// listing subdir/parent
	//   child true
	//   file2 false
	//   file3 false

	// Chdir lets us change the current working directory, similarly to cd.
	err = os.Chdir("subdir/parent/child")

	// Now we'll see the contents of subdir/parent/child when listing the current directory.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	// listing subdir/parent/child
	//   file4 false

	// cd back to where we started.
	err = os.Chdir("../../..")
	check(err)

	// We can also visit a directory recursively, including all its sub-directories. WalkDir accepts a callback function to handle every file or directory visited.
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
	// Visiting subdir
	//   subdir true
	//   subdir/file1 false
	//   subdir/parent true
	//   subdir/parent/child true
	//   subdir/parent/child/file4 false
	//   subdir/parent/file2 false
	//   subdir/parent/file3 false
}

// visit is called for every file or directory found recursively by filepath.WalkDir.
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
```

### 68) Temporary Files and Directories:
```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Throughout program execution, we often want to create data that isn't needed adter the program exits. Temporary files and directories are useful for this purpose since they don't pollute the file sustem over time.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// The easiest way to create a temporary file is by calling os.CreateTemp. It creates a file and opens it for reading and writing. We provide "" as the first argument, so os.CreateTemp will create the file in the default location for our OS.
	f, err := os.CreateTemp("", "sample")
	check(err)

	// Display the name of the temporary file. On Unix-based OSes the directory will likely be /tmp. The file name starts with the prefix given as the second argument to os.CreateTemp and the rest is chosen automatically to ensure that concurrent calls will always create different file names.
	fmt.Println("Temp file name:", f.Name())
	// Temp file name: /var/folders/kk/<something...>/T/sample3491378737

	// Cleanup teh file after wee're done. The OS is likely to clean up temporary files by itself after sometime, but it's good practice to do this explicitly
	defer os.Remove(f.Name())

	// We can write soem data to the file.
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// If we intend to write many temporary files, we may prefer to create a temporary directory. os.MkdirTemp's arguments are the same as CreateTemp's, but it returns a directory name rather than an open file.
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)
	// Temp dir name: /var/folders/kk/<something....>/T/sampledir3522086750

	defer os.RemoveAll(dname)

	// Now we can synthesize temporary file names by prefixing them with our temporary directory.
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
```

### 69) Embed Directive:
```go
package main

// `//go:embed` is a compiler directoive taht allows programs to include arbitrary files and folders in the Go binary at build time.

import (
	"embed"
)

// embed directives accept paths relative to the directory containing the Go source file. This directive embeds the cotents of the file into the string variable immediately following it.
//
//go:embed README.md
var fileString string

// or embed the contents of the file into a []byte.
//
//go:embed README.md
var fileByte []byte

// We can also embed multiple files or even folders with wildcards. THis uses a variable of the embed FS type, which implements a simple virtual file syste.
//
//go:embed folder/Learning.md
//go:embed folder/*.md
var folder embed.FS

func main() {
	// Print out the contents of README.md
	print(fileString)
	print(string(fileByte))
	// ### Learning Go by Practice

	// This is the practice ground for Prashant to test his GoLang code.

	// Retrieve some files from the embedded folder.
	content1, _ := folder.ReadFile("folder/Learning.md")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/Learning2.md")
	print(string(content2))
}
```
