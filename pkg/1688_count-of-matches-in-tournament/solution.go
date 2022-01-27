package countofmatchesintournament

func numberOfMatches(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 1 {
		return n/2 + numberOfMatches((n+1)/2)
	} else {
		return n/2 + numberOfMatches(n/2)
	}
}

func NumberOfMatches(n int) int {
	return numberOfMatches(n)
}
