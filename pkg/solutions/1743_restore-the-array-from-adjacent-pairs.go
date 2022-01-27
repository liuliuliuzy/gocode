package solutions

//n个不同元素组成的整数数组复原问题，“不同元素”是关键点
//解法1： 他妈的超时了
// func restoreArray(adjacentPairs [][]int) []int {
// 	ht := make(map[int][]int)
// 	for _, link := range adjacentPairs {
// 		ht[link[0]] = append(ht[link[0]], link[1])
// 		ht[link[1]] = append(ht[link[1]], link[0])
// 	}
// 	// fmt.Println(ht)
// 	res := adjacentPairs[0]
// 	length := 2
// 	for len(ht[res[length-1]]) > 1 {
// 		if ht[res[length-1]][0] != res[length-2] {
// 			res = append(res, ht[res[length-1]][0])
// 		} else {
// 			res = append(res, ht[res[length-1]][1])
// 		}
// 		length++
// 	}
// 	for len(ht[res[0]]) > 1 {
// 		if ht[res[0]][0] != res[1] {
// 			res = append([]int{ht[res[0]][0]}, res...)
// 		} else {
// 			res = append([]int{ht[res[0]][1]}, res...)
// 		}
// 		length++
// 	}
// 	return res
// }

func restoreArray(adjacentPairs [][]int) []int {
	ht := make(map[int][]int)
	for _, link := range adjacentPairs {
		ht[link[0]] = append(ht[link[0]], link[1])
		ht[link[1]] = append(ht[link[1]], link[0])
	}
	// fmt.Println(ht)
	length := len(adjacentPairs) + 1
	res := make([]int, length)

	for k, v := range ht {
		if len(v) == 1 {
			res[0] = k
			res[1] = v[0]
			break
		}
	}
	for i := 2; i < length; i++ {
		if res[i-2] != ht[res[i-1]][0] {
			res[i] = ht[res[i-1]][0]
		} else {
			res[i] = ht[res[i-1]][1]
		}
	}
	return res
}

func RestoreArray(adjacentPairs [][]int) []int {
	return restoreArray(adjacentPairs)
}
