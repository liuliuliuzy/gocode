package solutions

/*
给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。

请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-greater-element-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//O(n^2)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := []int{}
	dict := make(map[int]int)
	for _, x := range nums2 {
		dict[x] = -1
	}
	q := []int{}
	for _, x := range nums2 {
		if len(q) > 0 {
			for _, y := range q {
				if x > y && dict[y] < 0 {
					dict[y] = x
				}
			}
		}
		q = append(q, x)
	}
	for _, x := range nums1 {
		res = append(res, dict[x])
	}
	return res
}

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	return nextGreaterElement(nums1, nums2)
}
