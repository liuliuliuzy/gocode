package solutions

import "sort"

// houses 位于一条线上，heaters 可能有多个
// 对于这种问题，该用什么算法呢
// 贪心？
// 以下方法大失败，学习官方题解，反思一下问题在哪里

// func findRadius(houses []int, heaters []int) int {
// 	m := len(houses)
// 	n := len(heaters)
// 	radius := 0
// 	count := make(map[int]bool)
// 	for _, h := range houses {
// 		count[h] = true
// 	}
// 	maxInner := 0
// 	for i := 1; i < len(heaters); i++ {
// 		// space := int(math.Floor(float64(heaters[i]) - float64(heaters[i-1])/2))
// 		// if space > maxInner {
// 		// 	maxInner = space
// 		// }
// 		maxInner = max(maxInner, int(math.Floor((float64(heaters[i])-float64(heaters[i-1]))/2)))
// 	}
// 	flag := false
// 	// 内部还存在空间
// 	for radius < maxInner {
// 		innerFlag := true
// 		// 遍历所有加热范围区间的间隔
// 		for i := 0; i < n-1; i++ {
// 			offset := 0
// 			for j := heaters[i] + radius + 1; j < heaters[i+1]-radius; j++ {
// 				// 如果存在未覆盖到的房子
// 				if count[j] {
// 					offset = min(heaters[i+1]-radius-j, j-(heaters[i]+radius))
// 					break
// 				}
// 			}
// 			if offset > 0 {
// 				radius += offset
// 				innerFlag = false
// 				break
// 			}
// 		}
// 		if innerFlag {
// 			flag = true
// 			break
// 		}
// 	}
// 	if houses[0] < heaters[0]-radius || houses[m-1] > heaters[n-1]+radius {
// 		// radius += max(heaters[0]-radius-houses[0], houses[m-1]-(heaters[n-1]+radius))
// 		flag = false
// 	}
// 	if !flag {
// 		radius = min(max(abs(heaters[0]-houses[0]), abs(houses[m-1]-heaters[0])), max(abs(heaters[n-1]-houses[0]), abs(houses[m-1]-heaters[n-1])))
// 	}
// 	return radius
// }

// 正确做法应该是从房子去找最近的加热器

// 方法一：排序+二分查找
// func findRadius(houses []int, heaters []int) int {
// 	sort.Ints(heaters)
// 	ans := 0
// 	for _, house := range houses {
// 		j := sort.SearchInts(heaters, house+1) // 加1的目的：SearchInts返回的是大于等于value的索引，这里加1就会返回最近的大于value的索引
// 		minDis := math.MaxInt32
// 		if j < len(heaters) {
// 			minDis = heaters[j] - house
// 		}
// 		if j-1 >= 0 && house-heaters[j-1] < minDis {
// 			minDis = house - heaters[j-1]
// 		}
// 		if minDis > ans {
// 			ans = minDis
// 		}
// 	}
// 	return ans
// }

// 方法二：排序+双指针
// 两个指针i, j，分别指向houses 和 heaters
func findRadius(houses []int, heaters []int) (ans int) {
	sort.Ints(houses)
	sort.Ints(heaters)
	j := 0
	for _, house := range houses {
		dis := abs(heaters[j] - house)
		// 对每个当前的house找到离它最近的heater
		for j+1 < len(heaters) && abs(house-heaters[j]) >= abs(house-heaters[j+1]) {
			j++
			dis = abs(house - heaters[j])
		}
		if dis > ans {
			ans = dis
		}
	}
	return
}

func FindRadius(houses []int, heaters []int) int {
	return findRadius(houses, heaters)
}
