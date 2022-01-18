package solutions

import (
	"math"
	"sort"
)

func getMinutes(t string) int {
	return (int(t[0]-'0')*10+int(t[1]-'0'))*60 + int(t[3]-'0')*10 + int(t[4]-'0')
}
func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	ans := math.MaxInt32
	lastMinute := getMinutes(timePoints[0])
	for i := 1; i < len(timePoints); i++ {
		if ans == 0 {
			// 0 is the smallest result, so we could return it
			return ans
		}
		// calculate the minutes of now moment
		minutes := getMinutes(timePoints[i])
		// use the smaller value
		if minutes-lastMinute < ans {
			ans = minutes - lastMinute
		}
		// set now moment value to lastMinute
		lastMinute = minutes
	}
	if getMinutes(timePoints[0])+1440-lastMinute < ans {
		ans = getMinutes(timePoints[0]) + 1440 - lastMinute
	}
	return ans
}

func FindMinDifference(timePoints []string) int {
	return findMinDifference(timePoints)
}
