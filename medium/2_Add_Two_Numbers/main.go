package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintNodes(node *ListNode) {
	current := node
	fmt.Print(current)
	for current.Next != nil {
		current = current.Next
		fmt.Print(current)
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first := ListNode{}
	current := &first
	var one int
	var sum int

	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		one = sum / 10
		current.Val = sum % 10

		if l1 == nil && l2 == nil && one > 0 {
			current.Next = &ListNode{Val: 1, Next: nil}
			break
		} else if l1 != nil || l2 != nil {
			current.Next = &ListNode{}
			current = current.Next
		}
		sum = one
	}

	return &first
}

func reverseList(node *ListNode) string {

	if node == nil {
		return ""
	}
	return strings.Join([]string{reverseList(node.Next), strconv.Itoa(node.Val)}, "")
}

func main() {
	a := ListNode{Val: 9, Next: nil}
	b := ListNode{Val: 5, Next: nil}

	PrintNodes(addTwoNumbers(&a, &b))
}
