package numberofvalidwordsinasentence

import (
	"strings"
)

// 字符串处理函数
func countValidWords(sentence string) int {
	split := strings.Fields(sentence)
	a := 0
	for _, candidate := range split {
		if isValidWord(candidate) {
			a++
		}
	}
	return a
}

func isValidWord(word string) bool {
	hasMeetDash := false
	for i := 0; i < len(word); i++ {
		if word[i] >= '0' && word[i] <= '9' { // 不能含有数字
			return false
		} else if word[i] == '-' { // '-' 两边只能是小写字母
			if i == 0 || i == len(word)-1 || hasMeetDash || word[i-1] > 'z' || word[i-1] < 'a' || word[i+1] > 'z' || word[i+1] < 'a' {
				return false
			}
			hasMeetDash = true
		} else if word[i] == '.' || word[i] == '!' || word[i] == ',' { // 标点符号只能位于末尾
			if i != len(word)-1 {
				return false
			}
		} else {
			continue
		}
	}
	return true
}

func CountValidWords(sentence string) int {
	return countValidWords(sentence)
}
