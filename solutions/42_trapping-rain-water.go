package solutions

import "fmt"

//很有意思的一道题
//寻找区间
//单调栈
func trap(height []int) (ans int) {
	stack := []int{}
	for i, h := range height {
		// fmt.Println(stack)
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			//遇到更大的元素，先将栈顶元素出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
			fmt.Println(left, i, curWidth, curHeight, stack)
		}
		stack = append(stack, i)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Trap(height []int) int {
	return trap(height)
}
