package complexnumbermultiplication

import (
	"fmt"
	"strconv"
	"strings"
)

type Complex struct {
	R int // 实部
	I int // 虚部
}

func resolve(s string) Complex {
	t := strings.Split(s, "+")
	if len(t) != 2 {
		fmt.Println("invalid string input")
		return Complex{}
	}
	r, _ := strconv.Atoi(t[0])
	i, _ := strconv.Atoi(t[1][:len(t[1])-1])
	return Complex{r, i}
}

func (n Complex) mul(m Complex) Complex {
	r1 := n.R * m.R
	r2 := -1 * (n.I * m.I)
	i1 := n.R * m.I
	i2 := n.I * m.R
	return Complex{r1 + r2, i1 + i2}
}

func (n Complex) toString() string {
	return strconv.Itoa(n.R) + "+" + strconv.Itoa(n.I) + "i"
}

// 模拟
func complexNumberMultiply(num1 string, num2 string) string {
	// 字符串解析
	c1, c2 := resolve(num1), resolve(num2)
	return (c1.mul(c2)).toString()
}

func ComplexNumberMultiply(num1 string, num2 string) string {
	return complexNumberMultiply(num1, num2)
}
