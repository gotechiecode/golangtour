package main

import "fmt"

const p = "Im Global"

func main() {

	const q = 42 //local constant

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}

// a CONSTANT is a simple unchanging value
