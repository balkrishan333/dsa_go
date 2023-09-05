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

func main() {
	l1 := ListNode{Val: 2}
	l2 := ListNode{Val: 5}

	fmt.Println(addTwoNumbers(&l1, &l2).Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
