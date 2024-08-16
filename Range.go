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