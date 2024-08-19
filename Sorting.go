package main

import (
	"fmt"
	"slices"
)

// Go's slices pacakge implements sorting for builtins and user-defined types. We'll look at sorting for builtins first.

func main() {
	// Sorting functions are generic, and work for any ordered built-in type.
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("strings:", strs)
	// strings: [a b c]

	// An example of sorting ints
	ints := []int{7,5,2,9,100}
	slices.Sort(ints)
	fmt.Println("ints:", ints)
	// ints: [2 5 7 9 100]
	
	// We can also use the slices package to check if a slice is already in sorted order.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted:", s)
	// Sorted: true
}