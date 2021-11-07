package solutions

//不断地进行交集运算吧
func maxCount(m int, n int, ops [][]int) int {
	tmpM := m
	tmpN := n
	for _, op := range ops {
		tmpM = min(tmpM, op[0])
		tmpN = min(tmpN, op[1])
	}
	return tmpM * tmpN
}

func MaxCount(m int, n int, ops [][]int) int {
	return maxCount(m, n, ops)
}
