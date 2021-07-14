package solutions

import (
	"fmt"
	"sort"
)

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	// fmt.Println("entered")
	testNums := []int{2, 5}
	fmt.Println(search(testNums, 1)) //expected: 1, not 0
	res := []int{}
	res = append(res, nums1...)
	sort.Ints(res)
	sum, maxn, n := 0, 0, len(nums1)
	fmt.Println(res)
	for i := 0; i < n; i++ {
		tmp := abs(nums1[i] - nums2[i])
		sum += tmp
		j := search(res, nums2[i])
		fmt.Println(j, nums2[i])
		maxn = max(maxn, tmp-abs(nums2[i]-res[j]))
		if j < n-1 {
			maxn = max(maxn, tmp-abs(res[j]-nums2[i]))
		}
	}
	return (sum - maxn) % (1e9 + 7)
}

func MinAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	return minAbsoluteSumDiff(nums1, nums2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

//nums要求是升序排序后的数组
func search(nums []int, item int) int {
	s, e := 0, len(nums)-1
	if nums[e] < item {
		return e
	}
	for s < e {
		m := (s + e) / 2
		if nums[m] < item {
			s = m + 1
		} else {
			e = m
		}
	}
	return s
}
