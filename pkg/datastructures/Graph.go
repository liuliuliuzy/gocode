package datastructures

// type GraphWithNoWeight struct {
// 	edges [][]bool
// 	vn    int
// }

type GraphWithWeight struct {
	edges [][]int
	n     int
}

func Constructor(n int, edges [][3]int) GraphWithWeight {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for _, edge := range edges {
		matrix[edge[0]][edge[1]] = edge[2]
	}
	g := GraphWithWeight{
		edges: matrix,
		n:     n,
	}
	return g
}

func (g GraphWithWeight) BFS(s int, operation func(int)) {
	n := g.n
	seen := make([]bool, n)
	q := []int{s}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		operation(v)
		seen[v] = true
		for next, e := range g.edges[v] {
			if e > 0 && !seen[next] {
				// 如果边存在
				q = append(q, next)
			}
		}
	}
}

func (g GraphWithWeight) DFS(s int, operation func(int)) {
	n := g.n
	seen := make([]bool, n)
	var dfs func(int)
	dfs = func(node int) {
		operation(node)
		seen[node] = true
		// 递归
		for next, e := range g.edges[node] {
			if e > 0 && !seen[next] {
				dfs(next)
			}
		}
	}
	dfs(s)
}

// dijkstra 算法求单源最短路径
func (g GraphWithWeight) DJ(i, j int) (ans int) {
	return
}
