package findtheminimumnumberoffibonaccinumberswhosesumisk

/*
给你数字 k ，请你返回和为 k 的斐波那契数字的最少数目，其中，每个斐波那契数字都可以被使用多次。
*/
func findMinFibonacciNumbers(k int) int {
	// 这题的标签是贪心算法
	// 所以策略是尽可能用大的数
	// 贪心算法似乎和数学计算强相关啊
	// generate fibonacci numbers
	fib := []int{1, 1}
	length := 2
	for fib[length-1]+fib[length-2] <= k {
		fib = append(fib, fib[length-1]+fib[length-2])
		length++
	}
	ans := 0
	for k > 0 {
		k -= findMaxOfK(k, fib)
		ans++
	}
	return ans
}
func findMaxOfK(k int, fibs []int) int {
	i := len(fibs) - 1
	for fibs[i] > k {
		i--
	}
	return fibs[i]
}

func FindMinFibonacciNumbers(k int) int {
	return findMinFibonacciNumbers(k)
}
