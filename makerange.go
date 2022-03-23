package piscine

func MakeRange(min, max int) []int {
	size := 0
	if min >= max {
		table := make([]int, size)
		table = nil
		return table
	} else {
		size := max - min
		table := make([]int, size)
		for i := 0; i < size; i++ {
			table[i] = i + min
		}
		return table
	}
}
