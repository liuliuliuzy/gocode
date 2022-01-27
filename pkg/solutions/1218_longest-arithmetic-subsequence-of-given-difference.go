package solutions

//tags: array, hash, dynamic programming
func longestSubsequence(arr []int, difference int) (ans int) {
	//I think it's a tricky use of hash table in go
	dp := map[int]int{} //create a map, same a **dp := make(map[int]int)**
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return
}
func LongestSubsequence(arr []int, difference int) int {
	return longestSubsequence(arr, difference)
}
