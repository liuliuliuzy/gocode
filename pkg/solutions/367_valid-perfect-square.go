package solutions

//简单易懂，时间O(logn)
// func isPerfectSquare(num int) bool {
// 	start := 1
// 	for start*start < num {
// 		start += 1
// 	}
// 	return start*start == num
// }

//二分，同样O(logn)
func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (left + right) / 2
		sq := mid * mid
		if sq < num {
			left = mid + 1
		} else if sq > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

func IsPerfectSquare(num int) bool {
	return isPerfectSquare(num)
}
