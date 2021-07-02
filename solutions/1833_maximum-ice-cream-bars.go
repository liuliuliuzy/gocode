package solutions

import (
	"fmt"
	"sort"
)

//直接贪心算法
//先排序，后购买，一定能够达到最大购买数量
func maxIceCream(costs []int, coins int) int {
	res := 0
	sort.Ints(costs)
	fmt.Println(costs)
	if coins < costs[0] {
		return res
	}
	for _, cost := range costs {
		if cost <= coins {
			res++
			coins -= cost
		}
	}
	return res
}
func MaxIceCream(costs []int, coins int) int {
	return maxIceCream(costs, coins)
}
