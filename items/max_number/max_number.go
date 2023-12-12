package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var num string
	fmt.Scanln(&num)

	f := make([]int, 9)
	for i := 0; i < 9; i++ {
		fmt.Scan(&f[i])
	}

	result := getMaxNumber(num, f)
	fmt.Println(result)
}

func getMaxNumber(num string, f []int) string {
	// 将字符串转换为整数切片
	nums := make([]int, len(num))
	for i := 0; i < len(num); i++ {
		nums[i] = int(num[i] - '0')
	}

	// 找到第一个可以替换的位置
	start := 0
	for start < len(nums) && f[nums[start]-1] <= nums[start] {
		start++
	}

	// 从第一个可以替换的位置开始替换数字
	for i := start; i < len(nums); i++ {
		if f[nums[i]-1] > nums[i] {
			nums[i] = f[nums[i]-1]
		} else {
			// 如果遇到第一个不能替换的位置，后面的数字都不需要替换了
			break
		}
	}

	// 将整数切片转换为字符串
	result := ""
	for i := 0; i < len(nums); i++ {
		result += strconv.Itoa(nums[i])
	}

	return result
}
