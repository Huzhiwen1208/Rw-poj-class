package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{}
	result := []int{}
	var prev *TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		// 检查右子树是否已经访问过
		if root.Right == nil || root.Right == prev {
			result = append(result, root.Val)
			stack = stack[:len(stack)-1] // 弹出根节点
			prev = root                  // 记录上一个访问的节点
			root = nil                   // 继续弹出其他节点
		} else {
			root = root.Right // 继续遍历右子树
		}
	}

	return result
}

func Build(root *TreeNode, a [][]int) {
	if root == nil {
		return
	}

	left, right := a[root.Val-1][0], a[root.Val-1][1]
	if left != 0 {
		root.Left = &TreeNode{Val: left, Left: nil, Right: nil}
		Build(root.Left, a)
	}

	if right != 0 {
		root.Right = &TreeNode{Val: right, Left: nil, Right: nil}
		Build(root.Right, a)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	// 建树
	var a [][]int = make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, 2)
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&a[i][0], &a[i][1])
	}

	var root *TreeNode = &TreeNode{Val: 1, Left: nil, Right: nil}
	Build(root, a)

	result := postTraversal(root)
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
}
