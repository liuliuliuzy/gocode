package solutions

//这题理解题意占了40%可能...
func minCharacters(a string, b string) int {
	ca := make([]int, 26)
	cb := make([]int, 26)
	for _, char := range a {
		ca[char-'a']++
	}
	for _, char := range b {
		cb[char-'a']++
	}
	lenA, lenB := len(a), len(b)
	res := lenA + lenB
	var sumA, sumB int = 0, 0
	for i := 0; i < 25; i++ {
		//计算前缀和
		sumA += ca[i]
		sumB += cb[i]
		res = min(min(res, lenA+lenB-ca[i]-cb[i]), min(sumA+lenB-sumB, sumB+lenA-sumA))
	}
	res = min(res, lenA+lenB-ca[25]-cb[25])
	return res
}

func MinCharacters(a, b string) int {
	return minCharacters(a, b)
}
