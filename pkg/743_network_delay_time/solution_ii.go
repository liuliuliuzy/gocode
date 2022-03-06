package networkdelaytime

import (
	"container/heap"
	"fmt"
	"math"
)

// 小顶堆类型的dijkstra算法，go语言版本需要自己实现heap要求的接口
func networkDelayTime(times [][]int, n, k int) (ans int) {
	type edge struct{ to, time int }
	// 这里选择邻接表来表示图
	g := make([][]edge, n)
	// 创建邻接表
	for _, t := range times {
		x, y := t[0]-1, t[1]-1
		g[x] = append(g[x], edge{y, t[2]})
	}
	// 其它点的初始距离设置为无穷
	const inf int = math.MaxInt64 / 2
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	// 出发点的距离设为0
	dist[k-1] = 0

	// ========== 下面开始不断地从未确定节点集合中选取离起点最近的点，然后更新其它未确定的点的距离 ==========

	// 创建小根堆，获得距离最近的节点的优先队列，其中存储的就是未确定最短距离的节点
	h := &hp{{0, k - 1}} // 一开始最近的节点就是起点，索引为k-1，距离为0
	// 只要还存在未确定最短距离的节点
	for h.Len() > 0 {
		fmt.Println("now: ", h)
		// 取出当前最短距离的节点
		p := heap.Pop(h).(pair)
		x := p.x
		// 过滤无用的节点
		if dist[x] < p.d {
			continue
		}
		// 用取出的点去更新其邻居节点的距离
		for _, e := range g[x] {
			y := e.to
			if d := dist[x] + e.time; d < dist[y] {
				// 更新距离
				dist[y] = d
				// 加入优先队列
				heap.Push(h, pair{d, y}) // 那队列中岂不是会存在重复的节点对，只不过距离不同。是的，调试的时候发现确实会这样
			}
		}
	}

	for _, d := range dist {
		if d == inf {
			return -1
		}
		ans = max(ans, d)
	}
	return
}

type pair struct {
	d int // 与起点的距离
	x int // 节点索引编号
}
type hp []pair

// 实现 golang heap 要求的接口

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
