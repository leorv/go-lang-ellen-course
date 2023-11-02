package main

import (
	"fmt"
)

func main() {
	// string interpretada
	x := "oi bom dia\ncomo vai?\tespero que tudo bem."

	// string literal
	y := `oi bom dia\ntudo bem?`

	fmt.Println(x)
	fmt.Println(y)

	string1 := "oi"
	string2 := "bom dia"

	z := fmt.Sprint(string1, string2)

	fmt.Print(z)
}
