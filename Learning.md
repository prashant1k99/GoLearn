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

### Arrays:

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
