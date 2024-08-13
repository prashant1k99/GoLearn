package main

import "fmt"

func main() {
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
}