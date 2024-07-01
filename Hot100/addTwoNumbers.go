package main

type ListNode struct {
	Val int
	Next *ListNode
}

func getValueAndMoveNext(node **ListNode) int {
	if *node == nil {
		return 0
	} else {
		val := (*node).Val
		*node = (*node).Next
		return val
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	current := head
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		n1 := getValueAndMoveNext(&l1)
		n2 := getValueAndMoveNext(&l2)
		
		sum := (n1 + n2 + carry)
		current.Next = &ListNode{Val:  sum % 10}
		current = current.Next
		carry = sum / 10
	}
	return head.Next
}