package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 找到链表的中点，将链表拆分为两部分
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil // 断开链表

	// 递归地对两个子链表进行排序
	left := mergeSort(head)
	right := mergeSort(slow)

	// 合并两个已排序的子链表
	return merge(left, right)
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	// 比较两个子链表的节点值，按顺序合并
	for left != nil && right != nil {
		if left.Val < right.Val {
			curr.Next = left
			left = left.Next
		} else {
			curr.Next = right
			right = right.Next
		}
		curr = curr.Next
	}

	// 将剩余节点连接到合并后的链表末尾
	if left != nil {
		curr.Next = left
	} else if right != nil {
		curr.Next = right
	}

	return dummy.Next
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
	mergeSort(head.Next)
}
