package numberofrectanglesthatcanformthelargestsquare

// 哈希表
func countGoodRectangles(rectangles [][]int) int {
	nowLen := 0
	ht := make(map[int]int)
	for _, rec := range rectangles {
		t := min(rec[0], rec[1])
		if t > nowLen {
			nowLen = t
		}
		ht[t]++
	}
	return ht[nowLen]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func CountGoodRectangles(rectangles [][]int) int {
	return countGoodRectangles(rectangles)
}
