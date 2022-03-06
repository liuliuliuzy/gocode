package networkdelaytime

// import (
// 	"math"
// )

// // 相当于求有向加权图的单源最短路径
// // 枚举类型的dijkstra算法
// func networkDelayTime(times [][]int, n, k int) int {
// 	// 第一步肯定是建图，这里选择邻接矩阵
// 	edges := make([][]int, n+1)
// 	for i := 0; i < n+1; i++ {
// 		edges[i] = make([]int, n+1)
// 		for j := 0; j < n+1; j++ {
// 			edges[i][j] = math.MaxInt64 / 2
// 		}
// 	}
// 	for _, item := range times {
// 		edges[item[0]][item[1]] = item[2]
// 	}
// 	// 存放k节点到各个节点的最短路径长度，算法执行完成后，其会存储每个节点到出发节点的最短距离
// 	shortestPath := make([]int, n+1)
// 	for i := 1; i < n+1; i++ {
// 		shortestPath[i] = math.MaxInt64 / 2
// 	}
// 	// 下面开始使用dijkstra算法
// 	shortestPath[k] = 0 // k到自身的距离为0

// 	// 记录每个节点的选用情况
// 	used := make([]bool, n+1)
// 	// n 个节点，循环n次
// 	for i := 0; i < n; i++ {
// 		x := -1
// 		for y, u := range used[1:] {
// 			// 选出距起点最近的节点
// 			if !u && (x == -1 || shortestPath[y+1] < shortestPath[x]) { // 注意这里索引的处理
// 				x = y + 1
// 			}
// 		}
// 		// 标记为已使用，将其作为已确定节点
// 		used[x] = true
// 		// 用 已确定节点 来更新其他节点（已确定节点的邻居节点）
// 		for y, time := range edges[x][1:] {
// 			shortestPath[y+1] = min(shortestPath[x]+time, shortestPath[y+1])
// 		}
// 	}
// 	ans := -1
// 	for _, d := range shortestPath[1:] {
// 		if d == math.MaxInt64/2 {
// 			return -1
// 		}
// 		ans = max(ans, d)
// 	}
// 	return ans
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

func NetworkDelayTime(times [][]int, n, k int) int {
	return networkDelayTime(times, n, k)
}
