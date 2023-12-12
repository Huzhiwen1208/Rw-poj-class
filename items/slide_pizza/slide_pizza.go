package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	pizzas := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&pizzas[i])
	}

	result := getMaxSliceArea(pizzas, m)
	fmt.Printf("%.4f\n", result)
}

func getMaxSliceArea(pizzas []int, m int) float64 {
	// 计算每个披萨的面积
	areas := make([]float64, len(pizzas))
	for i := 0; i < len(pizzas); i++ {
		areas[i] = float64(pizzas[i] * pizzas[i])
	}

	// 将披萨面积按照从大到小排序
	sort.Sort(sort.Reverse(sort.Float64Slice(areas)))

	low := 0.0
	high := areas[0]
	eps := 1e-5

	// 二分查找最大面积
	for high-low > eps {
		mid := (low + high) / 2.0
		count := getCount(areas, mid)
		if count >= m+1 {
			low = mid
		} else {
			high = mid
		}
	}

	return low
}

// 计算给定面积下，披萨可以分给的人数
func getCount(areas []float64, area float64) int {
	count := 0
	for i := 0; i < len(areas); i++ {
		count += int(areas[i] / area)
	}
	return count
}
