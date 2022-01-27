package solutions

func isWin(role uint8, board []string) bool {
	if board[0][0] == role && board[0][1] == role {
		if board[0][2] == role {
			return true
		}
	} else if board[1][0] == role && board[1][1] == role {
		if board[1][2] == role {
			return true
		}
	} else if board[2][0] == role && board[2][1] == role {
		if board[2][2] == role {
			return true
		}
	} else if board[0][0] == role && board[1][0] == role {
		if board[2][0] == role {
			return true
		}
	} else if board[0][1] == role && board[1][1] == role {
		if board[2][1] == role {
			return true
		}
	} else if board[0][2] == role && board[1][2] == role {
		if board[2][2] == role {
			return true
		}
	} else if board[0][0] == role && board[1][1] == role {
		if board[2][2] == role {
			return true
		}
	} else if board[2][0] == role && board[1][1] == role {
		if board[0][2] == role {
			return true
		}
	} else {
		return false
	}
	return false
}

/*
有些坑点：
'X'是先下的一方
*/
func validTicTacToe(board []string) bool {
	xCount, oCount := 0, 0
	// count 'X' and 'O'
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch board[i][j] {
			case 'X':
				xCount++
			case 'O':
				oCount++
			}
		}
	}
	if xCount-oCount < 0 || xCount-oCount > 1 {
		return false
	}
	// 判断胜利条件是否成立
	// 有8种胜利场景，直接枚举就行吧
	if isWin('X', board) {
		if xCount-oCount != 1 || isWin('O', board) {
			return false
		}
	}
	if isWin('O', board) {
		if xCount != oCount || isWin('X', board) {
			return false
		}
	}
	return true
}

func ValidTicTacToe(board []string) bool {
	return validTicTacToe(board)
}
