package main

import (
	"bufio"
	"fmt"
	"os"
)

func isTree(n int, edges [][]int) string {
	if n != len(edges)+1 {
		return "NO" // 边的数量不满足树的条件
	}

	// 构建邻接表表示的无向图
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make(map[int]bool)
	parent := make(map[int]int)

	var dfs func(int, int) bool
	dfs = func(node, prev int) bool {
		visited[node] = true
		parent[node] = prev

		for _, neighbor := range graph[node] {
			if neighbor == prev {
				continue // 避免回退到上一个节点
			}

			if visited[neighbor] {
				return false // 存在环，不是树
			}

			if !dfs(neighbor, node) {
				return false
			}
		}

		return true
	}

	if !dfs(1, -1) {
		return "NO" // 存在不连通的节点，不是树
	}

	// 检查是否有孤立节点
	if len(visited) != n {
		return "NO" // 存在孤立节点，不是树
	}

	return "YES"
}

func main() {
	// 读取输入
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	edges := make([][]int, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		edges[i] = []int{a, b}
	}

	result := isTree(n, edges)
	fmt.Println(result)
}
