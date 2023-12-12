package main

import (
	"fmt"
)

func main() {
	var N, R int
	fmt.Scanln(&N, &R)

	coins := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&coins[i])
	}

	health := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&health[i])
	}

	maxCoins := getMaxCoins(N, R, coins, health)
	fmt.Println(maxCoins)
}

func getMaxCoins(N, R int, coins, health []int) int {
	// 创建一个二维数组 dp，用于记录在不同的生命值和不同的野怪数量下的最大金币数
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, R+1)
	}

	for i := 1; i <= N; i++ {
		for j := 0; j <= R; j++ {
			// 当前野怪的生命值小于等于剩余生命值时，可以选择击败该野怪或不击败该野怪
			if health[i-1] <= j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-health[i-1]]+coins[i-1])
			} else {
				// 当前野怪的生命值大于剩余生命值时，不能选择击败该野怪
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[N][R]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
