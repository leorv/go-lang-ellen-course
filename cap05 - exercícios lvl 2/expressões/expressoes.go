package main

import "fmt"

func main() {
	a := 1 == 10
	b := (2 != 3)
	c := (2 <= 3)
	d := (2 < 3)
	fmt.Printf("%v\n%v\n%v\n%v\n", a, b, c, d)
}
