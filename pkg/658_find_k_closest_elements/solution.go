package findkclosestelements

import (
	"container/heap"
	"fmt"
	"sort"
)

type item struct {
	value int // 数组元素值
	dis   int // 元素与x的差值
}

type hp []item

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	return h[i].dis > h[j].dis || (h[i].dis == h[j].dis && h[i].value > h[j].value)
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(item))
}

func (h *hp) Pop() interface{} {
	l := len(*h)
	x := (*h)[l-1]
	(*h) = (*h)[:l-1]
	return x
}

// 选出k个离x最近的元素
func findClosestElements(arr []int, k int, x int) []int {
	h := &hp{}
	for _, value := range arr {
		heap.Push(h, item{value, abs(value, x)})
		fmt.Println("before", h)
		if h.Len() > k {
			heap.Pop(h)
		}
		fmt.Println("after", h)
	}
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = heap.Pop(h).(item).value
	}
	sort.Ints(ans)
	return ans
}

func abs(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func FindClosestElements(arr []int, k int, x int) []int {
	return findClosestElements(arr, k, x)
}

// 击败 10% + 87%
