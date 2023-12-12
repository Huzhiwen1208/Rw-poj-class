package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// 读取输入
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = line[:len(line)-1]
	n, _ := strconv.Atoi(line)

	heights := make([]int, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		heights[i], _ = strconv.Atoi(line[:len(line)-1])
	}

	result := calculateMinSteps(heights)
	fmt.Println(result)
}

func calculateMinSteps(heights []int) int {
	// 对棋盘高度进行排序
	sort.Ints(heights)

	// 获取最大高度和最小高度
	maxHeight := heights[len(heights)-1]
	minHeight := heights[0]

	// 计算最少需要走的次数
	minSteps := 0
	for i := 1; i < len(heights); i++ {
		if heights[i] != heights[i-1] {
			minSteps++
		}
	}

	// 如果最大高度和最小高度不相等，则需要多走一次
	if maxHeight != minHeight {
		minSteps++
	}

	return minSteps
}
