package solutions

func MaxPower(s string) int {
	return maxPower(s)
}

func maxPower(s string) int {
	ans := 0
	count := 1
	startC := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != startC {
			ans = max(count, ans) // use the bigger value
			count = 1             // set count to 1
			startC = s[i]
		} else {
			count += 1
		}
	}
	return max(ans, count)
}
