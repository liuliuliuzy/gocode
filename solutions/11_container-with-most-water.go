package solutions

import "goleetcode/mylibs"

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。
*/

// double pointer
func maxArea(height []int) int {
	// define two pointers
	l, r := 0, len(height)-1
	// maintain the max value
	ans := 0
	for l < r {
		w := r - l
		h := 0
		if height[l] < height[r] {
			h = height[l]
			// if l is shorter than r
			// than we should move l forward
			// because if we move r backward, we couldn't gain a greater value
			l++
		} else {
			h = height[r]
			r--
		}
		ans = mylibs.Max(w*h, ans)
	}
	return ans
}
func MaxArea(height []int) int {
	return maxArea(height)
}
