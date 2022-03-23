package piscine

func AppendRange(min, max int) []int {
	var table []int
	if max <= min {
		return table
	} else {
		for i := min; i < max; i++ {
			table = append(table, i)
		}
	}
	return table
}
