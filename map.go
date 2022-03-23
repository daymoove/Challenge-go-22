package piscine

func Map(f func(int) bool, a []int) []bool {
	boolderiz := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		if f(a[i]) == true {
			boolderiz[i] = true
		} else {
			boolderiz[i] = false
		}
	}
	return boolderiz
}
