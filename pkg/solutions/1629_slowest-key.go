package solutions

import "fmt"

// 数组 字符串
func slowestKey(releaseTimes []int, keysPressed string) byte {
	// ans := keysPressed[0]
	// count := releaseTimes[0]
	// for i := 1; i < len(keysPressed); i++ {
	// 	tmp := releaseTimes[i] - releaseTimes[i-1]
	// 	fmt.Println(tmp, keysPressed[i])
	// 	if tmp > count {
	// 		ans = keysPressed[i]
	// 		count = tmp
	// 	} else {
	// 		if tmp == count && keysPressed[i] > ans {
	// 			ans = keysPressed[i]
	// 		}
	// 	}
	// }
	// return ans
	ans := keysPressed[0]
	count := releaseTimes[0]<<8 + int(keysPressed[0])
	fmt.Println(count)
	for i := 1; i < len(keysPressed); i++ {
		tmp := (releaseTimes[i]-releaseTimes[i-1])<<8 + int(keysPressed[i])
		fmt.Println(i, releaseTimes[i]-releaseTimes[i-1], tmp)
		if tmp > count {
			ans = keysPressed[i]
			count = tmp
		}
	}
	return ans
}

func SlowestKey(releaseTimes []int, keysPressed string) byte {
	return slowestKey(releaseTimes, keysPressed)
}
