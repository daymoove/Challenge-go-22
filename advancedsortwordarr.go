package piscine

func AdvancedSortWordArr(b []string, f func(a, c string) int) {
	for i := 1; i < len(b); i++ {
		for a := 0; a < len(b); a++ {
			if b[a] > b[i] {
				v := b[i]
				b[i] = b[a]
				b[a] = v
			}
		}
	}
}
