package solutions

// 哈希表，两次循环，真实
// func findErrorNums(nums []int) []int {
//     res := []int{}
//     mp := map[int]bool{}
//     n := 0
//     for _, number := range nums {
//         if mp[number] {
//             res = append(res, number)
//             continue
//         }
//         mp[number] = true
//         if n < number {
//             n = number
//         }
//     }
//     for i:=1; i<n; i++ {
//         if !mp[i] {
//             res = append(res, i)
//             break
//         }
//     }
//     //出现了特殊情况
//     if len(res) < 2 {
//         res = append(res, n+1)
//     }
//     return res
// }

//数学解法
func findErrorNums(nums []int) []int {
	res := []int{}
	mp := map[int]bool{}
	sum := 0
	for _, number := range nums {
		if mp[number] {
			res = append(res, number)
			continue
		}
		mp[number] = true
		sum += number
	}
	res = append(res, (1+len(nums))*len(nums)/2-sum)
	return res
}

func FindErrorNums(nums []int) []int {
	return findErrorNums(nums)
}
