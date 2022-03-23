package piscine

func SortIntegerTable(table []int) {
	j := 0
	l := len(table)
	for range table {
		if table[j] > table[l] {
			table[j] = table[l]
		} else {
			j = j + 1
		}
	}
}
