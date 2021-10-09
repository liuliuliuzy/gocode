package solutions

func destCity(paths [][]string) string {
	ht := make(map[int]bool)
	// start := paths[0][0]
	next := paths[0][1]
	ht[0] = true
	for {
		i := 0
		for i < len(paths) {
			if !ht[i] {
				//如果线路的起点就是上条线路的目的地
				if paths[i][0] == next {
					// start = paths[i][0]
					next = paths[i][1]
					ht[i] = true
					break
				}
			}
			i++
		}
		if i == len(paths) {
			break
		}
	}
	return next
}

func DestCity(paths [][]string) string {
	return destCity(paths)
}
