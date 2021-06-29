package solutions

import "fmt"

//很有意思的一道题
//寻找区间
func trap(height []int) int {
	type pair struct {
		left  int
		right int
	}
	subsets := []pair{}
	item := pair{-1, -1}
	upFlag := false
	for i := 1; i < len(height); i++ {
		if height[i] < height[i-1] {
			//之前处于上升阶段
			if upFlag {
				item.right = i - 1
			} else {
				item.left = i - 1
			}
			upFlag = false
			if item.right > 0 {
				subsets = append(subsets, item)
				item.left = 
			}
		} else if height[i] > height[i-1] {
			upFlag = true
		} else {
			continue
		}
	}
	fmt.Println(subsets)
	return 0
}

func Trap(height []int) int {
	return trap(height)
}
