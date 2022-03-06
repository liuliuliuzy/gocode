package numberofenclaves

var dirs [4][2]int = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	onesCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				onesCount++
			}
		}
	}
	seen := make([][]bool, m)
	for i := 0; i < m; i++ {
		seen[i] = make([]bool, n)
	}
	var dfs func(int, int) // 函数签名只需要写出函数参数类型以及返回值类型就行
	dfs = func(i, j int) {
		seen[i][j] = true
		onesCount--
		for _, d := range dirs {
			newI, newJ := i+d[0], j+d[1]
			// 前进的条件
			if newI >= 0 && newI < m && newJ >= 0 && newJ < n && grid[newI][newJ] == 1 && !seen[newI][newJ] {
				// 递归，深度优先
				dfs(newI, newJ)
			}
		}
	}
	for i := 0; i < n; i++ {
		if grid[0][i] == 1 && !seen[0][i] {
			dfs(0, i)
		}
		if grid[m-1][i] == 1 && !seen[m-1][i] {
			dfs(m-1, i)
		}
	}
	for j := 1; j < m-1; j++ {
		if grid[j][0] == 1 && !seen[j][0] {
			dfs(j, 0)
		}
		if grid[j][n-1] == 1 && !seen[j][n-1] {
			dfs(j, n-1)
		}
	}
	return onesCount
}

func NumEnclaves(grid [][]int) int {
	return numEnclaves(grid)
}
