package maximumnumberofballoons

import (
	"math"
)

// 统计ballon对应的五个字母的出现次数，然后取最小值（'l'的次数需要除以2）
func maxNumberOfBalloons(text string) int {
	counts := map[byte]int{}
	for i := 0; i < len(text); i++ {
		counts[text[i]]++
	}
	ans := math.MaxInt
	// 检查每个字母的出现次数
	for _, k := range []byte{'a', 'b', 'l', 'o', 'n'} {
		// 一个"balloon"中'l'和'o'都是出现两次
		if k == 'l' || k == 'o' {
			if ans > int(math.Floor(float64(counts[k])/2)) {
				ans = int(math.Floor(float64(counts[k]) / 2))
			}
		} else {
			if ans > counts[k] {
				ans = counts[k]
			}
		}
		// 小小优化
		if ans == 0 {
			break
		}
	}
	return ans
}

func MaxNumberOfBalloons(text string) int {
	return maxNumberOfBalloons(text)
}
