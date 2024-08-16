# GoLang Notes

To run Go project, we must include main package to execute: [main function is entry point to any Go code]

```go
package main

func main() {
    // ... your code here
}
```

To execute the go file:

```sh
go run HelloWorld.go
```

### 1) Print Hello World!

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

- To build it as an executable

```sh
go build HelloWorld.go
```

### 2) Values in Go:

```go
    fmt.Println("go" + "lang") // golang

	fmt.Println("1+1=", 1+1) // 1+1=2
	fmt.Println("7.0/3.0=", 7.0/3.0) // 7.0/3.0=2.3333333333333335

	fmt.Println(true && false) // false [if true, then false]
	fmt.Println(true || false) // true [true or false]
	fmt.Println(!true) // false [Not True]
```

### 3) Variables:

```go
	var a = "initial"
	fmt.Println("a=", a) // a= initial

	var b, c int = 1, 2
	fmt.Println(b, c) // 1 2 || Intialize 2 variables b and c with value 1 and 2 and set datatype int

	var d = true
	fmt.Println(d) // true

	var e int
	fmt.Println(e) // 0 || Default value for int is 0, string is <empty>, bool is false


	f := "apple"
	fmt.Println(f) // apple || shorthand Initialization and assignment expression [var f string = "apple"]
```

### 4) Constants:

```go
package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s) // constant || Accessible from outside of main func

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d) // 6e+11 || the value of d

	fmt.Println(int64(d)) // 600000000000 || Here it had no type but the type is given as int64 by an explicit conversion

	fmt.Println(math.Sin(n)) // -0.28470407323754404 || A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
}
```

### 5) For Loop:

```go
package main

import "fmt"

func main() {
	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	} // 1 \n 2 \n 3

	// The most basic type, with a single condition.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	} // 0 \n 1 \n 2

	// Another way of accomplishing the basic “do this N times” iteration is range over an integer.
	for l := range 3 {
		fmt.Println("range", l)
	} // range 0 \n range 1 \n range 2

	// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	// in this example we are looping infinetly, to add a break condition we are breaking the loop after 20 runs and keep on incrementing the variable
	k := 0
	for {
		fmt.Println("looping")
		k = k + 1
		if k == 20 {
			break
		}
	} // looping (x20)

	// You can also continue to the next iteration of the loop.
	for n := range 6 {
		if n%2 == 0   {
			continue
		}
		fmt.Println(n)
	} // 1 \n 3 \n 5
}
```

### 5) If/Else:

```go
package main

import "fmt"

func main() {
	// Here’s a basic example.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	} // 7 is odd

	// You can have an if statement without an else
	if 8%2 == 0 {
		fmt.Println("8 is divisible by 4")
	} //8 is divisible by 4

	// Logical operations like `&&` and `||` are often useful in conditions
	if (8%2 == 0 || 7 %2 == 0) {
		fmt.Println("Either 8 or 7 is even")
	} //Either 8 or 7 is even

	// A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.
	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    } // 9 has 1 digit
}

// NOTE: You can add parantheses around conditions but are not required in Go, but that the braces are required
```

### 6} Switch:

```go
// Switch statements express conditionals across many branches
package main

import (
	"fmt"
	"time"
)

func main() {
	// Here’s a basic switch.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	} // Write 2 as Two

	// You can use commas to separate multiple expressions in the same case statement. We use the optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	} // It's the weekend

	// switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	} // It's after noon

	// A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause.

	// A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true) // I'm a bool
	whatAmI(1) // I'm an int
	whatAmI("hey") //Don't know type string
}
```

### 7) Arrays:

```go
package main

import "fmt"

func main() {
	// Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for ints means 0s.
	var a [5]int
	fmt.Println("emp:", a)
	// emp: [0 0 0 0 0]

	// We can set a value at an index using the array[index] = value syntax, and get a value with array[index].
	a[4] = 100
	fmt.Println("set:", a) //set: [0 0 0 0 100]
	fmt.Println("get:", a[4]) // get: 100

	// The builtin len returns the length of an array.
	fmt.Println("len:", len(a))
	// len: 5

	// Use this syntax to declare and initialize an array in one line.
	b := [5]int {1,2,3,4,5}
	fmt.Println("dcl:", b)
	// dcl: [1 2 3 4 5]

	// You can also have the compiler count the number of elements for you with ...
	b = [...]int{1,2,3,4,5}
	fmt.Println("idx:", b)
	// idx: [1 2 3 4 5]

	// If you specify the index with :, the elements in between will be zeroed.
	b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)
	// idx: [100 0 0 400 500]

	// Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	var twoD [2][3]int
	for i :=0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d: ", twoD)
	// 2d:  [[0 1 2] [1 2 3]]

	// You can create multidimensional Arrays at once too
	twoD = [2][3]int{
		{1,2,3},
		{4,5,6},
	}
	fmt.Println("2d: ", twoD)
	// 2d:  [[1 2 3] [4 5 6]]
}
```

