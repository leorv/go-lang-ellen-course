package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	a := 12
	b = 12
	fmt.Println(b, a)

	// Apesar de serem baseados em int, a e b não podem ser somados, por exemplo. São tipos diferentes.
	// O tipo subjacente dele é um int.
}
