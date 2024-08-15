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