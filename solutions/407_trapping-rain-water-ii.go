package solutions

import "container/heap"

type cell struct {
	h, x, y int
}

type hp []cell

//实现heap的接口
func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].h < h[j].h
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(cell))
}

func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

//#42接雨水问题 扩展到二维场景
//太难了啊喂
func trapRainWater(heightMap [][]int) int {
	ans := 0
	m, n := len(heightMap), len(heightMap[0])
	if m <= 2 || n <= 2 {
		return ans
	}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	h := &hp{}
	for i, row := range heightMap {
		for j, v := range row {
			//遍历最外层
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				heap.Push(h, cell{v, i, j})
				vis[i][j] = true //mark
			}

		}
	}

	dirs := []int{-1, 0, 1, 0, -1} // 用长度为5的一维数组来表示上下左右4个方向的移动。你学废了吗
	for h.Len() > 0 {
		cur := heap.Pop(h).(cell) // . means type assertion?
		for k := 0; k < 4; k++ {
			nx, ny := cur.x+dirs[k], cur.y+dirs[k+1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && !vis[nx][ny] {
				if heightMap[nx][ny] < cur.h {
					ans += cur.h - heightMap[nx][ny]
				}
				vis[nx][ny] = true
				heap.Push(h, cell{max(heightMap[nx][ny], cur.h), nx, ny})
			}
		}
	}
	return ans
}

func TrapRainWaterII(heightMap [][]int) int {
	return trapRainWater(heightMap)
}
