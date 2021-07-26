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

//辗转相除
func Gcd(a, b int) int {
	var tmp int
	for {
		tmp = a % b
		if tmp > 0 {
			a = b
			b = tmp
		} else {
			return b
		}
	}
}

func QuickSort(data []int, start, end int) {
	if start < end {
		midIndex := findIndex(data, start, end)
		QuickSort(data, start, midIndex-1)
		QuickSort(data, midIndex+1, end)
	}
}

func findIndex(data []int, start, end int) int {
	//哨兵元素
	tmp := data[start]
	l := start
	r := end
	for l < r {
		for l < r && data[r] >= tmp {
			r--
		}
		data[l] = data[r]
		for l < r && data[l] <= tmp {
			l++
		}
		data[r] = data[l]
	}
	//放置哨兵元素
	data[l] = tmp
	return l
}

//nums为递增数组，返回nums中最小的大于等于target的元素索引，如果不存在，则返回len(nums)-1
func MyBiSearch(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		//这一步取整操作，mid>=left
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] < target {
		left++
	}
	return left
}
