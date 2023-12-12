package main

import (
	"bufio"
	"fmt"
	"os"
)

func hasPath(n, m, s, t int, edges [][]int) int {
	graph := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
	}

	visited := make([]bool, n+1)

	var dfs func(int) bool
	dfs = func(node int) bool {
		if node == t {
			return true
		}

		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				if dfs(neighbor) {
					return true
				}
			}
		}

		return false
	}

	if dfs(s) {
		return 1
	}
	return 0
}

func main() {
	// 读取输入
	reader := bufio.NewReader(os.Stdin)
	var n, m, s, t int
	fmt.Fscan(reader, &n, &m, &s, &t)

	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}

	result := hasPath(n, m, s, t, edges)
	fmt.Println(result)
}
