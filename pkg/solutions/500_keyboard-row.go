package solutions

//纯纯的哈希表做法
//需要注意大小写
func findWords(words []string) []string {
	keyboard := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	runeHT := make(map[rune]int)
	for i, row := range keyboard {
		for _, r := range row {
			runeHT[r] = i
			runeHT[r-32] = i //转大写
		}
	}
	res := []string{}
	for _, word := range words {
		index := -1
		flag := true
		for _, r := range word {
			//set initial value
			if index < 0 {
				index = runeHT[r]
				continue
			}
			if runeHT[r] != index {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, word)
		}
	}
	return res
}

func FindWords(words []string) []string {
	return findWords(words)
}
