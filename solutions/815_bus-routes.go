package solutions

//广度优先吧，得先确定一个终止条件
func numBusesToDestination(routes [][]int, source int, target int) int {
	//map记录各个车站是否到过
	visited := map[int]bool{source: true}
	var nextStations func(status int) {
		return {}
	}
}

func NumBusesToDestination(routes [][]int, source int, target int) int {
	return numBusesToDestination(routes, source, target)
}
