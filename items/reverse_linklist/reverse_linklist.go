package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 翻转m到n的链表
func reverseBetween(head *ListNode, m int, n int) {
	// 1. 找到m的前一个节点
	pre := &ListNode{Next: head}
	for i := 1; i < m; i++ {
		pre = pre.Next
	}

	// 2. 翻转m到n的链表
	cur := pre.Next
	for i := m; i <= n; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
}

func main() {
	// 输入N, m, n
	var N, m, n int
	fmt.Scan(&N, &m, &n)
	// 尾插法构造带头结点的链表
	head := &ListNode{}
	tail := head
	for i := 0; i < N; i++ {
		var val int
		fmt.Scan(&val)
		node := &ListNode{Val: val}
		tail.Next = node
		tail = node
	}
	// 翻转m到n的链表
	reverseBetween(head.Next, m, n)
}
