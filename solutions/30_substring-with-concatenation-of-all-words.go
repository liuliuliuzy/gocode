package solutions

import "fmt"

// use hash table and sliding window
func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	if len(s) == 0 || len(words) == 0 {
		return res
	}
	words_ht := make(map[string]int)
	for _, item := range words {
		words_ht[item] += 1
	}
	// fmt.Print(words_ht)
	total_len := len(words) * len(words[0])
	for i := 0; i < len(s)-total_len; i++ {
		tmp_s := s[i : i+total_len]
		fmt.Println(tmp_s)
	}
	return res
}

func FindSubstring(s string, words []string) []int {
	return findSubstring(s, words)
}
