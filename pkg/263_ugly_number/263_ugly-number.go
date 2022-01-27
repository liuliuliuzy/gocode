package uglynumber

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	if n < 4 {
		return true
	}
	for n%2 == 0 {
		n = n / 2
	}
	for n%3 == 0 {
		n = n / 3
	}
	for n%5 == 0 {
		n = n / 5
	}
	return n == 1
}

func IsUgly(n int) bool {
	return isUgly(n)
}
