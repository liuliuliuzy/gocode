package pathwithminimumeffort

import (
	"container/heap"
	"fmt"
	"math"
)

var dirs [4][2]int = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// 变种dijkstra
// 一遍过，开森！
func minimumEffortPath(heights [][]int) (ans int) {
	const inf int = math.MaxInt64 / 2
	m, n := len(heights), len(heights[0])
	costs := make([][]int, m)
	for i := range costs {
		costs[i] = make([]int, n)
		for j := range costs[i] {
			costs[i][j] = inf
		}
	}
	costs[0][0] = 0
	h := &hp{{0, 0, 0}}
	for h.Len() > 0 {
		now := heap.Pop(h).(element)
		nowi, nowj, nowcost := now.i, now.j, now.cost
		if nowcost > costs[nowi][nowj] {
			continue
		}
		for _, dir := range dirs {
			nexti, nextj := nowi+dir[0], nowj+dir[1]
			if nexti >= 0 && nexti < m && nextj >= 0 && nextj < n {
				newC := max(abs(heights[nowi][nowj], heights[nexti][nextj]), costs[nowi][nowj])
				if newC < costs[nexti][nextj] {
					costs[nexti][nextj] = newC
					heap.Push(h, element{nexti, nextj, newC})
				}
			}
		}
	}
	fmt.Println(costs)
	ans = costs[m-1][n-1]
	return
}

type element struct {
	i, j int
	cost int
}
type hp []element

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].cost < h[j].cost }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(element)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func abs(i, j int) int {
	if i > j {
		return i - j
	} else {
		return j - i
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinimumEffortPath(heights [][]int) int {
	return minimumEffortPath(heights)
}
