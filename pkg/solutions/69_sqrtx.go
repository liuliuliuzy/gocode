package solutions

import "fmt"

//二分法计算平方根取整
func mySqrt(x int) int {
	left, right := 0, x
	fmt.Printf("%T\n", left)
	ans := left
	for left <= right {
		mid := (left + right) / 2
		// fmt.Println(left, right, mid, ans)
		if x >= mid*mid {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func MySqrt(x int) int {
	return mySqrt(x)
}
