package solutions

func Permutation(s string) []string {
	return permutation(s)
}

//简单的回溯
//考虑到s中可能有重复字符，所以不能简单地遍历每个字符然后选择是否使用
func permutation(s string) []string {
	visited := make([]bool, len(s))
	res := make([]string, 0)
	var backtrack func(int, string)
	backtrack = func(pos int, path string) {
		if pos == len(s) {
			if !isStringIn(res, path) {
				res = append(res, path)
			}
		}
		for i := 0; i < len(s); i++ {
			if !visited[i] {
				visited[i] = true
				backtrack(pos+1, path+string(s[i]))
				visited[i] = false
			}
		}
	}
	backtrack(0, "")
	return res
}

func isStringIn(lists []string, item string) bool {
	for _, i := range lists {
		if i == item {
			return true
		}
	}
	return false
}
