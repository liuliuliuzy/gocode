package findmedianfromdatastream

import (
	"container/heap"
	"sort"
)

/*
堆 优先队列
经典数据结构题，用一个大顶堆＋一个小顶堆来维护中位数
*/

type hp struct{ sort.IntSlice }

func (h *hp) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp) Pop() interface{} {
	l := len(h.IntSlice)
	x := h.IntSlice[l-1]
	h.IntSlice = h.IntSlice[:l-1]
	return x
}

type MedianFinder struct {
	qMin, qMax hp // qMin：大顶堆，记录小于等于中位数的所有数; qMax: 小顶堆，记录大于中位数的所有数
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

// add number & maintain the heap
func (m *MedianFinder) AddNum(num int) {
	minQ, maxQ := &m.qMin, &m.qMax
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num) // 通过添加相反数的方式来将小顶堆变为大顶堆，小小技巧
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (m *MedianFinder) FindMedian() float64 {
	minQ, maxQ := m.qMin, m.qMax
	// minQ has one more element than maxQ, so the top of minQ is the median number
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	// otherwise, the median number will be obtained by adding the two top elements in the heap and divide the sum by 2
	return float64(-minQ.IntSlice[0]+maxQ.IntSlice[0]) / 2
}
