package solutions

//固定长度的滑动窗口，加哈希表记录出现次数
func findRepeatedDnaSequences(s string) []string {
	res := []string{}
	l := 0
	ht := make(map[string]int)
	for l+10 <= len(s) {
		sub := s[l : l+10]
		ht[sub]++
		if ht[sub] == 2 {
			res = append(res, sub)
		}
		l++
	}
	return res
}

func FindRepeatedDnaSequences(s string) []string {
	return findRepeatedDnaSequences(s)
}
