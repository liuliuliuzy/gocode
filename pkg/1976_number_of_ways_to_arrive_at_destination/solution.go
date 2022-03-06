package numberofwaystoarriveatdestination

/*
寻找最短路径的条数
*/

// 先想一下动态规划咋做

func countPaths(n int, roads [][]int) (ans int) {
	// 稠密图用邻接矩阵表示更合适
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = 1e18
		}
	}

	for _, road := range roads {
		src, dst, costTime := road[0], road[1], road[2]
		// 题目给的边是无向图的边，将其当作有向图来看的话，就表示两条相反方向的边
		g[src][dst] = costTime
		g[dst][src] = costTime
	}

	// 我们先求最短路径
	d := make([]int, n)
	d[0] = 0
	for i := 1; i < n; i++ {
		d[i] = 1e18
	}
	// 下面开始基于枚举的dijkstra
	used := make([]bool, n)
	for {
		useNode := -1
		for node, isUsed := range used {
			if !isUsed && (useNode < 0 || d[node] < d[useNode]) {
				useNode = node
			}
		}

		// 如果找不到离起点最近的节点，那就代表dijkstra算法完成，起点到所有节点的最短路径已经计算完毕
		if useNode < 0 {
			break
		}

		// 将最近的节点视作已经确定其最短距离，然后将其标志位置位
		used[useNode] = true
		// 用选取的节点来更新其所有邻居节点的最短路径长度
		for nextNode, costTime := range g[useNode] {
			// 计算距离，如果比记录的还小，则更新记录的节点距离
			if newD := d[useNode] + costTime; newD < d[nextNode] {
				d[nextNode] = newD
			}
		}
	}

	// 统计最短路径构成的有向图中各个节点的入度
	deg := make([]int, n)
	for node, neighbors := range g {
		for nextNode, costTime := range neighbors {
			if d[node]+costTime == d[nextNode] {
				// 说明边 node -> nextNode 在最短路径上
				deg[nextNode]++ // nextNode节点的入度++
			}
		}
	}

	// 动态规划统计0到n的最短路径数量
	dp := make([]int, n)
	// 0 到 0 的最短路径有1条
	dp[0] = 1
	// 从节点0出发
	q := []int{0}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for nextNode, costTime := range g[node] {
			if d[node]+costTime == d[nextNode] {
				// 如果边 node -> nextNode 在最短路径组成的有向图上
				dp[nextNode] = (dp[nextNode] + dp[node]) % (1e9 + 7)
				deg[nextNode]--
				// 如果统计完了所有的指向nextNode的边，那么dp[nextNode]就已经确定了，将其加入到队列中
				if deg[nextNode] == 0 {
					q = append(q, nextNode)
				}
			}
		}
	}
	ans = dp[n-1]
	return
}

func CountPaths(n int, roads [][]int) int {
	return countPaths(n, roads)
}

// learned from: https://leetcode-cn.com/problems/number-of-ways-to-arrive-at-destination/solution/dan-yuan-zui-duan-lu-tuo-bu-xu-dp-by-end-i3ml/
