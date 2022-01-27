package solutions

//要求：
//O(n)时间复杂度，O(1)空间复杂度
//后续：被空间复杂度要求给难住了，看一下官方题解...
func majorityElement(nums []int) int {
	var (
		count     = 0
		candidate = 0
		res       = 0
	)
	//Boyer-Moore 投票算法
	for _, item := range nums {
		if count == 0 {
			candidate = item
			count++
		} else {
			if item == candidate {
				count++
			} else {
				count--
			}
		}
	}

	//按照摩尔投票算法，经过一轮遍历之后，candidate一定是主要元素
	//或者是其它任意一个元素
	//如果candidate不是主要元素，那么就没有主要元素。
	//因为按照主要元素的定义，如果它存在，那么它一定是candidate
	count = 0
	for _, item := range nums {
		if item == candidate {
			count++
		}
	}

	if 2*count <= len(nums) {
		res = -1
	} else {
		res = candidate
	}
	return res
}

func MajorityElement(nums []int) int {
	return majorityElement(nums)
}
