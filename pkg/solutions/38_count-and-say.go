package solutions

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	res := ""
	s := countAndSay(n - 1)
	start := 0
	for start < len(s) {
		count := 1
		for start < len(s)-1 && s[start] == s[start+1] {
			start++
			count++
		}
		res += (strconv.Itoa(count) + string([]byte{s[start]}))
		start++
	}
	return res
}

func CountAndSay(n int) string {
	return countAndSay(n)
}
