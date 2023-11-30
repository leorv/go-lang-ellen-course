package main

import (
	"fmt"
)

const (
	a = iota + 10
	b
	_
	d
	x
	y
	z
)

func main() {
	fmt.Println(a, b, d, x, y, z)
	// 10 11 13 14 15 16
}
