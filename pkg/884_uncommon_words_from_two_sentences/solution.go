package uncommonwordsfromtwosentences

import "strings"

// 集合问题
// 哈希表
func uncommonFromSentences(s1 string, s2 string) []string {
	w1 := strings.Fields(s1)
	w2 := strings.Fields(s2)
	ans := []string{}
	// 两次遍历
	m := make(map[string]int)
	for _, w := range w1 {
		m[w]++
	}
	for _, w := range w2 {
		m[w]++
	}
	// 遍历哈希表
	for k, v := range m {
		if v == 1 {
			ans = append(ans, k)
		}
	}
	return ans
}

func UncommonFromSentences(s1 string, s2 string) []string {
	return uncommonFromSentences(s1, s2)
}
