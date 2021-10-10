package solutions

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid/n][mid%n] < target {
			left = mid + 1
		} else if matrix[mid/n][mid%n] == target {
			return true
		} else {
			right = mid - 1
		}
	}
	return false
}

func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}
