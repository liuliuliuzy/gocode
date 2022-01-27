package solutions

func toHex(num int) string {
	convert := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f"}
	ans := ""
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = 0x100000000 + num
	}
	for num > 0 {
		ans = convert[num%16] + ans
		num = num / 16
	}
	return ans
}

func ToHex(num int) string {
	return toHex(num)
}
