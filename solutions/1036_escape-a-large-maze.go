package solutions

// BFS search
// 纯广度优先可能会超时吧？
func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	// BFS search
	bfsQ := [][]int{}
    var getNext(start [][]int) (bool, []int)
    getnext = func (start []int) (bool, []int) {
        dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
        res := [][]int{}
        for _, dir := range dir {
            tmpX := dir[0] + start[0]
            tmpY := dir[1] + start[1]
            if 
        }
    }

	return false
}

func contains(blocked [][]int, p []int) bool {
    for _, point := range blocked {
        if point[0] == p[0] && point[1] == p[1] {
            return true
        }
    }
    return false
}

func IsEscapePossible(blocked [][]int, source []int, target []int) bool {
	return isEscapePossible(blocked, source, target)
}
