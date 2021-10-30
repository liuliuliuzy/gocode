package solutions

/*
给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。



进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//线性时间很好满足
// func singleNumber(nums []int) []int {
// 	seen := make(map[int]bool)
// 	for _, num := range nums {
// 		seen[num] = !seen[num]
// 	}
// 	res := []int{}
// 	for _, num := range nums {
// 		if seen[num] {
// 			res = append(res, num)
// 		}
// 	}
// 	return res
// }

func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum = xorSum ^ num
	}
	//得到了两个数异或的结果sum，接下来获取其二进制表示的最低非0位
	mask := 1
	for xorSum&mask == 0 {
		mask <<= 1
	}
	res := []int{0, 0}
	//根据这一结果，可以将所有的数分为 num&mask==0 num&mask==1
	//两组数一起异或就可以分别得到两个只出现一次的数
	for _, num := range nums {
		if num&mask == 0 {
			res[0] ^= num
		} else {
			res[1] ^= num
		}
	}
	return res
}

func SingleNumber(nums []int) []int {
	return singleNumber(nums)
}
