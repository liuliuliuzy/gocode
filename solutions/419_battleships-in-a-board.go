package solutions

// 不考虑空间复杂度O(1)的限制
func countBattleships(board [][]byte) int {
	m, n := len(board), len(board[0])
	result := 0
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] {
				continue
			}
			visited[i][j] = true
			if board[i][j] == 'X' {
				result++
				// 继续搜索，将战舰剩余的点也都标记为visited
				if i+1 < m && board[i+1][j] == 'X' {
					tmp := i + 1
					for tmp < m && board[tmp][j] == 'X' {
						visited[tmp][j] = true
						tmp++
					}
				} else if j+1 < n && board[i][j+1] == 'X' {
					tmp := j + 1
					for tmp < n && board[i][tmp] == 'X' {
						visited[i][tmp] = true
						tmp++
					}
				} else {
					continue
				}
			}
		}
	}
	return result
}

func CountBattleships(board [][]byte) int {
	return countBattleships(board)
}
