package solutions

import "fmt"

func maxDepth(s string) int {
	// 等价于计算栈的最大高度
	count, ans := 0, 0
	for _, c := range s {
		if c == '(' {
			count++
		} else if c == ')' {
			ans = max(ans, count)
			count--
		} else {
			continue
		}
		fmt.Println(c)
	}
	return ans
}

func MaxDepth(s string) int {
	return maxDepth(s)
}
