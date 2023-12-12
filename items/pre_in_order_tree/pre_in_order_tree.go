package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 根据前序遍历结果找到根节点的值
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历结果中找到根节点的索引
	rootIndex := 0
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	// 递归构建左子树和右子树
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append([]int{node.Val}, result...)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return result
}

func main() {
	// 读取输入
	var n int
	fmt.Scanln(&n)

	preorder := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&preorder[i])
	}

	inorder := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&inorder[i])
	}

	// 构建二叉树
	root := buildTree(preorder, inorder)

	// 后序遍历并输出结果
	result := postorderTraversal(root)
	for i := 0; i < len(result); i++ {
		fmt.Printf("%d ", result[i])
	}
	fmt.Println()
}
