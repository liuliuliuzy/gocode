package solutions

// import (
// 	"fmt"
// 	"math"
// )

// // BFS search
// // 纯广度优先可能会超时吧？
// func isEscapePossible(blocked [][]int, source []int, target []int) bool {
// 	// 目标节点与起点之间的距离
// 	distance := (source[0]-target[0])*(source[0]-target[0]) + (source[1]-target[1])*(source[1]-target[1])
// 	// BFS search
// 	bfsQ := [][]int{}
// 	// 哈希表纪录走过的点
// 	// 但是这操作直接就内存爆炸了...
// 	seen := make([][]bool, 100)
// 	for i := range seen {
// 		seen[i] = make([]bool, 100)
// 	}
// 	fmt.Println(len(seen))
// 	// 获取下个点
// 	var getNext = func(start []int) (bool, [][]int) {
// 		// getNext = func(start []int) (bool, [][]int) {
// 		dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
// 		res := [][]int{}
// 		for _, dir := range dirs {
// 			tmpX := dir[0] + start[0]
// 			tmpY := dir[1] + start[1]
// 			// 判断是否合法
// 			if !contains(blocked, tmpX, tmpY) && tmpX >= 0 && tmpX <= 999999 && tmpY >= 0 && tmpY <= 999999 {
// 				res = append(res, []int{tmpX, tmpY})
// 			}
// 		}
// 		return len(res) > 0, res
// 	}
// 	// 如果没有阻拦
// 	// if len(blocked) == 0 {
// 	// 	return true
// 	// }
// 	// 开始搜索
// 	bfsQ = append(bfsQ, source)
// 	seen[source[0]][source[1]] = true
// 	for len(bfsQ) > 0 {
// 		// // 队列中取出一个点
// 		p := bfsQ[0]
// 		bfsQ = bfsQ[1:]

// 		// seen[p[0]][p[1]] = true
// 		flag, candidates := getNext(p)
// 		if flag {
// 			for _, nextP := range candidates {
// 				if !seen[nextP[0]][nextP[1]] {
// 					// 插入队列
// 					bfsQ = append(bfsQ, nextP)
// 					// 标记
// 					seen[nextP[0]][nextP[1]] = true
// 					// 搜索到节点时进行的操作
// 					fmt.Println("searched: ", nextP)
// 					if nextP[0] == target[0] && nextP[1] == target[1] {
// 						return true
// 					}
// 					// 有限步数判断
// 					if (nextP[0]-source[0])*(nextP[0]-source[0])+(nextP[1]-source[1])*(nextP[1]-source[1]) > 2*distance+3*int(math.Sqrt(float64(distance)))+1 {
// 						fmt.Println("hit:", nextP)
// 						return false
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return false
// }

// // 判断(x, y)是否在block列表中
// func contains(blocked [][]int, x, y int) bool {
// 	for _, point := range blocked {
// 		if point[0] == x && point[1] == y {
// 			return true
// 		}
// 	}
// 	return false
// }

// 还是不太对，看看官方题解吧
// https://leetcode-cn.com/problems/escape-a-large-maze/solution/tao-chi-da-mi-gong-by-leetcode-solution-qxhz/

// 定义坐标类型
type point struct{ x, y int }

// 分别对应坐标系中的左、右、下、上四个方向
var dirs = []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func isEscapePossible(block [][]int, source, target []int) bool {
	// 定义常量
	const (
		blocked = -1 // 在包围圈中
		valid   = 0  // 不在包围圈中
		found   = 1  // 无论在不在包围圈中，但在 n(n-1)/2 步搜索的过程中经过了 target

		boundary = 1e6
	)

	// 计算包围圈中最多能围住多少个坐标
	n := len(block)
	// 0、1个block无法形成包围圈
	if n < 2 {
		return true
	}
	blockSet := map[point]bool{}
	for _, b := range block {
		blockSet[point{b[0], b[1]}] = true
	}

	// 限制了步数的BFS
	check := func(start, finish []int) int {
		sx, sy := start[0], start[1]
		fx, fy := finish[0], finish[1]
		countdown := n * (n - 1) / 2

		q := []point{{sx, sy}}
		vis := map[point]bool{{sx, sy}: true}
		for len(q) > 0 && countdown > 0 {
			// 队列中取出坐标
			p := q[0]
			q = q[1:]
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				np := point{x, y}
				if 0 <= x && x < boundary && 0 <= y && y < boundary && !blockSet[np] && !vis[np] {
					// 搜索过程中遇到了目标节点
					if x == fx && y == fy {
						return valid
					}
					countdown--
					vis[np] = true
					q = append(q, np)
				}
			}
		}
		// 被包围了
		if countdown > 0 {
			return blocked
		}
		// 没有被包围
		return valid
	}
	res := check(source, target)
	// 返回true的情况：
	// 1、从source出发寻找target的过程中找到了target
	// 2、从target出发寻找source的过程中找到了source
	// 3、source和target都没有处于包围圈中
	return res == found || res == valid && check(target, source) != blocked
}

func IsEscapePossible(blocked [][]int, source []int, target []int) bool {
	return isEscapePossible(blocked, source, target)
}
