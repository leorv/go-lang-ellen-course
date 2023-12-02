/* Escreva um programa que mostre um número em decimal, binário, e hexadecimal.*/
package main

import "fmt"

func main() {
	numero := 20

	fmt.Printf("%d\t%b\t%#x\n", numero, numero, numero)
}
