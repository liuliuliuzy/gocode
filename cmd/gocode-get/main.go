/*
根据输入的题号，自动生成待填充的解题模板以及题目描述
*/

package main

import (
	"flag"
	"fmt"
	gocodeget "goleetcode/pkg/cmd/gocode-get"
)

var problemNumber int

func main() {
	// 初始化参数变量
	flag.IntVar(&problemNumber, "n", 1, "problem number")
	// 传递命令行参数
	flag.Parse()
	fmt.Println("you entered: ", problemNumber)
	gocodeget.Execute(problemNumber)
}
