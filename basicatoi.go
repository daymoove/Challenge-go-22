package piscine

func BasicAtoi(s string) int {
	o := 0
	c := 0
	r := []rune(s)
	for _, word := range r {
		for i := '0'; i < word; i++ {
			c++
		}
		o = o*10 + c
		c = 0
	}
	return o
}
