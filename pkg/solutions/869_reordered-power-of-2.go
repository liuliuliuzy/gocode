package solutions

import (
	"math"
	"strconv"
)

//由于题目给定了n的范围[1, 1e9]，所以直接可以知道这个范围内所有的2的幂次。
//于是直接哈希表就可以做了。
func countDigits(n int) [10]int {
	res := [10]int{}
	for _, item := range strconv.Itoa(n) {
		// fmt.Println(item)
		res[item-0x30] += 1
	}
	return res
}

func reorderedPowerOf2(n int) bool {
	//哈希表
	hashTable := make(map[[10]int]bool)
	// fmt.Println(countDigits(113143))
	i := 0
	for {
		if int(math.Pow(2, float64(i))) > 1e9 {
			break
		}
		hashTable[countDigits(int(math.Pow(2, float64(i))))] = true
		i += 1
	}
	// fmt.Println(hashTable)
	return hashTable[countDigits(n)]
}

func ReorderedPowerOf2(n int) bool {
	return reorderedPowerOf2(n)
}
