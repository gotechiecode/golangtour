package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20818a220

	// b is an int pointer, which points to the memory address of variable 'a' where an int is stored
	var b = &a
	fmt.Println(b)  // 0x20818a220
	// to see the value of b, add a * in front of b, this is known as dereferencing
	fmt.Println(*b) // 43
}
