package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)

	res := LongestUnRepeatSubstring(s)
	fmt.Printf("%s \n", res)
}

func LongestUnRepeatSubstring(s string) string {
	record := map[rune]int{}
	currentLeft, currentRight, resultLeft, resultRight := 0, 0, 0, 0
	for i, item := range []rune(s) {
		if v, ok := record[item]; !ok || v < currentLeft {
			record[item] = i
		} else {
			idx := record[item]
			currentLeft = idx + 1
		}

		record[item] = i
		currentRight++

		if resultRight-resultLeft < currentRight-currentLeft {
			resultRight = currentRight
			resultLeft = currentLeft
		}
	}

	return s[resultLeft:resultRight]
}
