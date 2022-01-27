package solutions

//跟168刚好是反着来的
func titleToNumber(columnTitle string) int {
	res := 0
	for i := 0; i < len(columnTitle); i++ {
		res = res*26 + int((columnTitle[i] - 'A')) + 1
	}
	return res
}

func TitleToNumber(columnTitle string) int {
	return titleToNumber(columnTitle)
}
