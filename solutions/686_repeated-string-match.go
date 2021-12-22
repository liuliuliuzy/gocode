package solutions

import (
	"math"
	"strings"
)

// 略显啰嗦
func repeatedStringMatch(a string, b string) int {
	var ans int
	m, n := len(a), len(b)
	q := int(math.Floor(float64(n-1) / float64(m)))
	if q == 0 {
		if strings.Contains(a, b) {
			ans = 1
		} else {
			if strings.Contains(a+a, b) {
				ans = 2
			} else {
				ans = -1
			}
		}
	} else if q == 1 {
		if strings.Contains(a+a, b) {
			ans = 2
		} else {
			if strings.Contains(a+a+a, b) {
				ans = 3
			} else {
				ans = -1
			}
		}
	} else {
		content := strings.Repeat(a, q)
		if !strings.Contains(b, content) {
			ans = -1
		} else {
			ans = q
			start := strings.Index(b, a)
			head := b[0:start]
			tail := b[start+q*m : n]
			lh, lt := len(head), len(tail)
			// 得到了2个新的字符串
			if lh > 0 {
				if a[m-lh:m] != head {
					return -1
				}
				ans += 1
			}
			if lt > 0 {
				if a[0:lt] != tail {
					return -1
				}
				ans += 1
			}
		}
	}
	return ans
}

// 比我的方法差
// func repeatedStringMatch(a string, b string) int {
// 	ans := 1
// 	tmpA := a
// 	maxLen := len(a)*2 + len(b)
// 	l := len(a)
// 	for l < maxLen {
// 		if strings.Contains(tmpA, b) {
// 			return ans
// 		} else {
// 			ans++
// 			tmpA += a
// 			l += len(a)
// 		}
// 	}
// 	return -1
// }

func RepeatedStringMatch(a string, b string) int {
	return repeatedStringMatch(a, b)
}
