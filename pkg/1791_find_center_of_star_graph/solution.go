package findcenterofstargraph

func findCenter(edges [][]int) (ans int) {
	seen := make(map[int]bool)
	for _, edge := range edges {
		s, d := edge[0], edge[1]
		if seen[s] {
			ans = s
			break
		}
		seen[s] = true
		if seen[d] {
			ans = d
			break
		}
		seen[d] = true
	}
	return
}

func FindCenter(edges [][]int) int {
	return findCenter(edges)
}
