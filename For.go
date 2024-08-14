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