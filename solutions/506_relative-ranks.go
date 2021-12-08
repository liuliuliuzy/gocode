/*
字典快排
*/
package solutions

import (
	"strconv"
)

type element struct {
	score int
	index int
}

func qc(elements []element, start, end int) {
	if start < end {
		ep := elements[start]
		l, r := start, end
		for l < r {
			for l < r && elements[r].score <= ep.score {
				r--
			}
			elements[l] = elements[r]
			for l < r && elements[l].score >= ep.score {
				l++
			}
			elements[r] = elements[l]
		}
		elements[l] = ep
		// fmt.Println(elements)
		qc(elements, start, l-1)
		qc(elements, l+1, end)
	}
}

func findRelativeRanks(score []int) []string {
	elements := []element{}
	for i, s := range score {
		elements = append(elements, element{s, i})
	}
	// quick sort
	qc(elements, 0, len(elements)-1)
	// fmt.Println(elements)
	ans := make([]string, len(score))
	ranks := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	for i, e := range elements {
		if i < 3 {
			ans[e.index] = ranks[i]
		} else {
			ans[e.index] = strconv.Itoa(i + 1)
		}
	}
	return ans
}

func FindRelativeRanks(score []int) []string {
	return findRelativeRanks(score)
}
