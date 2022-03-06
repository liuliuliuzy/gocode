package zigzagconversion

func convert(s string, numRows int) string {
	// 直接数学规律
	n, r := len(s), numRows
	// 特殊情况处理
	if r == 1 || n <= r {
		return s
	}
	t := r*2 - 2
	ans := make([]byte, 0, n) // length -> 0, cap -> n
	for i := 0; i < r; i++ {  // 枚举行
		for j := 0; j+1 < n; j += t { // 去寻找可能存在的行中的每一个元素
			ans = append(ans, s[j+i]) // 当前周期的第一个字符
			if 0 < i && i < r-1 && j+t-i < n {
				ans = append(ans, s[j+t-i]) // 当前周期的第二个字符
			}
		}
	}
	return string(ans)
}

func Convert(s string, numRows int) string {
	return convert(s, numRows)
}
