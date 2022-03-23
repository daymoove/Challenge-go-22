package piscine

func Join(strs []string, sep string) string {
	var m string
	for i := 0; i < len(strs); i++ {
		m = m + strs[i]
		if i == len(strs)-1 {
		} else {
			m = m + sep
		}
	}
	return m
}
