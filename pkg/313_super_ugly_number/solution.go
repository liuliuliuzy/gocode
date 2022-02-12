package superuglynumber

import "math"

// 堆做法超时了...

// import (
// 	"container/heap"
// 	"sort"
// )

// // 堆 优先队列
// type hp struct{ sort.IntSlice } // 匿名类型字段，类型名就是字段名

// func (h *hp) Push(x interface{}) {
// 	// 类型断言只能用于接口类型，返回两个值，第一个值为类型，第二个值为断言是否正确
// 	h.IntSlice = append(h.IntSlice, x.(int))
// }

// func (h *hp) Pop() interface{} {
// 	l := len(h.IntSlice)
// 	v := h.IntSlice[l-1]
// 	h.IntSlice = h.IntSlice[:l-1]
// 	return v
// }

// func nthSuperUglyNumber(n int, primes []int) (ans int) {
// 	// 创建小顶堆
// 	h := &hp{sort.IntSlice{1}}
// 	seen := map[int]struct{}{1: {}} // 空成员变量的结构体与其值 {}
// 	for i := 1; ; i++ {
// 		item := heap.Pop(h).(int)
// 		if i == n {
// 			ans = item
// 			break
// 		}
// 		for _, p := range primes {
// 			next := item * p
// 			if _, has := seen[next]; !has {
// 				heap.Push(h, next)
// 				seen[next] = struct{}{}
// 			}
// 		}
// 	}
// 	return
// }

// func NthSuperUglyNumber(n int, primes []int) int {
// 	return nthSuperUglyNumber(n, primes)
// }

// 看题解得用动态规划

func nthSuperUglyNumber(n int, primes []int) int {
	dp := []int{0, 1}
	l := len(primes)
	// 创建索引数组并初始化
	pointers := make([]int, l)
	for i := range pointers {
		pointers[i] = 1
	}
	count := 1
	for count < n {
		tmp := math.MaxInt
		for i := range pointers {
			if dp[pointers[i]]*primes[i] < tmp {
				tmp = dp[pointers[i]] * primes[i]
			}
		}
		dp = append(dp, tmp)
		for i := range pointers {
			if dp[pointers[i]]*primes[i] == tmp {
				pointers[i]++
				// 这里不能想当然地break，因为存在相同计算结果的可能，所以相同就应该将索引加一
				// break
			}
		}
		count++
	}
	return dp[n]
}

func NthSuperUglyNumber(n int, primes []int) int {
	return nthSuperUglyNumber(n, primes)
}

/*
指针（索引）数组这种思路还是无法想到，觉得不是目前能够掌握的知识
*/
