package solutions

//感觉本质上和走方格是一样的
//4位数字
//广度优先
func openLock(deadends []string, target string) int {
	res := 0
	dis := 0
	targetDis := 0
	for i := 0; i < len(target); i++ {
		targetDis += int(target[i]) - 48
	}
	start := "0000"
	for dis < targetDis && start != target {

	}
	if dis == targetDis {
		return -1
	}
	return res
}

func OpenLock(deadends []string, target string) int {
	return openLock(deadends, target)
}
