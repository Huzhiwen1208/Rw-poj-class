package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// 读取输入
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = line[:len(line)-1]
	params := strings.Split(line, " ")
	strconv.Atoi(params[0])
	m, _ := strconv.Atoi(params[1])

	line, _ = reader.ReadString('\n')
	line = line[:len(line)-1]

	result := calculateMaxScore(line, m)
	fmt.Println(result)
}

func calculateMaxScore(cards string, m int) int {
	// 统计每张卡牌的出现次数
	counts := make(map[byte]int)
	for i := 0; i < len(cards); i++ {
		counts[cards[i]]++
	}

	// 将卡牌的出现次数放入数组中
	frequencies := make([]int, 0)
	for _, count := range counts {
		frequencies = append(frequencies, count)
	}

	// 对卡牌的出现次数进行排序
	sort.Ints(frequencies)

	// 计算最大分值总和
	maxScore := 0
	for i := len(frequencies) - 1; i >= 0; i-- {
		if m > 0 {
			maxScore += frequencies[i]
			m--
		} else {
			break
		}
	}

	return maxScore
}
