package minimumcosttoreachdestinationintime

import (
	"math"
)

// ------------------------------------ dijkstra 尝试 -----------------------------
// import (
// 	"container/heap"
// 	"math"
// )

// 存在限制条件下的最短路径
// 解答错误...
// func minCost(maxTime int, edges [][]int, passingFees []int) (ans int) {
// 	n := len(passingFees)
// 	// 邻接表建图
// 	g := make([][]node, n)
// 	for _, e := range edges {
// 		s, d, c := e[0], e[1], e[2]
// 		g[s] = append(g[s], node{d, passingFees[d], c})
// 		g[d] = append(g[d], node{s, passingFees[s], c})
// 	}
// 	const inf int = math.MaxInt64 / 2
// 	minCosts := make([]int, n)
// 	for i := range minCosts {
// 		minCosts[i] = inf
// 	}
// 	minCosts[0] = passingFees[0]
// 	h := &hp{{0, passingFees[0], 0}}
// 	for h.Len() > 0 {
// 		now := heap.Pop(h).(node)
// 		nowi, nowc, nowt := now.index, now.cost, now.time
// 		// bfs剪枝条件
// 		if minCosts[nowi] < nowc && nowt > maxTime {
// 			continue
// 		}
// 		for _, nextNode := range g[nowi] {
// 			nexti, nextc, rt := nextNode.index, nextNode.cost, nextNode.time
// 			if newt, newc := nowt+rt, nowc+nextc; newt <= maxTime && newc < minCosts[nexti] {
// 				minCosts[nexti] = newc
// 				heap.Push(h, node{nexti, newc, newt})
// 			}
// 		}
// 	}
// 	ans = minCosts[n-1]
// 	if ans == inf {
// 		ans = -1
// 	}
// 	return
// }

// type node struct {
// 	index int // 目的地的索引
// 	cost  int // 到该点时的总收费
// 	time  int // 到该点时的总耗时
// }

// type hp []node

// func (h hp) Len() int              { return len(h) }
// func (h hp) Less(i, j int) bool    { return h[i].cost < h[j].cost }
// func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
// func (h *hp) Push(v interface{})   { *h = append(*h, v.(node)) }
// func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// ---------------------------------- 动态规划 ----------------------------------------
func minCost(maxTime int, edges [][]int, passingFees []int) int {
	const inf int = math.MaxInt64 / 2
	// 城市数量
	n := len(passingFees)
	dp := make([][]int, maxTime+1) // dp[t][n]: 恰好花费t分钟到达n情况下的最小费用
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][0] = passingFees[0]
	for t := 1; t <= maxTime; t++ {
		for _, edge := range edges {
			i, j, cost := edge[0], edge[1], edge[2]
			if cost <= t {
				dp[t][i] = min(dp[t][i], dp[t-cost][j]+passingFees[i])
				dp[t][j] = min(dp[t][j], dp[t-cost][i]+passingFees[j])
			}
		}
	}
	ans := inf
	for t := 1; t <= maxTime; t++ {
		ans = min(ans, dp[t][n-1])
	}
	if ans == inf {
		ans = -1
	}
	return ans
}

// 直接dp，平庸啊，太平庸了

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func MinCost(maxTime int, edges [][]int, passingFees []int) int {
	return minCost(maxTime, edges, passingFees)
}
