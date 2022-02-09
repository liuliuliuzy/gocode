package pathwithmaximumgold

// 不要觉得图遍历的复杂度很高

// 四个方向
// var dirs [4][2]int = [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func getMaximumGold(grid [][]int) (ans int) {
	var dfs func(x, y, gold int)
	dfs = func(x, y, gold int) {
		// 更新当前金子量
		gold += grid[x][y]
		// 更新结果
		if gold > ans {
			ans = gold
		}
		// 记录原始状态
		rec := grid[x][y]
		// 走过的方格不能再走
		grid[x][y] = 0
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			// 递归
			if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[0]) && grid[nx][ny] > 0 {
				dfs(nx, ny, gold)
			}
		}
		// 回溯
		grid[x][y] = rec
	}

	for i, row := range grid {
		for j, gold := range row {
			// 从有金子的方格开始
			if gold > 0 {
				// 初始金子量为0
				dfs(i, j, 0)
			}
		}
	}
	return
}

func GetMaximumGold(grid [][]int) int {
	return getMaximumGold(grid)
}
