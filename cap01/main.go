package main

import (
	"fmt"
)

var a int

func main() {
	_, erros := fmt.Println("Hello, world!")
	fmt.Println(erros)

	numeroBytes, _ := fmt.Println("Olá, mundo!")
	fmt.Println(numeroBytes)

	x := 5
	y := "strings"
	z := true

	x = 44
	a = 11

	fmt.Println(a, x, y, z)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
}
