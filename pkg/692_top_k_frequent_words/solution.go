package topkfrequentwords

import (
	"container/heap"
	"fmt"
)

// 继续用堆来解决这一问题
type item struct {
	value string
	times int
}

type hp []item

// 堆顶的是最不需要的元素
// func (h hp) Less(i, j int) bool {
// 	// 字符串出现次数更少
// 	if h[i].times < h[j].times {
// 		return true
// 	} else {
// 		if h[i].times == h[j].times {
// 			l1, l2 := len(h[i].value), len(h[j].value)
// 			l := min(l1, l2)
// 			// 字母顺序在后
// 			for k := 0; k < l; k++ {
// 				if h[i].value[k] > h[j].value[k] {
// 					return true
// 				} else if h[i].value[k] < h[j].value[k] {
// 					return false
// 				} else {
// 					continue
// 				}
// 			}
// 			// 字符串长度更长
// 			if l1 > l2 {
// 				return true
// 			} else {
// 				return false
// 			}
// 		} else {
// 			return false
// 		}
// 	}
// }

// 原来golang string是支持对比运算符的...
func (h hp) Less(i, j int) bool {
	return h[i].times < h[j].times || (h[i].times == h[j].times && h[i].value > h[j].value)
}

func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(item))
}

func (h *hp) Pop() interface{} {
	l := len(*h)
	x := (*h)[l-1]
	*h = (*h)[:l-1]
	return x
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

func topKFrequent(words []string, k int) []string {
	ht := make(map[string]int)
	for _, s := range words {
		ht[s]++
	}
	h := &hp{}
	for s, times := range ht {
		heap.Push(h, item{s, times})
		fmt.Println("b: ", h)
		if h.Len() > k {
			heap.Pop(h)
		}
		fmt.Println("a: ", h)
	}
	ans := make([]string, k)
	for i := 0; i < k; i++ {
		ans[k-i-1] = heap.Pop(h).(item).value
	}
	return ans
}

func TopKFrequent(words []string, k int) []string {
	return topKFrequent(words, k)
}
