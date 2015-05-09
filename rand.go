package main

import (
	"fmt"
)

/*
#include <stdlib.h>
#include "foo.h"
*/
import "C"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	fmt.Println("Hello")
	fmt.Println("Canned number is", 4)
	fmt.Println("Random number from C is", Random())
	Seed(42)
	fmt.Println("After seeding with 42, random number from C is", Random())
	life := C.meaningOfLife()
	fmt.Println("The meaning of life is", life)
}
