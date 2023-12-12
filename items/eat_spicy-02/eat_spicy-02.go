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

	// 读取每一排辣条的数量
	line, _ := reader.ReadString('\n')
	line = line[:len(line)-1]
	n, _ := strconv.Atoi(line)

	// 读取第一排每根辣条的长度
	line, _ = reader.ReadString('\n')
	line = line[:len(line)-1]
	lengthsStr := strings.Split(line, " ")
	lengths1 := make([]int, n)
	for i := 0; i < n; i++ {
		lengths1[i], _ = strconv.Atoi(lengthsStr[i])
	}

	// 读取第二排每根辣条的长度
	line, _ = reader.ReadString('\n')
	line = line[:len(line)-1]
	lengthsStr = strings.Split(line, " ")
	lengths2 := make([]int, n)
	for i := 0; i < n; i++ {
		lengths2[i], _ = strconv.Atoi(lengthsStr[i])
	}

	result := calculateMaxTotalLength(lengths1, lengths2)
	fmt.Println(result)
}

func calculateMaxTotalLength(lengths1, lengths2 []int) int {
	// 动态规划数组，dp[i][j]表示以第一排第i根辣条和第二排第j根辣条结尾的最大总长度
	dp := make([][]int, len(lengths1)+1)
	for i := 0; i <= len(lengths1); i++ {
		dp[i] = make([]int, 2)
	}

	// 计算最大总长度
	for i := 1; i <= len(lengths1); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+lengths1[i-1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+lengths2[i-1])
	}

	return max(dp[len(lengths1)][0], dp[len(lengths1)][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
