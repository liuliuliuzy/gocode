package pathwithmaximumprobability

import (
	"container/heap"
)

/*
给你一个由 n 个节点（下标从 0 开始）组成的无向加权图，该图由一个描述边的列表组成，其中 edges[i] = [a, b] 表示连接节点 a 和 b 的一条无向边，且该边遍历成功的概率为 succProb[i] 。

指定两个节点分别作为起点 start 和终点 end ，请你找出从起点到终点成功概率最大的路径，并返回其成功概率。

如果不存在从 start 到 end 的路径，请 返回 0 。只要答案与标准答案的误差不超过 1e-5 ，就会被视作正确答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-with-maximum-probability
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 感觉是翻版的dijkstra
// 内存炸了，什么鬼测试用例...
// func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) (ans float64) {
// 	g := make([][]float64, n)
// 	for i := range g {
// 		g[i] = make([]float64, n)
// 	}
// 	for i, edge := range edges {
// 		s, d, p := edge[0], edge[1], succProb[i]
// 		// 无向加权图
// 		g[s][d] = p
// 		g[d][s] = p
// 	}
// 	probs := make([]float64, n)
// 	// 将起点的概率设为1
// 	probs[start] = 1
// 	used := make([]bool, n)
// 	for {
// 		// 枚举法寻找当前最大概率到达的点
// 		v := -1
// 		for w, u := range used {
// 			if !u && (v < 0 || probs[w] > probs[v]) {
// 				v = w
// 			}
// 		}
// 		// 如果所有节点都已经确定了最大概率
// 		if v == -1 {
// 			break
// 		}
// 		used[v] = true
// 		for n, p := range g[v] {
// 			// 未计算过概率的新点
// 			// if probs[n] == 0 {
// 			// 	probs[n] = probs[v] * p
// 			// } else {
// 			if newP := probs[v] * p; newP > probs[n] {
// 				probs[n] = newP
// 			}
// 			// }
// 		}
// 	}
// 	fmt.Println(probs)
// 	ans = probs[end]
// 	return
// }

// 改用大根堆再写一遍
func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) (ans float64) {
	g := make([][]pair, n)
	for i, edge := range edges {
		s, d := edge[0], edge[1]
		p := succProb[i]
		g[s] = append(g[s], pair{d, p})
		g[d] = append(g[d], pair{s, p})
	}
	probs := make([]float64, n)
	// 初始设置，将起点的概率置为1。这是至关重要的，因为后续所有的结果都来自于这个1
	probs[start] = 1
	// 堆中也是在初始状态下只放置了起点
	h := &hp{{start, 1}}
	for len(*h) > 0 {
		node := heap.Pop(h).(pair)
		// 剪枝操作：
		// 节点第一次被当作已确定节点时，其值必然是最优的
		if node.index == end {
			ans = node.probs
			return
		}
		// 回想一下，其实这里的“过滤无用点“操作是很有用的
		if node.probs < probs[node.index] {
			continue
		}
		for _, r := range g[node.index] {
			if newP := probs[node.index] * r.probs; newP > probs[r.index] {
				probs[r.index] = newP
				heap.Push(h, pair{r.index, newP})
			}
		}
	}
	ans = probs[end]
	return
}

type pair struct {
	index int
	probs float64
}

type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].probs > h[j].probs }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(pair)) }

// 更为优雅的Pop实现，直接一行语句进行两次赋值，然后返回
func (h *hp) Pop() (v interface{}) { v, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]; return }

func MaxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	return maxProbability(n, edges, succProb, start, end)
}
