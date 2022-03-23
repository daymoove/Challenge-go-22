package piscine

func Index(s string, toFind string) int {
	l := []rune(toFind)
	i := 0
	nb := 0
	if len(toFind) != 0 {
		for index, letter := range s {
			if letter == l[0] {
				i = index
				nb++
				break
			} else if index == len(s)-1 && nb == 0 {
				return -1
			}
		}
		for i := 0; i < len(l); i++ {
			if l[i] == l[i] {
			} else {
				return -1
			}
		}
	}
	return i
}
