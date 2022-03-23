package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 0 && nb < 100 {
		for i := 1; i < nb+1; i++ {
			result = result * i
		}
	} else if nb == 0 {
		result = 1
	} else {
		result = 0
	}
	return result
}
