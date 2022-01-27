package solutions

//遍历 + 哈希表
//艹，每个单词里面可以有重复字母
func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	ht := make(map[[26]int]int)
	for _, value := range strs {
		tmpIndex := [26]int{}
		for _, char := range value {
			tmpIndex[char-97]++
		}
		if ht[tmpIndex] == 0 {
			ht[tmpIndex] = len(res) + 1
			res = append(res, []string{value})
		} else {
			res[ht[tmpIndex]-1] = append(res[ht[tmpIndex]-1], value)
		}
	}
	return res
}

func GroupAnagrams(strs []string) [][]string {
	return groupAnagrams(strs)
}
