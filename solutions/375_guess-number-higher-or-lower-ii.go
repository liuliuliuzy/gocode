package solutions

/*
我们正在玩一个猜数游戏，游戏规则如下：

    我从 1 到 n 之间选择一个数字。
    你来猜我选了哪个数字。
    如果你猜到正确的数字，就会 赢得游戏 。
    如果你猜错了，那么我会告诉你，我选的数字比你的 更大或者更小 ，并且你需要继续猜数。
    每当你猜了数字 x 并且猜错了的时候，你需要支付金额为 x 的现金。如果你花光了钱，就会 输掉游戏 。

给你一个特定的数字 n ，返回能够 确保你获胜 的最小现金数，不管我选择那个数字 。

1 <= n <= 200
*/

//加权的一种二分？试试动态规划吧
// monimize the maximum

/*
又想了一遍，感觉是很经典的动态规划题
TODO:以后准备八股文和leetcode的话一定要记得看这个
*/
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	//逆序遍历挺关键的，从右下到左上角
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			dp[i][j] = j + dp[i][j-1]
			for k := i; k < j; k++ {
				//min(max)
				cost := k + max(dp[i][k-1], dp[k+1][j])
				if cost < dp[i][j] {
					dp[i][j] = cost
				}
			}
		}
	}
	return dp[1][n]
}

func GetMoneyAmount(n int) int {
	return getMoneyAmount(n)
}
