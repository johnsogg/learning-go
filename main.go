package main

import (
	"fmt"
)

func foo() {
	var a string = "initial"
	fmt.Println(a)
	var x, y int = 2, 6
	fmt.Println("2 + 6 is", x+y)
	fmt.Println("pi is roughly", y)
}
