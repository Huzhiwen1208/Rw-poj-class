package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 读取输入
	reader := bufio.NewReader(os.Stdin)

	// 读取同学人数
	line, _ := reader.ReadString('\n')
	line = line[:len(line)-1]
	n, _ := strconv.Atoi(line)

	// 读取每个同学左侧和右侧需要空出的座位数
	seats := make([][]int, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = line[:len(line)-1]
		seatStr := []rune(line)
		l, _ := strconv.Atoi(string(seatStr[:len(seatStr)/2]))
		r, _ := strconv.Atoi(string(seatStr[len(seatStr)/2+1:]))
		seats[i] = []int{l, r}
	}

	result := calculateMinSeats(seats)
	fmt.Println(result)
}

func calculateMinSeats(seats [][]int) int {
	// 计算左侧和右侧需要空出的座位数的最大值
	maxLeft := 0
	maxRight := 0
	for _, seat := range seats {
		if seat[0] > maxLeft {
			maxLeft = seat[0]
		}
		if seat[1] > maxRight {
			maxRight = seat[1]
		}
	}

	// 返回左侧和右侧需要空出的座位数的最大值的两倍加1
	return maxLeft*2 + 1
}
