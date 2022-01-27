package solutions

// 方法一：回溯
// 力扣提交超时...
// func grayCode(n int) []int {
// 	ans := []int{0} // 创建slice存储遍历过的数
// 	total := 1 << n
// 	var bc func(int, int) bool
// 	bc = func(num, count int) bool {
// 		// 终止条件
// 		if count == total {
// 			return true
// 		}
// 		// 否则
// 		flag := false
// 		// 选取一个不同的bit位
// 		for i := 0; i < n; i++ {
// 			// 计算对应的数
// 			tmp := 0
// 			if (num>>i)&1 > 0 {
// 				tmp = num - 1<<i
// 			} else {
// 				tmp = num + 1<<i
// 			}
// 			// fmt.Println(num, tmp)
// 			// 如果用过
// 			if contains(ans, tmp) {
// 				continue
// 			}
// 			// 尝试将其加入到ans中
// 			ans = append(ans, tmp)
// 			if bc(tmp, count+1) {
// 				flag = true
// 				break
// 			}
// 			// 尝试失败，将其移除
// 			ans = ans[:len(ans)-1]
// 		}
// 		return flag
// 	}
// 	bc(0, 1)
// 	return ans
// }

// func contains(s []int, value int) bool {
// 	for _, v := range s {
// 		if v == value {
// 			return true
// 		}
// 	}
// 	return false
// }

/*=============================================*/

// 方法二：对称生成，就是知道方法之后模拟
// 生成方法：https://leetcode-cn.com/problems/gray-code/solution/ge-lei-bian-ma-by-leetcode-solution-cqi7/
// func grayCode(n int) []int {
// 	ans := make([]int, 1, 1<<n) // 创建slice，length: 1, capacity: 1<<n
// 	for i := 1; i <= n; i++ {
// 		for j := len(ans) - 1; j >= 0; j-- {
// 			ans = append(ans, ans[j]|1<<(i-1))
// 		}
// 	}
// 	return ans
// }

/*=============================================*/

// 方法三：二进制数直接转格雷码（利用格雷码的性质）
// func grayCode(n int) []int {
// 	var l uint = 1 << uint(n)
// 	out := make([]int, l)
// 	for i := uint(0); i < l; i++ {
// 		out[i] = int((i >> 1) ^ i)
// 	}
// 	return out
// }

// 方法四：对称生成之递归
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	half := grayCode(n - 1)
	for i := (1<<(n-1) - 1); i >= 0; i-- {
		half = append(half, (1<<(n-1))+half[i])
	}
	return half
}

func GrayCode(n int) []int {
	return grayCode(n)
}
