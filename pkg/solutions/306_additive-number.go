package solutions

import "strconv"

func isAdditiveNumber(num string) bool {
	// 判断分割方式是否符合要求
	var isCanAdded func(num string, first, second, sumIdx int) bool
	isCanAdded = func(num string, first, second, sumIdx int) bool {
		// 递归终止条件：索引走到了字符串末尾
		if sumIdx == len(num) {
			return true
		}
		// 计算应该出现的字符串
		sumStr := strconv.Itoa(first + second)
		if sumIdx+len(sumStr) > len(num) {
			return false
		}
		actualSum := num[sumIdx : sumIdx+len(sumStr)]
		// 递归调用，当前相等则继续判断后续内容
		return sumStr == actualSum && isCanAdded(num, second, first+second, sumIdx+len(sumStr))
	}
	// 遍历所有分割可能
	first := 0
	for i := 0; i < len(num); i++ {
		// 选取第一个数的切分处
		// 第一个数为0的话，i只能为0
		if i > 0 && num[0] == '0' {
			return false
		}

		// 累加计算切分方案下的第一个数
		first = first*10 + int(num[i]-'0')
		second := 0
		// 选取第二个数的切分处
		for j := i + 1; j < len(num); j++ {
			second = second*10 + int(num[j]-'0')
			// 同理，假如第二个数的开头为0的话，j只能为i+1
			if j > i+1 && num[i+1] == '0' {
				break
			}
			// 判断这种切分方案下是否满足要求
			if j+1 < len(num) && isCanAdded(num, first, second, j+1) {
				return true
			}
		}
	}
	return false
}

func IsAdditiveNumber(num string) bool {
	return isAdditiveNumber(num)
}
