package solutions

//其实就是合并区间问题
func isCovered(ranges [][]int, left int, right int) bool {
	sections := make(map[int]bool)
	for _, section := range ranges {
		for i := section[0]; i < section[1]+1; i++ {
			sections[i] = true
		}
	}
	for i := left; i < right+1; i++ {
		if sections[i] == false {
			return false
		}
	}
	return true
}

func IsCovered(ranges [][]int, left int, right int) bool {
	return isCovered(ranges, left, right)
}
