package solutions

//对每一行进行二分
//行开头元素大于target或者行结尾元素小于target则结束
// func searchMatrixII(matrix [][]int, target int) bool {
// 	m, n := len(matrix), len(matrix[0])
// 	for i := 0; i < m; i++ {
// 		if matrix[i][n-1] < target {
// 			continue
// 		}
// 		if matrix[i][0] > target {
// 			break
// 		}
// 		left, right := 0, n-1
// 		for left <= right {
// 			mid := (left + right) / 2
// 			if matrix[i][mid] < target {
// 				left = mid + 1
// 			} else if matrix[i][mid] == target {
// 				return true
// 			} else {
// 				right = mid - 1
// 			}
// 		}
// 	}
// 	return false
// }

//从右上角出发
//是比二分更好的方法
func searchMatrixII(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i < m && j > -1 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

func SearchMatrixII(matrix [][]int, target int) bool {
	return searchMatrixII(matrix, target)
}
