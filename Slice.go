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