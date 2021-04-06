package solutions

//感觉 似乎可以原地算法？
func merge(nums1 []int, m int, nums2 []int, n int) {
	startIndex := 0
	j := 0
	for {
		if startIndex > m-1 || j > n-1 {
			break
		}
		if nums1[startIndex] <= nums2[j] {
			startIndex += 1
		} else {
			nums1[startIndex] = nums2[j]
			startIndex += 1
		}
		j += 1
	}
	if j > n-1 {
		return
	} else {
		for j < n {
			nums1[startIndex] = nums2[j]
			startIndex += 1
			j += 1
		}
	}
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	merge(nums1, m, nums2, n)
}
