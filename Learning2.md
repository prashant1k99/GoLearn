### 17) Runes:

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

// A Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of “characters”. In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.

func main() {
	// s is a string assigned a literal value representing the word “hello” in the Thai language. Go string literals are UTF-8 encoded text.
	const s = "สวัสดี"

	// Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.
	fmt.Println("Len:", len(s))
	// 18

	// Indexing into a string produces the raw byte values at each index. This loop generates the hex values of all the bytes that constitute the code points in s.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
		// e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
	}
	fmt.Println()

	// To count how many runes are in a string, we can use the utf8 package. Note that the run-time of RuneCountInString depends on the size of the string, because it has to decode each UTF-8 rune sequentially. Some Thai characters are represented by UTF-8 code points that can span multiple bytes, so the result of this count may be surprising.
	fmt.Println("Rune Count:", utf8.RuneCountInString(s))
	// Rune Count: 6

	// A range loop handles strings specially and decodes each rune along with its offset in the string.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
		// U+0E2A 'ส' starts at 0
		// U+0E27 'ว' starts at 3
		// U+0E31 'ั' starts at 6
		// U+0E2A 'ส' starts at 9
		// U+0E14 'ด' starts at 12
		// U+0E35 'ี' starts at 15
	}

	// We can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, width)
		w = width

		// This demonstrates passing a rune value to a function.
		examineRune(runeValue)
	}
	// Using DecodeRuneInString
	// U+0E2A 'ส' starts at 3
	// found so sua
	// U+0E27 'ว' starts at 3
	// U+0E31 'ั' starts at 3
	// U+0E2A 'ส' starts at 3
	// found so sua
	// U+0E14 'ด' starts at 3
	// U+0E35 'ี' starts at 3

}

func examineRune(r rune) {
	// Values enclosed in single quotes are rune literals. We can compare a rune value to a rune literal directly.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
        fmt.Println("found so sua")
	}
}
```

### 18) Structs;

```go
package main

import "fmt"

// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

// This person struct type has name and age fields.
type person struct {
	name string
	age int
}

// newPerson constructs a new person struct with the given name.
func newPerson(name string) *person {
	// Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// This syntax created a new struct
	fmt.Println(person{"Bob", 20})
	// {Bob 20}

	// You can name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 20})
	// {Alice 20}

	// Omited fields will be zero-valued
	fmt.Println(person{name: "Fred"})
	// {Fred 0}
	fmt.Println(person{age: 30})
	// { 30}

	// An & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 35})
	// &{Ann 35}

	// It's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))
	// &{Jon 42}

	// Access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	// Sean

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)
	// 50

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)
	// 51

	// If a struct type is only used for a single value, we don't have to give it a name.
	// The value can have an anonymous struct type. This technique is commonly used for tabledriven tests
	dog := struct {
		name string
		isGood bool
	} {
		"Rex",
		true,
	}
	fmt.Println(dog)
	// {Rex true}
}
```
