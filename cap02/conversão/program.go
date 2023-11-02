package main

import (
	"fmt"
)

type hotdog int

var b hotdog

func main() {
	b = 12
	fmt.Printf("%v\n", b)

	a := 3
	fmt.Printf("%v\n", a)

	a = int(b)
	fmt.Printf("%v\n", a)
}
