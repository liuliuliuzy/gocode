package gocodeget

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var mainUrl string = "https://leetcode-cn.com/api/problems/all/"

// var graphUrl string = "https://leetcode-cn.com/graphql"

// TODO: 模拟登录过程
// func userLogin(username, passwd string) {

// }

// type problem struct {
// 	id   int
// 	slug string
// }

// 根据题号获取slug
func getBaseInfoByNum(n int) {
	resp, err := http.Get(mainUrl)
	if err != nil {
		fmt.Println("get basic info from leetcode api failed.")
		fmt.Println(err)
		// 终止函数运行
		return
	}
	// resp.
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body)[:300])
}

// 根据slug信息获取题目的详细信息，包括模板代码以及题目描述
// func getDetailInfoBySlug(slug string) {

// }

func Execute(n int) {
	// userLogin()
	getBaseInfoByNum(n)
}

// 参考资料：
// https://www.jianshu.com/p/199881dc86e2?utm_campaign=maleskine
// https://github.com/jaydenwen123/GolangSpider/tree/master/GolangSpider/example/leetcode
