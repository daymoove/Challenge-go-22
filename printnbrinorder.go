package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	o := n
	l := 0
	e := 1
	r := 0
	var m []int
	if n < 0 {
	} else {
		for o > 9 {
			o = o / 10
			l++
		}
		o = n
		for i := 0; i <= l; i++ {
			if o > 9 {
				for o > 9 {
					o = o / 10
					e = e * 10
				}
				m = append(m, o)
				r = r + o*e
				o = n - r
				if o < e/10 {
					m = append(m, 0)
				}
				e = 1
			} else {
				m = append(m, o)
			}
			if m[i] == 0 {
				i++
			}
		}
		for k := 0; k <= 9; k++ {
			for c := 0; c <= len(m)-1; c++ {
				if m[c] == k {
					z01.PrintRune(rune(m[c] + 48))
				}
			}
		}
	}
}
