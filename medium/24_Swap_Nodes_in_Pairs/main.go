package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	swapPairs(&a)
}

func swapPairs(head *ListNode) *ListNode {
	result := new(ListNode)
	result.Next = head
	current := result
	for current.Next != nil && current.Next.Next != nil {
		a := current.Next // 3
		b := a.Next       // 4
		a.Next = b.Next
		b.Next = a
		current.Next = b
		current = current.Next.Next
	}
	return result.Next
}
