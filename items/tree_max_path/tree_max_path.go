package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	// 初始化一个变量来保存最大路径和
	maxSum := math.MinInt32
	// 递归求解最大路径和
	maxPathSumHelper(root, &maxSum)
	return maxSum
}

func maxPathSumHelper(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	// 递归计算左子树的最大路径和，若结果小于0，则取0
	leftSum := max(maxPathSumHelper(node.Left, maxSum), 0)
	// 递归计算右子树的最大路径和，若结果小于0，则取0
	rightSum := max(maxPathSumHelper(node.Right, maxSum), 0)
	// 计算当前节点为根节点的最大路径和
	currSum := node.Val + leftSum + rightSum
	// 更新最大路径和
	*maxSum = max(*maxSum, currSum)
	// 返回以当前节点为根节点的最大路径和（只能选择左子树或右子树的路径）
	return node.Val + max(leftSum, rightSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scanln(&n)

	nodes := make([]*TreeNode, n+1)
	for i := 1; i <= n; i++ {
		var l, r, val int
		fmt.Scanln(&l, &r, &val)
		nodes[i] = &TreeNode{Val: val}
		if l > 0 {
			nodes[i].Left = nodes[l]
		}
		if r > 0 {
			nodes[i].Right = nodes[r]
		}
	}

	result := maxPathSum(nodes[1])
	fmt.Println(result)
}
