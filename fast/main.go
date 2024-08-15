package main

import ("fmt"; "math")

func main() {
	var i int = 10
	i = 20
	x := 20

	// Different primary datatype in Go
	// int, string, bool, float

	if x < i {
		fmt.Println("X is smaller than I")
	} else if x == i {
		fmt.Println("X is equal to I")
	} else {
		fmt.Println("X is greater than I")
	}

	fmt.Println(x)

	// Array
	var arr [5]int // Showing that arr is the variable name and is assigned with the fixed size of 5 and datatype of string
	// [0, 0, 0, 0, 0]
	arr1 := [5]int {1, 2, 3, 4, 5} // Alternative way to assign fix length Array
	// [1,2,3,4,5]
	fmt.Println(arr)
	fmt.Println(arr1)
	arr1[4] = 6
	// [1,2,3,4,6] || At index 4 replace with 6
	fmt.Println(arr1)
	// You cannot perform Append to this Array as it's of fix size
	
	// To create arrays without any fix length, we use slice
	sli := []int {1,2,3,4,5}
	fmt.Println(sli) // [1,2,3,4,5]
	
	sli = append(sli, 6)
	fmt.Println(sli) // [1,2,3,4,5,6]

	// Creating Maps in Go
	a := make(map[string]int) // So we are creating variable a with an empty map, where key will be string and value should be integer
	a["test"] = 10
	fmt.Println(a) // map[test:10]
	a["test1"] = 15
	a["test2"] = 20

	delete(a, "test") // We can delete any particular key value pair in Map using delete function
	fmt.Println(a) // map[test1:15 test2:20]

	// Loop: Go only has for loop and using for loop we can create different types of loops
	for k := 0; k < 10; k++ {
		fmt.Println(k)
	} // 0 \n 1 \n 2 \n 3 \n 4 \n 5 \n 6 \n 7 \n 8 \n 9

	// To make while loop
	k1 := 0
	for k1 < 10 {
		k1++
	} // Works same as while loop

	// To have endless running loop
	for {
		if k1 > 100 {
			fmt.Println("Breaking endless loop", k1)
			break //  we have both break and continue in Go
		}
		k1++
	}

	val := add(10, 20)
	fmt.Println(val) // 30

	// There can be multiple return types of a function
	a2, err := sqrt(2) // We handle all return params here
	if err != nil {
		panic("") 
		/*  panic:

			goroutine 1 [running]:
			main.main()
					E:/GoLang/fast/main.go:78 +0x68a
			exit status 2 
		*/
	}
	fmt.Println(a2)

	// We can have named returns
	_, error := divide(25, 5) // For returns we do not wish to handle, we can replace with _
	if error != nil {
		fmt.Println("You cannot perform division.")
	}
}

func add(a int, b int) int { // We define a function `add`, we have params `a` and `b` with the type int. And Return type int
	return a + b
}

func sub(a, b int) int { // We can have also define datatype for all variables in one go
	return a -b
}

func someWeirdFn(a, b int, c string, d bool, e int) {
	// Here we have multiple types of params and return type is set to void as no return datatype is defined
}

func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("Cannot get sqrt for values less than 0")
	}

	return math.Sqrt(a), nil
}

func divide(x1, x2 float64) (res float64, err error) {
	if x2 == 0 {
		return 0.0, fmt.Errorf("Cannot divide with 0")
	}
	return x1 / x2, nil
}
