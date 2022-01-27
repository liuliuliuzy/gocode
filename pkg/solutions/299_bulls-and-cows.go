package solutions

import "strconv"

/*

虽然这题目翻译的很垃圾，但是看了1分钟终于还是看懂了

提示：

    1 <= secret.length, guess.length <= 1000
    secret.length == guess.length
    secret 和 guess 仅由数字组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bulls-and-cows
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//纯纯的哈希表
func getHint(secret string, guess string) string {
	bullsCount := 0
	cowsCount := len(secret)
	otherDigitsCount := make(map[byte]int)
	for i := range secret {
		if secret[i] == guess[i] {
			bullsCount += 1
			cowsCount -= 1
		} else {
			otherDigitsCount[secret[i]] += 1
			otherDigitsCount[guess[i]] -= 1
		}
	}
	c := byte('0')
	for c < 0x3a {
		if otherDigitsCount[c] > 0 {
			cowsCount -= otherDigitsCount[c]
		}
		c += 1
	}
	return strconv.Itoa(bullsCount) + "A" + strconv.Itoa(cowsCount) + "B"
}

func GetHint(secret string, guess string) string {
	return getHint(secret, guess)
}
