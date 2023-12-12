package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(node int, dist []int, adj [][]int, visited []bool) (int, int) {
	visited[node] = true
	maxDistNode := node
	maxDist := dist[node]

	for _, neighbor := range adj[node] {
		if visited[neighbor] {
			continue
		}

		dist[neighbor] = dist[node] + 1
		neighborNode, neighborDist := dfs(neighbor, dist, adj, visited)
		if neighborDist > maxDist {
			maxDist = neighborDist
			maxDistNode = neighborNode
		}
	}

	return maxDistNode, maxDist
}

func minTotalDistance(N int, edges [][]int) int {
	adj := make([][]int, N+1)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		adj[x] = append(adj[x], y)
		adj[y] = append(adj[y], x)
	}

	visited := make([]bool, N+1)
	dist := make([]int, N+1)

	// 从根节点1开始DFS遍历，找到离根节点最远的节点
	farthestNode, _ := dfs(1, dist, adj, visited)

	// 从离根节点最远的节点出发，继续DFS遍历，找到离其最远的节点
	visited = make([]bool, N+1)
	dist = make([]int, N+1)
	_, diameter := dfs(farthestNode, dist, adj, visited)

	return diameter
}

func main() {
	// 读取输入
	reader := bufio.NewReader(os.Stdin)
	N := 0
	fmt.Fscan(reader, &N)

	edges := make([][]int, N-1)
	for i := 0; i < N-1; i++ {
		x, y := 0, 0
		fmt.Fscan(reader, &x, &y)
		edges[i] = []int{x, y}
	}

	result := minTotalDistance(N, edges)
	fmt.Println(result)
}
