package solutions

//位运算 O(n) + O(1)
func singleNumber_i(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}
func SingleNumber_i(nums []int) int {
	return singleNumber_i(nums)
}
