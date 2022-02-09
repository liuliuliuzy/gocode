package sumofuniqueelements

func sumOfUnique(nums []int) (ans int) {
	// 哈希表
	seen := make(map[int]int)
	for _, num := range nums {
		if seen[num] == 0 {
			ans += num
			seen[num]++
		} else if seen[num] == 1 {
			ans -= num
			seen[num]++
		} else {
			continue
		}
	}
	return
}

func SumOfUnique(nums []int) int {
	return sumOfUnique(nums)
}
