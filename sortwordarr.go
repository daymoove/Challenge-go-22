package piscine

func SortWordArr(b []string) {
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
