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
	params := splitString(line)
	n, _ := strconv.Atoi(params[0])
	m, _ := strconv.Atoi(params[1])

	lengths := make([]float64, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		lengths[i], _ = strconv.ParseFloat(line[:len(line)-1], 64)
	}

	result := calculateMaxLength(lengths, m)
	fmt.Printf("%.2f\n", result)
}

func splitString(str string) []string {
	result := make([]string, 0)
	curr := ""
	for _, ch := range str {
		if ch == ' ' {
			if curr != "" {
				result = append(result, curr)
				curr = ""
			}
		} else {
			curr += string(ch)
		}
	}
	if curr != "" {
		result = append(result, curr)
	}
	return result
}

func calculateMaxLength(lengths []float64, m int) float64 {
	// 对火腿肠长度进行排序
	sort.Float64s(lengths)

	// 计算最小和最大长度
	minLength := lengths[0]
	maxLength := lengths[len(lengths)-1]

	// 二分查找最大长度
	left := minLength
	right := maxLength
	maxLengthCut := 0.0
	for left <= right {
		mid := (left + right) / 2
		if isValidCut(lengths, m, mid) {
			maxLengthCut = mid
			left = mid + 0.01
		} else {
			right = mid - 0.01
		}
	}

	return maxLengthCut
}

func isValidCut(lengths []float64, m int, maxLength float64) bool {
	count := 0
	for _, length := range lengths {
		count += int(length / maxLength)
	}
	return count >= m
}
