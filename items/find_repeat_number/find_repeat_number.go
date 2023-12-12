package main

import (
	"fmt"
	"sort"
)

func FindRepeatNumber(arr []int64) int64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return arr[i]
		}
	}

	return -1
}

func main() {
	var (
		n   int
		arr []int64
	)
	fmt.Scanln(&n)

	for i := 0; i < n+1; i++ {
		var e int64
		fmt.Scan(&e)

		arr = append(arr, e)
	}

	// for i := 0; i < n; i++ {
	// 	fmt.Printf("%d ", arr[i])
	// }

	fmt.Println(FindRepeatNumber(arr))
}
