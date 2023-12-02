package main

import "fmt"

func main() {
	var a int = 200
	fmt.Printf("decimal: %d, exadecimal: %#x, binário: %b\n", a, a, a)

	b := (a << 1)
	fmt.Println("deslocando 1 bit para a esquerda...")
	fmt.Printf("%b\n", b)
}
