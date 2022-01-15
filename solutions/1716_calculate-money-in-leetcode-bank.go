package solutions

func totalMoney(n int) int {
	k := int(n / 7)
	l := int(n % 7)
	part1 := int(k * (56 + 7*(k-1)) / 2)
	part2 := int(k*l + (l+1)*l/2)
	return part1 + part2
}

func TotalMoney(n int) int {
	return totalMoney(n)
}
