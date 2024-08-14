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
