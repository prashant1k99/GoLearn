package main

import "fmt"

/*
	Go supports embedding of structs and itnerfaces to express a more seamless composition of types.
	This is not to be confused with //go:embed which is a go directive introduced in Go version 1.16+ to embed files and folders into the application binary.
*/

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base. An embedding looks like a field without a name.
type container struct {
	base
	str string
	num int
}

func main() {
	// When creating structs with literals, we have to initialize the embedding explicitly; here the embedded type serves as the field name.
	co := container{
		base: base{num: 1,},
		str: "some name",
		// num: 20,
	}

	// We can access the base's fields directly on co, e.g. co.num.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	// co={num: 1, str: some name}
	// co={num: 20, str: some name} || if the same num exists on container

	// Alternatively, we can spell out the full path using the embedded type name.
	fmt.Println("alos num:", co.base.num)
	// alos num: 1

	// Since container embeds base, the methods of base also become methods of a container. Here we invoke a method that was emedded from base directly on co.
	type describer interface {
		describe() string
	}
	// Embedding structs with methods may be used to bestow interface implementations onto other structs.
	// Here we see that a container now implements the describer interface because it embeds base.
	var d describer = co
	fmt.Println("describer:", d.describe())
	// describer: base with num=1
}