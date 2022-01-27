package solutions

func maximumTime(time string) string {
	res := ""
	for index, char := range time {
		if index == 2 {
			res += ":"
			continue
		}
		if char != '?' {
			res += time[index : index+1]
		} else {
			switch index {
			case 0:
				if time[1] == '?' || time[1] < '4' {
					res += "2"
				} else {
					res += "1"
				}
			case 1:
				if time[0] == '?' {
					res += "3"
				} else if time[0] < '1' {
					res += "9"
				} else {
					res += "3"
				}
			case 3:
				res += "5"
			case 4:
				res += "9"
			}
		}
	}
	return res
}

func MaximumTime(time string) string {
	return maximumTime(time)
}
