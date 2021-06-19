package solutions

import "strings"

func IsNumber(s string) bool {
	return isNumber(s)
}

func isNumber(s string) bool {

	//step1: 判断'e'或'E'的存在
	if strings.Contains(s, "e") || strings.Contains(s, "E") {
		//如果是，对前后两部分进行判断
		return
	}
	var res bool = true
	return res
}

func isFloat(sub string) bool {
	var res = false
	return res
}

func isInt(sub string) bool {
	res := true
	return res
}
