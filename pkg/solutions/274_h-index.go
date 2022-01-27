package solutions

import "sort"

func hIndex(citations []int) int {
	res := 0
	sort.Ints(citations)
	startIndex := 0
	length := len(citations)
	for startIndex < length {
		if citations[startIndex] < length-startIndex {
			startIndex++
			continue
		}
		res = length - startIndex
		break
	}
	return res
}

func HIndex(citations []int) int {
	return hIndex(citations)
}
