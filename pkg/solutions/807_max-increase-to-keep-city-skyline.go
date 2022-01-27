package solutions

// 我想到了
func maxIncreaseKeepingSkyline(grid [][]int) int {
	ans := 0
	var n int = len(grid)
	maxrows := make([]int, n)
	maxcols := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > maxrows[i] {
				maxrows[i] = grid[i][j]
			}
			if grid[j][i] > maxcols[i] {
				maxcols[i] = grid[j][i]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += min(maxrows[i]-grid[i][j], maxcols[j]-grid[i][j])
		}
	}
	return ans
}

func MaxIncreaseKeepingSkyline(grid [][]int) int {
	return maxIncreaseKeepingSkyline(grid)
}
