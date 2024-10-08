package main

import (
	"fmt"
	"os"
)

// Go offers excellent support for stirng formatting in the printf tradition.
// Here are some examples of common string formatting tasks.
type point struct {
	x, y int
}

func main() {
	// Go offers several printng "verbs" designed to format general Go values.
	// For example, this prints an instance of our point struct.
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)
	// struct1: {1 2}

	// If the value is a struct, the %+v variant will include the struct's field anmes.
	fmt.Printf("struct2: %+v\n", p)
	// struct2: {x:1 y:2}

	// The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.
	fmt.Printf("struct3: %#v\n", p)
	// struct3: main.point{x:1, y:2}

	// To print the type of a value, use %T.
	fmt.Printf("type: %T\n", p)
	// type: main.point

	// Formatting booleans is straight-forward.
	fmt.Printf("bool: %t\n", true)
	// bool: true

	// There are many options for formatting integers. Use %d for standard, base-10 formatting.
    fmt.Printf("int: %d\n", 123)
	// 	int: 123
	
	// This prints a binary representation.
    fmt.Printf("bin: %b\n", 14)
	// bin: 1110
	
	// This prints the character corresponding to the given integer.
    fmt.Printf("char: %c\n", 33)
	// char: !
	
	// %x provides hex encoding.
    fmt.Printf("hex: %x\n", 456)
	// hex: 1c8
	
	// There are also several formatting options for floats. For basic decimal formatting use %f.
    fmt.Printf("float1: %f\n", 78.9)
	// float1: 78.900000
	
	// %e and %E format the float in (slightly different versions of) scientific notation.
    fmt.Printf("float2: %e\n", 123400000.0)
	// float2: 1.234000e+08
    fmt.Printf("float3: %E\n", 123400000.0)
	// float3: 1.234000E+08
	
	// For basic string printing use %s.
    fmt.Printf("str1: %s\n", "\"string\"")
	// str1: "string"
	
	// To double-quote strings as in Go source, use %q.
    fmt.Printf("str2: %q\n", "\"string\"")
	// str2: "\"string\""
	
	// As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.
    fmt.Printf("str3: %x\n", "hex this")
	// str3: 6865782074686973
	
	// To print a representation of a pointer, use %p.
    fmt.Printf("pointer: %p\n", &p)
	// pointer: 0x140000a0020
	
	// When formatting numbers you will often want to control the width and precision of the resulting figure. To specify the width of an integer, use a number after the % in the verb. By default the result will be right-justified and padded with spaces.
    fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	// width1: |    12|   345|
	
	// You can also specify the width of printed floats, though usually you’ll also want to restrict the decimal precision at the same time with the width.precision syntax.
    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	// width2: |  1.20|  3.45|
	
	// To left-justify, use the - flag.
    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	// width3: |1.20  |3.45  |
	
	// You may also want to control width when formatting strings, especially to ensure that they align in table-like output. For basic right-justified width.
    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	// width4: |   foo|     b|
	
	// To left-justify use the - flag as with numbers.
    fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
	// width5: |foo   |b     |
	
	// So far we’ve seen Printf, which prints the formatted string to os.Stdout. Sprintf formats and returns a string without printing it anywhere.
    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)
	// sprintf: a string
	
	// You can format+print to io.Writers other than os.Stdout using Fprintf.
    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
	// io: an error
}