### 8) Slices:

```go
package main

import (
	"fmt"
	"slices"
)

// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.
func main() {
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements). An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	// uninit: [] true true

	// To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued). By default a new slice’s capacity is equal to its length; if we know the slice is going to grow ahead of time, it’s possible to pass a capacity explicitly as an additional parameter to make.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	// emp: [  ] len: 3 cap: 3

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	// set: [a b c]
    fmt.Println("get:", s[2])
	// get: c

	// In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)
	// [a b c d e f]

	// Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s) // Copy to from
	fmt.Println("Copy:", c)
	// Copy: [a b c d e f]

	// Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].
	// From to[excluding] index
	l := s[2:4]
	fmt.Println("sl1:", l)
	// sl1: [c d]

	// This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sli2:", l)
	// sli2: [a b c d e]

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sli3:", l)
	// sli3: [c d e f]

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string{"a","b","c"}
	fmt.Println("dcl:", t)
	// dcl: [a b c]

	// The slices package contains a number of useful utility functions for slices.
	t2 := []string{"a", "b","c"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}
	// t == t2

	// Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i +1
		twoD[i] = make([]int, innerLen) // For that particular index, we set the length of the slice
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
	// twoD: [[0] [1 2] [2 3 4]]
}
```

### 9) Map:

```go
package main

import (
	"fmt"
	"maps"
)

// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
func main() {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 1
	m["k2"] = 2

	// Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("Map:", m)
	// Map: map[k1:1 k2:2]

	// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	// v1: 1

	// If the key doesn’t exist, the zero value of the value type is returned.
	v3 := m["k3"]
	fmt.Println("v3:", v3)
	// v3: 0

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))
	// len: 2

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k1")
	fmt.Println("Del:", m)
	// Del: map[k2:2]

	// To remove all key/value pairs from a map, use the clear builtin.
	clear(m)
	fmt.Println("Clr:", m)
	// Clr: map[]

	// The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	// prs: false

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map:", n)
	// map: map[bar:2 foo:1]

	// The maps package contains a number of useful utility functions for maps.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
		// n == n2
	}
}
```

### 10) Range:

```go
package main

import "fmt"

// range iterates over elements in a variety of data structures. Let’s see how to use range with some of the data structures we’ve already learned.
func main() {
	// Here we use range to sum the numbers in a slice. Arrays work like this too.
	nums:= []int{1,2,3}
	sum := 0
	for _, num := range nums {
        sum += num
    }
	fmt.Println("sum:", sum)
	// sum: 6

	// range on arrays and slices provides both the index and value for each entry. Above we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
			// index: 2
		}
 	}

	// range on map iterates over key/value pairs.
	kvs := map[string]string{"foo":"bar", "john": "doe"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// foo -> bar
	// john -> doe

	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
		// key: foo
		// key: john
	}

	// range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself. See Strings and Runes for more details.
	for i, c := range "go" {
        fmt.Println(i, c)
		// 0 103
		// 1 111
    }

	for i, c := range "😆😀😇" {
        fmt.Println(i, c)
		// 0 128518
		// 4 128512
		// 8 128519
    }
}
```

### 11) Functions:

```go
package main

import "fmt"

// Functions are central in Go. We’ll learn about functions with a few different examples.

// Here’s a function that takes two ints and returns their sum as an int.
func plus(a int, b int) int {
	// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
    return a + b
}

// When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {

	// Call a function just as you’d expect, with name(args).
    res := plus(1, 2)
    fmt.Println("1+2 =", res)
	// 1+2 = 3

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
	// 1+2+3 = 6
}
```

### 12) Multiple Return Functions:

```go
package main

import "fmt"

// Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

// The (int, int) in this function signature shows that the function returns 2 ints.
func vals() (int, int) {
	return 2, 3
}

func main() {
	// Here we use the 2 different return values from the call with multiple assignment.
	a, b := vals()
	fmt.Println(a)
	// 2
	fmt.Println(b)
	// 3

	// If you only want a subset of the returned values, use the blank identifier _.
	_, c := vals()
	fmt.Println(c)
	// 3
}
```

### 13) Variadic Functions:

```go
package main

import "fmt"

// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function

// Here’s a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Println(nums, " ")

	total := 0
	// Within the function, the type of nums is equivalent to []int. We can call len(nums), iterate over it with range, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println("Sum:", total)
}

func main() {
	// Variadic functions can be called in the usual way with individual arguments.
	sum(1,2,3,4,5,6,7)
	// Sum: 28

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1,2,3,64756,34,53467,486,4345}
	sum(nums...)
	// Sum: 123094
}
```
