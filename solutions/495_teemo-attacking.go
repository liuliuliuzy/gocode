package solutions

func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := 0
	i := 1
	ans += duration
	nextMoment := timeSeries[0] + duration - 1
	for i < len(timeSeries) {
		if timeSeries[i] <= nextMoment {
			ans += duration - (nextMoment - timeSeries[i] + 1)
		} else {
			ans += duration
		}
		nextMoment = timeSeries[i] + duration - 1
		i++
	}
	return ans
}

func FindPoisonedDuration(timeSeries []int, duration int) int {
	return findPoisonedDuration(timeSeries, duration)
}
