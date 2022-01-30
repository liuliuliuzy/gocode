package thenumberofweakcharactersinthegame

import (
	"fmt"
	"sort"
)

func numberOfWeakCharacters(properities [][]int) int {
	// 按照攻击力降序排序
	sort.Slice(properities, func(i, j int) bool {
		return properities[i][0] > properities[j][0] || properities[i][0] == properities[j][0] && properities[i][1] < properities[j][1]
	})
	fmt.Println(properities)
	tmp := properities[0][1]
	a := 0
	for i := 1; i < len(properities); i++ {
		if tmp > properities[i][1] {
			a++
		} else {
			tmp = properities[i][1]
		}
	}
	return a
}

func NumberOfWeakCharacters(properities [][]int) int {
	return numberOfWeakCharacters(properities)
}
