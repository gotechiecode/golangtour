package main

import (
	"fmt"
)

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	// b is a pointer to the memory address of variable 'a' where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int
	var b = &a

	fmt.Println(b)
}
