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