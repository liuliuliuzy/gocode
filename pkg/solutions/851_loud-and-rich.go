package solutions

type person struct {
	id    int
	value int
}

/*
有向图的图算法
*/
func loudAndRich(richer [][]int, quiet []int) []int {
	// 创建有向图
	n := len(quiet)
	directedG := make([][]int, n)
	seen := make([]bool, n)
	// 邻接表形式
	for _, link := range richer {
		// link中的元素表示：link[0] 比 link[1] 更富有
		// 那么我们就创建一条从link[1]到link[0]的有向边
		directedG[link[1]] = append(directedG[link[1]], link[0])
		// 将link[0]标记为非开始节点
		seen[link[0]] = true
	}
	ans := make([]int, n)
	// 记录每个节点对应的符合条件最安静的人
	persons := make([]person, n)
	// 开始节点集合
	starts := []int{}
	for i := range seen {
		if seen[i] {
			seen[i] = false
			continue
		}
		starts = append(starts, i)
	}
	var dfs func(node int) person
	dfs = func(node int) person {
		// 如果访问过了，那么直接返回先前存储的结果
		if seen[node] {
			return persons[node]
		}
		// 标记为已经访问过
		seen[node] = true
		// 如果到达了边缘节点（出度为0）
		if len(directedG[node]) == 0 {
			persons[node] = person{
				node,
				quiet[node],
			}
			ans[node] = node
			return persons[node]
		}
		p := person{
			node,
			quiet[node],
		}
		// 深度优先搜索
		for _, next := range directedG[node] {
			var tmp person
			if seen[next] {
				tmp = persons[next]
			} else {
				tmp = dfs(next)
			}
			if p.value > tmp.value {
				p = tmp
			}
		}
		persons[node] = p
		ans[node] = p.id
		return p
	}
	// 从开始节点集合开始遍历，深度优先
	for _, snode := range starts {
		dfs(snode)
	}
	return ans
}

// func helper(dg [][]int, quiet []int, ans []int, start int) person {
// 	// 终止条件
// 	if len(dg[start]) == 0 {
// 		ans[start] = start
// 		p := person{
// 			start,
// 			quiet[start],
// 		}
// 		return p
// 	}
// 	// value := quiet[start]
// 	// id := start
// 	p := person{
// 		start,
// 		quiet[start],
// 	}
// 	for next := range dg[start] {
// 		tmpP := helper(dg, quiet, ans, next)
// 		if tmpP.value < p.value {
// 			p = tmpP
// 		}
// 	}
// 	ans[start] = p.id
// 	return p
// }

func LoudAndRich(richer [][]int, quiet []int) []int {
	return loudAndRich(richer, quiet)
}
