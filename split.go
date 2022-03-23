package piscine

func Split(s, sep string) []string {
	l := []rune(s)
	z := []rune(sep)
	var table []string
	var m string
	c := 0
	for i := 0; i < len(s); i++ {
		if l[i] == z[0] {
			for a := 0; a < len(sep); a++ {
				if l[a+i] == z[a] {
					c++
				} else {
					m = m + string(l[i])
					break
				}
				if c == len(sep) {
					i = i + len(sep) - 1
					table = append(table, m)
					m = ""
					c = 0
				}
			}
		} else {
			m = m + string(l[i])
		}
		c = 0
	}
	table = append(table, m)
	return table
}
