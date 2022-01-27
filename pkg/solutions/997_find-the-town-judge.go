package solutions

// 我自己的愚蠢解法...
// func findJudge(n int, trust [][]int) int {
// 	graph := make([][]int, n+1)
// 	for i := range graph {
// 		graph[i] = make([]int, n+1)
// 	}
// 	for i := range trust {
// 		graph[trust[i][0]][trust[i][1]] = 1
// 	}
// 	ans := -1
// 	for i := 1; i < n+1; i++ {
// 		// // 发现一个可能是小镇法官的人
// 		// if i > 1 && graph[1][i] == 1 {
// 		j := 1
// 		for ; j < n+1; j++ {
// 			if j == i && graph[j][i] == 1 {
// 				// 不符合条件1
// 				break
// 			}
// 			if j != i && graph[j][i] != 1 {
// 				// 不符合条件2
// 				break
// 			}
// 		}
// 		if j < n+1 {
// 			continue
// 		}
// 		// 初步通过判断
// 		// 检查看他有没有信任其他人
// 		flag := true
// 		for k := 1; k < n+1; k++ {
// 			if graph[i][k] == 1 {
// 				flag = false
// 				break
// 			}
// 		}
// 		if !flag {
// 			continue
// 		}
// 		if ans > 0 {
// 			// 但是已经有了一个小镇法官了，所以这里不符合条件3
// 			ans = -1
// 			break
// 		}
// 		ans = i
// 		// }
// 	}
// 	return ans
// }

func findJudge(n int, trust [][]int) int {
	// 统计入度和出度
	count := make([][]int, n+1)
	for i := range count {
		count[i] = make([]int, 2)
	}
	for i := range trust {
		count[trust[i][0]][0]++ // 出度++
		count[trust[i][1]][1]++ // 入度++
	}
	// fmt.Println(count)
	ans := -1
	for i := 1; i < n+1; i++ {
		if count[i][0] == 0 && count[i][1] == n-1 {
			if ans > 0 {
				ans = -1
				break
			}
			ans = i
		}
	}
	return ans
}

func FindJudge(n int, trust [][]int) int {
	return findJudge(n, trust)
}
