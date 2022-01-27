package solutions

//给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
func missingNumber(nums []int) int {
	n := len(nums)
	vis := map[int]bool{}
	for _, num := range nums {
		vis[num] = true
	}
	for i := 0; i <= n; i++ {
		if !vis[i] {
			return i
		}
	}
	return -1
}

func MissingNumber(num []int) int {
	return missingNumber(num)
}
