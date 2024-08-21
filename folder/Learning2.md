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

### 19) Methods:

```go
package main

import "fmt"

// Go supports methods defined on struct types.

type rect struct {
	width, height int
}

// This area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Here we call teh 2 methods defined for our struct.
	fmt.Println("Area:", r.area())
	// Area: 50
	fmt.Println("Perimeter:", r.perim())
	// Perimeter: 30

	// Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("area:", rp.area())
	// area: 50
	fmt.Println("peri:", rp.perim())
	// peri: 30
}
```

### 20) Interfaces:

```go
package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures

// Here’s a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// For our example we’ll implement this interface on rect and circle types.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to implement all the methods in the interface. Here we implement geometry on rects.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for circles.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call methods that are in the named interface. Here’s a generic measure function taking advantage of this to work on any geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 6, height: 4}
	c := circle{radius: 5}

	// The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
	measure(r)
	// {6 4}
	// 24
	// 20
	measure(c)
	// {5}
	// 78.53981633974483
	// 31.41592653589793
}
```

### 21) Enums:

```go
package main

import "fmt"

/*
	Enumerated types(enums) are a special case of sum types.
	An enum is a type that has a fixed number of possible values, each with a distant name.
	Go doesn't have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms.
*/

// Our enum type ServerState has an underlying int type.
type ServerState int

// The possible values for ServerState are defined as constants. The special keyword iota generates successive constant values automatically; in this case 0, 1,2 and so on.
const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

// By implementing the fmt.Stringer interface, values of ServerState can be printed out or converted to strings.
var stateName = map[ServerState]string{
	StateIdle: "idle",
	StateConnected: "connected",
	StateError: "error",
	StateRetrying: "retrying",
}
// This can get cumbersome if there are many possible values. In such cases the stringer tool can used in conjunction with go:generate to automate the process.

func (ss ServerState) String() string {
	return stateName[ss]
}

// If we have a value of type int, we cannot pass it to transition - the compiler will complain about type mismathc. This provides some degree of compile-time type safety for enums.
func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// connected

	ns2 := transition(ns)
	fmt.Println(ns2)
	// idle
}

// transition emulates a state transition for a server; it takes the existing state and returns a new state.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Suppose we check some predicated here to determine the next state...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
```

### 22) StructEmbedding:

```go
package main

import "fmt"

/*
	Go supports embedding of structs and itnerfaces to express a more seamless composition of types.
	This is not to be confused with //go:embed which is a go directive introduced in Go version 1.16+ to embed files and folders into the application binary.
*/

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base. An embedding looks like a field without a name.
type container struct {
	base
	str string
}

func main() {
	// When creating structs with literals, we have to initialize the embedding explicitly; here the embedded type serves as the field name.
	co := container{
		base: base{num: 1,},
		str: "some name",
	}

	// We can access the base's fields directly on co, e.g. co.num.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	// co={num: 1, str: some name}

	// Alternatively, we can spell out the full path using the embedded type name.
	fmt.Println("alos num:", co.base.num)
	// alos num: 1

	// Since container embeds base, the methods of base also become methods of a container. Here we invoke a method that was emedded from base directly on co.
	type describer interface {
		describe() string
	}
	// Embedding structs with methods may be used to bestow interface implementations onto other structs.
	// Here we see that a container now implements the describer interface because it embeds base.
	var d describer = co
	fmt.Println("describer:", d.describe())
	// describer: base with num=1
}
```

### 23) Generics:

```go
package main

import "fmt"

// Starting with version1.18, Go has added support for generics, also known as type parameters.

// As an example of a generic fuction, MapKeys takes a map of any type and returns a slice of its keys.
// This funciton has two type parameters - K and V; K as the comparable constraint, meaning that we can compare values of this type with the == and != operators.
// This is required for map keys in Go. V has the any constraint, meaning that it's not restricted in any way (any is an alias for interface{})
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// As an example of a generic type, List is a singly-linked list with values of any type.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

// We can define methods on generic types just like we do on regular types, but we have to keep the type parameters in place. The type is List[T], not List.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{ val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// When invoking generic functions, we can often rely on type inference.
	// Note that we don't have to specify teh types for K and V when calling MapKeys - the compiler infers them automatically.
	fmt.Println("Keys:", MapKeys(m))
	// [1 2 4]

	// ...though we could also specify them explicitly
	_ = MapKeys[int, string](m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
	// list: [10 13 23]
}
```

### 24) Errors:

```go
package main

import (
	"errors"
	"fmt"
)

// In Go it’s idiomatic to communicate errors via an explicit, separate return value.
// This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C.
// Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for other, non-error tasks.

// By convention, errors are the last return value and have type error, a built-in interface.
func f(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message.
		return -1, errors.New("Can't work with 42")
	}

	// A nil value in the error position indicates that there was no error.
	return arg + 3, nil
}

// A sentinel error is a predeclared variable that is used to signify a specific error condition.
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("cant boil water")

// We can wrap errors with higher-level errors to add context. The simplest way to do this is with the %w verb in fmt.Errorf. Wrapped errors create a logical chain (A wraps B, which wraps C, etc.) that can be queried with functions like errors.Is and errors.As.
func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7,42} {
		// It’s common to use an inline error check in the if line.
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			// errors.Is checks that a given error (or any error in its chain) matches a specific error value. This is especially useful with wrapped or nested errors, allowing you to identify specific error types or sentinel errors in a chain of errors.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Println("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}

	// f worked: 10
	// f failed: Can't work with 42
	// Tea is ready!
	// Tea is ready!
	// We should buy new tea!
	// Tea is ready!
	// Now it is dark.
}
```

### 25) Custom Errors:

```go
package main

import (
	"errors"
	"fmt"
)

// It’s possible to use custom types as errors by implementing the Error() method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.

// A custom error type usually has the suffix "Error".
type argError struct {
	arg int
	message string
}

// Adding this Error method makes argError implement the error interface.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		// Return our custom error.
		return -1, &argError{arg, "can't work with it."}
	}
	return arg +3, nil
}

// errors.As is a more advanced version of errors.Is. It checks that a given error (or any error in its chain) matches a specific error type and converts to a value of that type, returning true. If there’s no match, it returns false.
func main() {
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		// 42
		fmt.Println(ae.message)
		// can't work with it
	} else {
		fmt.Println("err doesn't match argError")
	}
}
```
