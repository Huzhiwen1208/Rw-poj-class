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

	// 读取辣条的总数
	line, _ := reader.ReadString('\n')
	line = line[:len(line)-1]
	n, _ := strconv.Atoi(line)

	// 读取每根辣条的长度
	line, _ = reader.ReadString('\n')
	line = line[:len(line)-1]
	lengthsStr := strings.Split(line, " ")
	lengths := make([]int, n)
	for i := 0; i < n; i++ {
		lengths[i], _ = strconv.Atoi(lengthsStr[i])
	}

	result := calculateMaxCandies(lengths)
	fmt.Println(result)
}

func calculateMaxCandies(lengths []int) int {
	// 动态规划数组，dp[i]表示以第i根辣条结尾的最长递增子序列的长度
	dp := make([]int, len(lengths))
	dp[0] = 1

	// 计算最长递增子序列的长度
	for i := 1; i < len(lengths); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if lengths[i] > lengths[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// 找到最长递增子序列的最大长度
	maxLength := 0
	for _, length := range dp {
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
