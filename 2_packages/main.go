package main

import (
	"fmt"
	"github.com/gotechiecode/golangtour/2_packages/publicpack"
)

func main()  {
	fmt.Println(publicpack.MyName)
	//fmt.Println(privatepack.name) If we enable this comment, it gives compilation error. As the name variable starts with a small letter
}
