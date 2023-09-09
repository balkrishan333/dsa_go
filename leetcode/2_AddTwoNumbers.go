package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, answer := 0, new(ListNode)
	for node := answer; l1 != nil || l2 != nil || carry > 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{carry % 10, nil}
		carry /= 10
	}
	return answer.Next
}

// func main() {
// 	l3 := ListNode{Val: 7}
// 	l4 := ListNode{Val: 9}
// 	l1 := ListNode{2, &l3}
// 	l2 := ListNode{5, &l4}

// 	fmt.Println(addTwoNumbers(&l1, &l2))
// }

func (ln *ListNode) String() string {
	str := ""

	for node := ln; node != nil; node = node.Next {
		str = fmt.Sprint(str, node.Val, ",")
	}
	return str
}

type ListNode struct {
	Val  int
	Next *ListNode
}
