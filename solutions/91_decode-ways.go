package solutions

// 动态规划
func numDecodings(s string) int {
	// 使用数组进行dp
	// dp := make([]int, n+1)
	// dp[0] = 1
	// for i := 0; i < n; i++ {
	// 	fmt.Println(s[i])
	// 	if s[i] != '0' {
	// 		dp[i+1] += dp[i]
	// 	}
	// 	if i > 0 && s[i-1] != '0' && ((s[i-1]-'0')*10+(s[i]-'0')) <= 26 {
	// 		dp[i+1] += dp[i-1]
	// 	}
	// }
	// return dp[n]

	// 进一步优化，使用三个变量进行动态规划
	n := len(s)
	// a: f_n-2
	// b: f_n-1
	// c: f_n
	a, b, c := 0, 1, 0
	for i := 0; i < n; i++ {
		c = 0
		if s[i] != '0' {
			c += b
		}
		if i > 0 && s[i-1] != '0' && ((s[i-1]-'0')*10+s[i]-'0') <= 26 {
			c += a
		}
		a, b = b, c
	}
	return c
}

func NumDecodings(s string) int {
	return numDecodings(s)
}
