package topkfrequentelements

import (
	"container/heap"
	"fmt"
)

/*
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
*/

type item struct {
	value, times int
}

type hp []item

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].times < h[j].times }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(item))
}

func (h *hp) Pop() interface{} {
	l := len(*h)
	v := (*h)[l-1]
	(*h) = (*h)[:l-1]
	return v
}

// 优先队列了
// 统计频率类型的可以先哈希表＋一次遍历来获得所有次数元素
func topKFrequent(nums []int, k int) []int {
	ht := map[int]int{}
	for _, num := range nums {
		ht[num]++
	}
	h := &hp{}
	fmt.Println(ht)
	// 啥比才会把局部变量名写得跟函数参数名一样，还发现不了...
	// for k, v := range ht {
	for key, value := range ht {
		heap.Push(h, item{value: key, times: value})
		fmt.Println(h.Len(), k)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	fmt.Println(h)
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[k-i-1] = heap.Pop(h).(item).value
	}
	return ans
}

func TopKFrequent(nums []int, k int) []int {
	return topKFrequent(nums, k)
}

// 这题主要是对时间复杂度有要求，不然也可以不用堆来做，而且更加简洁且便于理解
