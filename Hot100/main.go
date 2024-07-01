package main

import "fmt"

func printList(node *ListNode) {
	for node != nil {
			fmt.Print(node.Val)
			if node.Next != nil {
					fmt.Print(" -> ")
			}
			node = node.Next
	}
	fmt.Println()
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15},9))

	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	
	result := addTwoNumbers(l1, l2)
	printList(result)
}
