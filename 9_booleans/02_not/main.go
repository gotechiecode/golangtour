package main

import "fmt"

func main() {

	if !true {
		fmt.Println("This is not printed")
	}

	if !false {
		fmt.Println("This is printed")
	}

}
