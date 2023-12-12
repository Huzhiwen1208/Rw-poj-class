package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// rotate a linklist by k
func rotateLinklist(head *ListNode, k int) {
	// get the length of linklist
	length := 0
	for p := head; p != nil; p = p.Next {
		length++
	}

	// get the real rotate count
	k = k % length

	// rotate linklist
	p := head
	for i := 0; i < length-k-1; i++ {
		p = p.Next
	}
	newHead := p.Next
	p.Next = nil
	p = newHead
	for p.Next != nil {
		p = p.Next
	}
	p.Next = head
}

func main() {
	// 输入N, k
	var N, k int
	fmt.Scan(&N, &k)
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
	rotateLinklist(head.Next, k)
}
