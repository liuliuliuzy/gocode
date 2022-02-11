package uglynumberii

import (
	"container/heap"
	"sort"
)

// 下面的做法无法保证大小顺序关系，比如4会被放在5的后面
// func nthUglyNumber(n int) int {
// 	ht := map[int]bool{1: true}
// 	q := []int{1}
// 	count := 0
// 	for len(q) < n {
// 		base := q[count]
// 		if !ht[2*base] {
// 			q = append(q, 2*base)
// 			ht[2*base] = true
// 		}
// 		if !ht[3*base] {
// 			q = append(q, 3*base)
// 			ht[3*base] = true
// 		}
// 		if !ht[5*base] {
// 			q = append(q, 5*base)
// 			ht[5*base] = true
// 		}
// 		count++
// 	}
// 	return q[n-1]
// }

var factors = []int{2, 3, 5}

// sort.IntSlice 是实现了sort.Interface的类型
type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	l := h.IntSlice.Len()
	a := h.IntSlice[l-1]
	h.IntSlice = h.IntSlice[:l-1]
	return a
}

func nthUglyNumber(n int) (ans int) {
	h := &hp{sort.IntSlice{1}}
	seen := make(map[int]bool)
	// 表示有或者没有，还可以用结构体作为value值类型
	// seen := make(map[int]struct{})
	for i := 1; ; i++ {
		item := heap.Pop(h).(int)
		if i == n {
			ans = item
			break
		}
		for _, f := range factors {
			next := item * f
			if _, ok := seen[next]; !ok {
				heap.Push(h, next)
				seen[next] = true
			}
		}
	}
	return
}

func NthUglyNumber(n int) int {
	return nthUglyNumber(n)
}
