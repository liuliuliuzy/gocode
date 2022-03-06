package cheapestflightswithinkstops

// 动态规划
// 状态设置：我们用 f[t][i] 表示通过恰好 t 次航班，从出发城市 src 到达城市 i 需要的最小花费。
func findCheapestPrice(n int, flights [][]int, src, dst, k int) int {
	const inf int = 10000*101 + 1
	// dp二维数组
	f := make([][]int, k+2) // 为什么长度是k+2呢？懂了，因为k个中转站，就意味着最多搭乘k+1次航班，所以索引需要到k+1
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[0][src] = 0
	for t := 1; t < k+2; t++ {
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			f[t][i] = min(f[t-1][j]+cost, f[t][i])
		}
	}
	ans := inf
	for t := 1; t < k+2; t++ {
		ans = min(ans, f[t][dst])
	}
	if ans == inf {
		ans = -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 看起来 动态规划 比 bfs+优先队列 要简单易懂许多啊
