package main

import "github.com/01-edu/z01"

func main() {
	numero := "0123456789"
	for i := 0; i < len(numero); i++ {
		z01.PrintRune(rune(numero[i]))
	}
	z01.PrintRune('\n')
}
