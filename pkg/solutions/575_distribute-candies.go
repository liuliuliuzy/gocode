package solutions

//计算一共有多少种就行了吧，然后和半数对比一下
func distributeCandies(candyType []int) int {
	types := make(map[int]bool)
	typeCount := 0
	for _, candy := range candyType {
		if !types[candy] {
			types[candy] = true
			typeCount += 1
		}
	}
	if typeCount >= len(candyType)/2 {
		return len(candyType) / 2
	} else {
		return typeCount
	}
}

func DistributeCandies(candyType []int) int {
	return distributeCandies(candyType)
}
