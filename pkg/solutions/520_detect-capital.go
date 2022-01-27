package solutions

func detectCapitalUse(word string) bool {
	//根据前两位的大小写情况确定整个单词中应该的大小写情况
	if len(word) == 1 {
		return true
	}
	if word[0] < 'a' { // 首字母大写
		if word[1] < 'a' { // 第二个字母也是大写
			for i := 2; i < len(word); i++ { // 那么所有字母都应该大写
				if word[i] >= 'a' {
					return false
				}
			}
		} else { // 首字母小写
			for i := 2; i < len(word); i++ {
				if word[i] < 'a' {
					return false
				}
			}
		}
	} else { // 首字母小写，那么所有字母的都要小写
		for i := 1; i < len(word); i++ {
			if word[i] < 'a' {
				return false
			}
		}
	}
	return true
}

func DetectCapitalUse(word string) bool {
	return detectCapitalUse(word)
}
