package main

import (
	"fmt"
	"strings"
	"test/items/utils"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p, q := l1, l2
	var preQ *ListNode
	carry := 0
	for p != nil && q != nil {
		curr := p.Val + q.Val + carry
		carry = curr / 10
		q.Val = curr % 10

		p = p.Next
		preQ = q
		q = q.Next
	}
	if p != nil && q == nil {
		preQ.Next = p
		q = p
	}

	for q != nil {
		curr := q.Val + carry
		carry = curr / 10
		q.Val = curr % 10

		preQ = q
		q = q.Next
	}
	if carry > 0 {
		s := &ListNode{
			Val:  carry,
			Next: nil,
		}

		preQ.Next = s
	}
	return l2
}

func CreateList() *ListNode {
	var p *ListNode
	r := p

	var input string
	fmt.Scanf("%s", &input)
	splits := strings.Split(input, ",")
	for _, item := range splits {
		e := utils.StringToInt(item)

		s := &ListNode{
			Val:  e,
			Next: nil,
		}
		if r != nil {
			r.Next = s
			r = s
		} else {
			p = s
			r = p
		}
	}

	return p
}

func PrintList(p *ListNode) {
	for p != nil {
		fmt.Printf("%v ", p.Val)
		p = p.Next
	}
	fmt.Println()
}

func main() {
	l1 := CreateList()
	l2 := CreateList()

	res := addTwoNumbers(l1, l2)
	PrintList(res)
}
