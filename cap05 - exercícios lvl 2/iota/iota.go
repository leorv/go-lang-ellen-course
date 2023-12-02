package main

import "fmt"

const ano = 2023

const (
	a = iota + ano
	b
	c
	d
)

func main() {
	fmt.Printf("próximos 3 anos: %v, %v, %v, %v.\n", a, b, c, d)
}
