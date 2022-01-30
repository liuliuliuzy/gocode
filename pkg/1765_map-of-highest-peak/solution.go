// package mapofhighestpeak

// import (
// 	"fmt"
// 	"math"
// )

// // 二维数组类型全局变量的初始化
// var dirs [4][2]int = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// // 从每个值为0的点开始做广度优先搜索
// func highestPeak(isWater [][]int) [][]int {
// 	m, n := len(isWater), len(isWater[0])
// 	ans := make([][]int, m)
// 	for i := range ans {
// 		ans[i] = make([]int, n)
// 		for j := range ans[i] {
// 			ans[i][j] = math.MaxInt
// 		}
// 	}
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if isWater[i][j] == 1 {
// 				fmt.Println("bfs", i, j)
// 				bfs(ans, m, n, i, j)
// 				return ans
// 			}
// 		}
// 	}
// 	return ans
// }

// func bfs(height [][]int, m, n, i, j int) {
// 	height[i][j] = 0
// 	fmt.Println("before bfs: ", height)
// 	hasSeen := make([][]bool, m)
// 	for i := range hasSeen {
// 		hasSeen[i] = make([]bool, n)
// 	}
// 	hasSeen[i][j] = true
// 	q := [][2]int{{i, j}}
// 	for len(q) > 0 {
// 		nowi, nowj := q[0][0], q[0][1]
// 		q = q[1:]
// 		// 当前节点的高度
// 		nowHeight := height[nowi][nowj]
// 		for _, dir := range dirs {
// 			tmpi, tmpj := nowi+dir[0], nowj+dir[1]
// 			if tmpi >= 0 && tmpi < m && tmpj >= 0 && tmpj < n && !hasSeen[tmpi][tmpj] {
// 				hasSeen[tmpi][tmpj] = true
// 				height[tmpi][tmpj] = min(height[tmpi][tmpj], nowHeight+1)
// 				q = append(q, [2]int{tmpi, tmpj})
// 			}
// 		}
// 	}
// 	fmt.Println("after bfs: ", height)
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// func HighestPeak(isWater [][]int) [][]int {
// 	return highestPeak(isWater)
// }

// --------------------- 第一遍没写出来，看一下题解思路，再写一遍 ---------------------------

package mapofhighestpeak

// 相邻的四个方向
var dirs [4][2]int = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// 记录下所有水域的位置，开始广度优先搜索
func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	hasSeen := make([][]bool, m)
	for i := range hasSeen {
		hasSeen[i] = make([]bool, n)
	}

	q := [][2]int{} // 队列
	level := 0      // 记录轮数
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				q = append(q, [2]int{i, j})
				hasSeen[i][j] = true
				count++
			}
		}
	}
	// 增加count变量，进行分轮次的广度优先搜索
	for len(q) > 0 {
		tmpCount := 0
		for count > 0 {
			// 从队列中取出索引
			nowi, nowj := q[0][0], q[0][1]
			q = q[1:]
			count--
			// 高度赋值
			ans[nowi][nowj] = level
			for _, dir := range dirs {
				tmpi, tmpj := nowi+dir[0], nowj+dir[1]
				// 遇到没有遍历过的相邻节点
				if tmpi >= 0 && tmpi < m && tmpj >= 0 && tmpj < n && !hasSeen[tmpi][tmpj] {
					hasSeen[tmpi][tmpj] = true
					q = append(q, [2]int{tmpi, tmpj})
					tmpCount++
				}
			}
		}
		// 每轮结束，更新状态变量
		count = tmpCount
		level++
	}
	return ans
}

func HighestPeak(isWater [][]int) [][]int {
	return highestPeak(isWater)
}
