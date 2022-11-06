package longestuncommonsubsequencei

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

/*
寻找两个字符串中【最长独特子序列】的长度
【最长独特子序列】定义：该序列为某字符串独有的最长子序列（即不能是其他字符串的子序列）。
*/
func findLUSlength(a, b string) int {
	if a != b {
		return max(len(a), len(b))
	}
	return -1
}

func FindLUSlength(a, b string) int {
	return findLUSlength(a, b)
}
