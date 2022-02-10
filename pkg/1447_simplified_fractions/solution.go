package simplifiedfractions

import "strconv"

// 数学方法直接枚举
func simplifiedFractions(n int) (ans []string) {
	// 暴力枚举分子和分母
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(j, i) == 1 {
				// go 允许字符串之间直接的加法
				ans = append(ans, strconv.Itoa(j)+"/"+strconv.Itoa(i))
			}
		}
	}
	return
}

// 求最大公约数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func SimplifiedFractions(n int) []string {
	return simplifiedFractions(n)
}
