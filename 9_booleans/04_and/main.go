package main

import "fmt"

func main() {

	if true && false {
		fmt.Println("Didn't get executed")
	}

}
