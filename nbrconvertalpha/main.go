package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	o := 0
	c := 0
	a_s := []rune(s)
	if len(s) == 0 {
	} else {
		for _, word := range a_s {
			for i := '0'; i < word; i++ {
				c++
			}
			o = o*10 + c
			c = 0
		}
	}
	return o
}

func main() {
	nlet := os.Args
	carnb := 0
	if len(nlet) > 1 {
		size := 96
		if nlet[1] == "--upper" {
			size = 64
		}
		for index, elems := range nlet {
			if index > 0 {
				carnb = Atoi(elems) + size
				if carnb >= 97 && carnb <= 122 || carnb >= 65 && carnb <= 90 {
					z01.PrintRune(rune(carnb))
				} else if nlet[1] == "--upper" {
					continue
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
