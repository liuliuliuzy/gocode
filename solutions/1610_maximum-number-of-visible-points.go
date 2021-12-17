package solutions

import (
	"fmt"
	"math"
	"sort"
)

// points中的点都位于第一象限
// 任意旋转，返回能够看到的点数的最大值
// 看了题解，没第一时间想到极坐标，害
func visiblePoints(points [][]int, angle int, location []int) int {
	sameCnt := 0
	polarDegrees := []float64{}
	// 因为是坐标系，所以这里求极坐标中的极角，最方便的函数是acrtan
	for i := range points {
		if points[i][0] == location[0] && points[i][1] == location[1] {
			sameCnt++
			continue
		}
		polarDegrees = append(polarDegrees, math.Atan2(float64(points[i][0]-location[0]), float64(points[i][1]-location[1])))
	}
	sort.Float64s(polarDegrees)
	fmt.Println(polarDegrees)

	n := len(polarDegrees)
	for _, deg := range polarDegrees {
		polarDegrees = append(polarDegrees, deg+2*math.Pi)
	}

	maxCnt := 0
	degree := float64(angle) * math.Pi / 180
	for i, deg := range polarDegrees[:n] {
		j := sort.Search(n*2, func(j int) bool { return polarDegrees[j] > deg+degree })
		if j-i > maxCnt {
			maxCnt = j - i
		}
	}
	return sameCnt + maxCnt
}

func VisiblePoints(points [][]int, angle int, location []int) int {
	return visiblePoints(points, angle, location)
}
