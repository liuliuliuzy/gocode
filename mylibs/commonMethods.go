package mylibs

/*一些common methods*/

func IsIn(slices [][2]int, item [2]int) (int, bool) {
	for i, items := range slices {
		if items == item {
			return i, true
		}
	}
	return -1, false
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func IsContain(arr []uint32, value uint32) bool {
	for i := 0; i < len(arr); i++ {
		if value == arr[i] {
			return true
		}
	}
	return false
}
