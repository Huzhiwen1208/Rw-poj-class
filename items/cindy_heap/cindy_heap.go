package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 读取输入
	reader := bufio.NewReader(os.Stdin)

	// 读取糖果堆的个数
	line, _ := reader.ReadString('\n')
	line = line[:len(line)-1]
	n, _ := strconv.Atoi(line)

	// 读取每个糖果堆的数量
	line, _ = reader.ReadString('\n')
	line = line[:len(line)-1]
	candies := make([]int, n)
	candiesStr := line
	candiesStrArr := strings.Split(candiesStr, " ")
	for i := 0; i < n; i++ {
		candies[i], _ = strconv.Atoi(candiesStrArr[i])
	}

	result := calculateMaxCandies(candies)
	fmt.Println(result)
}

func calculateMaxCandies(candies []int) int {
	// 统计每个糖果堆的数量
	counts := make(map[int]int)
	for _, candy := range candies {
		counts[candy]++
	}

	// 动态规划数组，dp[i]表示前i个糖果堆中能获得的最大糖果数量
	dp := make([]int, len(counts)+1)
	dp[0] = 0
	if count, ok := counts[1]; ok {
		dp[1] = count
	} else {
		dp[1] = 0
	}

	// 动态规划计算最大糖果数量
	for i := 2; i <= len(counts); i++ {
		if count, ok := counts[i]; ok {
			dp[i] = max(dp[i-1], dp[i-2]+count*i)
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(counts)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
