package main

import (
	"fmt"
	s "strings"
)

// The standard library's strings package provides many useful string-related functions. Here are some examples to give you a sense of the package.

// We alias fmt.Println to a shorter name as we'll use it a log below
var p = fmt.Println

func main() {
	// Here's a sample of the functions available in strings. Since these are functions from the package, not methods on the string object itself, we need pass the string in question as the first argument to the function.
	p("Contains:  ", s.Contains("test", "es"))
	// true
    p("Count:     ", s.Count("test", "t"))
	// 2
    p("HasPrefix: ", s.HasPrefix("test", "te"))
	// true
    p("HasSuffix: ", s.HasSuffix("test", "st"))
	// true
    p("Index:     ", s.Index("test", "e"))
	// 1
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	// a-b
    p("Repeat:    ", s.Repeat("a", 5))
	// aaaaa
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
	// f00
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
	// f0o
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
	// [a b c d e]
    p("ToLower:   ", s.ToLower("TEST"))
	// test
    p("ToUpper:   ", s.ToUpper("test"))
	// TEST
}