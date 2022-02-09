package countnumberofpairswithabsolutedifferencek

import "fmt"

// 哈希表
func countKDifference(nums []int, k int) (ans int) {
	ht := make(map[int]int)
	// 对于每一个元素，其索引都不同，所以不会造成重复
	for _, item := range nums {
		ans += ht[item+k] + ht[item-k]
		ht[item]++
		fmt.Println(item, ht, ans)
	}
	return
}

func CountKDifference(nums []int, k int) int {
	return countKDifference(nums, k)
}
