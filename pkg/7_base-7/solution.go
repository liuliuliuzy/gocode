package base7

import "strconv"

func convertToBase(num int) (ans string) {
	if num == 0 {
		return "0"
	}
	flag := false
	if num < 0 {
		num = -num
		flag = true
	}
	for num > 0 {
		ans = strconv.Itoa(num%7) + ans
		num = num / 7
	}
	if flag {
		ans = "-" + ans
	}
	return
}

func ConvertToBase(num int) string {
	return convertToBase(num)
}
