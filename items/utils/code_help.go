package utils

import "strconv"

func StringToInt(a string) int {
	i, _ := strconv.ParseInt(a, 10, 32)
	return int(i)
}
