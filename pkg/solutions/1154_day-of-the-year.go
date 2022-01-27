package solutions

import "strconv"

func dayOfYear(date string) int {
	counts := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	ans := 0
	year, _ := strconv.Atoi(date[0:4])
	mounth, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:10])
	if mounth > 2 && year%4 == 0 {
		counts[2] += 1
	}
	for i := mounth; i > 0; i-- {
		ans += counts[i-1]
	}
	ans += day
	return ans
}

func DayOfYear(date string) int {
	return dayOfYear(date)
}
