package solutions

//数学题，从全部从左到右标号的二叉树出发，寻找思路
func pathInZigZagTree(label int) []int {
	//首先确定在第几行
	i := 0
	tmp := label
	for tmp > 0 {
		tmp = tmp >> 1
		i++
	}
	res := make([]int, i)
	start := label
	if i%2 == 0 {
		start = (1<<i + 1<<(i-1) - 1) - label
	}
	for i > 0 {
		if i%2 == 0 {
			res[i-1] = (1<<i + 1<<(i-1) - 1) - start
		} else {
			res[i-1] = start
		}
		start = start / 2
		i--
	}

	return res
}

func PathInZigZagTree(label int) []int {
	return pathInZigZagTree(label)
}
