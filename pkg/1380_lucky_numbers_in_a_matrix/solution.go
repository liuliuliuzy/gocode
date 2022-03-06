package luckynumbersinamatrix

import (
	"math"
)

func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	ans := []int{}
	judge := make([][]bool, m)
	for i := 0; i < m; i++ {
		judge[i] = make([]bool, n)
	}
	// 第一遍：遍历每一行
	for i, row := range matrix {
		minValue := math.MaxInt
		minIndex := -1
		for j, v := range row {
			if v < minValue {
				minValue = v
				minIndex = j
			}
		}
		judge[i][minIndex] = true
	}
	// 第二遍：遍历每一列
	for j := 0; j < n; j++ {
		maxValue := -1
		maxIndex := -1
		for i := 0; i < m; i++ {
			if matrix[i][j] > maxValue {
				maxValue = matrix[i][j]
				maxIndex = i
			}
		}
		// 逐列对bool值进行操作
		for i := 0; i < m; i++ {
			if i == maxIndex {
				judge[i][j] = judge[i][j] && true
			} else {
				judge[i][j] = judge[i][j] && false
			}
		}
	}
	for i, row := range judge {
		for j, b := range row {
			if b {
				ans = append(ans, matrix[i][j])
			}
		}
	}
	return ans
}

func LuckyNumbers(matrix [][]int) []int {
	return luckyNumbers(matrix)
}
