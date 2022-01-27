package solutions

import (
	"container/heap"
)

/*
给定两个以 升序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
*/

type pair struct{ i, j int }
type priorHeap struct {
	data         []pair
	nums1, nums2 []int
}

func (h priorHeap) Len() int {
	return len(h.data)
}

func (h priorHeap) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}

func (h priorHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *priorHeap) Push(v interface{}) {
	h.data = append(h.data, v.(pair))
}

func (h *priorHeap) Pop() interface{} {
	a := h.data
	v := a[len(a)-1]
	h.data = a[:len(a)-1]
	return v
}

// 优先队列
func kSmallestPairs(nums1 []int, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := priorHeap{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1})
		}
	}
	return
}

func KSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	return kSmallestPairs(nums1, nums2, k)
}
