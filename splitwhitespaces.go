package piscine

func SplitWhiteSpaces(s string) []string {
	l := []rune(s)
	var table []string
	var m string
	for i := 0; i < len(s); i++ {
		if l[i] == 32 || l[i] == '\n' || l[i] == 9 {
			if l[i+1] == 32 || l[i+1] == '\n' || l[i+1] == 9 {
			} else {
				table = append(table, m)
				m = ""
			}
		} else {
			m = m + string(l[i])
		}
	}
	table = append(table, m)
	return table
}
