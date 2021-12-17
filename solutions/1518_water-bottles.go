package solutions

import "math"

// 空瓶兑换酒，但是不能借瓶子
func numWaterBottles(numBottles int, numExchange int) int {
	ans := numBottles
	for numBottles >= numExchange {
		remain := numBottles % numExchange
		quotient := int(math.Floor(float64(numBottles) / float64(numExchange)))
		ans += quotient
		numBottles = quotient + remain
	}
	return ans
}

func NumWaterBottles(numBottles int, numExchange int) int {
	return numWaterBottles(numBottles, numExchange)
}
