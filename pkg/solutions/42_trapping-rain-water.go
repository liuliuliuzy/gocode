package solutions

//很有意思的一道题
//寻找区间

//解法一：动态规划，计算坐标i的左右两边的最大高度
// func trap(height []int) (ans int) {
// 	n := len(height)
// 	if n == 0 {
// 		return
// 	}
// 	leftMax := make([]int, n)
// 	rightMax := make([]int, n)
// 	// 设置初始值
// 	leftMax[0] = height[0]
// 	rightMax[n-1] = height[n-1]
// 	for i := 1; i < n; i++ {
// 		leftMax[i] = max(leftMax[i-1], height[i])
// 	}
// 	for i := n - 2; i > -1; i-- {
// 		rightMax[i] = max(height[i], rightMax[i+1])
// 	}
// 	//计算每个坐标i处所能接的雨水量
// 	for i, h := range height {
// 		ans += min(leftMax[i], rightMax[i]) - h
// 	}
// 	//提前声明好的返回值，在这里也是需要执行return来返回的
// 	return
// }

//解法二：单调栈
// func trap(height []int) (ans int) {
// 	stack := []int{}
// 	for i, h := range height {
// 		// fmt.Println(stack)
// 		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
// 			//遇到更大的元素，先将栈顶元素出栈
// 			top := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			if len(stack) == 0 {
// 				break
// 			}
// 			left := stack[len(stack)-1]
// 			curWidth := i - left - 1
// 			curHeight := min(height[left], h) - height[top]
// 			ans += curWidth * curHeight
// 			fmt.Println(left, i, curWidth, curHeight, stack)
// 		}
// 		stack = append(stack, i)
// 	}
// 	return
// }

//解法三：双指针，可以看作是解法一的空间优化，用指针与变量来代替数组
//非常巧妙的一种解法
func trap(height []int) int {
	ans := 0
	n := len(height)
	if n == 0 {
		return ans
	}
	left := 0
	right := n - 1
	leftMax := 0
	rightMax := 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left += 1
		} else {
			ans += rightMax - height[right]
			right -= 1
		}
	}
	return ans
}

func Trap(height []int) int {
	return trap(height)
}
