package solutions

// 广度优先
// func numWays(n int, relation [][]int, k int) int {
// 	// 建图
// 	edges := map[int][]int{}
// 	for _, rela := range relation {
// 		edges[rela[0]] = append(edges[rela[0]], rela[1])
// 	}
// 	que := []int{0}
// 	steps := 0
// 	for len(que) > 0 && steps < k {
// 		steps++
// 		l := len(que)
// 		for l > 0 {
// 			start := que[0]
// 			que = que[1:]
// 			que = append(que, edges[start]...) //还有这种操作
// 			l--
// 		}
// 	}
// 	res := 0
// 	if steps == k {
// 		for _, des := range que {
// 			if des == n-1 {
// 				res++
// 			}
// 		}
// 	}

// 	return res
// }

// 深度优先
// func numWays(n int, relation [][]int, k int) int {
// 	// 建图
// 	edges := map[int][]int{}
// 	for _, rela := range relation {
// 		edges[rela[0]] = append(edges[rela[0]], rela[1])
// 	}
// 	res := 0
// 	// 注意，如果函数闭包存在自身的递归调用，那么就不能将其 declaration 和 assignment 合并在一起
// 	var dfs func(int, int)
// 	dfs = func(start, steps int) {
// 		if steps == k {
// 			if start == n-1 {
// 				res++
// 			}
// 			return //终止
// 		}
// 		for _, to := range edges[start] {
// 			dfs(to, steps+1)
// 		}
// 	}
// 	dfs(0, 0)
// 	return res
// }

//再记个官方题解的dp解法
func numWays(n int, relation [][]int, k int) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < k; i++ {
		next := make([]int, n)
		for _, r := range relation {
			src, dst := r[0], r[1]
			next[dst] += dp[src]
		}
		dp = next
	}
	return dp[n-1]
}

func NumWays(n int, relation [][]int, k int) int {
	return numWays(n, relation, k)
}
