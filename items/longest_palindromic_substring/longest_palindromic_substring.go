package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	var dp [][]bool = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if j-i+1 > 2 {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j]
			}
		}
	}

	l, r := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] && j-i+1 > r-l {
				r = j + 1
				l = i
			}
		}
	}

	return s[l:r]
}

func main() {
	var s string
	fmt.Scanln(&s)

	res := longestPalindrome(s)
	fmt.Println(res)
}
