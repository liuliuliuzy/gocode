package mylibs

func IsIn(slices [][2]int, item [2]int) (int, bool) {
	for i, items := range slices {
		if items == item {
			return i, true
		}
	}
	return -1, false
}
