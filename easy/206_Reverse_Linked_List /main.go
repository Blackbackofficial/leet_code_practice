package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func (*Node) add(val int) {
	current := root
	for {
		if current.Next == nil {
			current.Next = &Node{Val: val, Next: nil}
			return
		}
		current = current.Next
	}
}

func (*Node) reverse() {
	current := root
	var prev *Node
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	root = prev
}

func (*Node) printAll() {
	current := root
	for current.Next != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println()
}

var root = new(Node)

func main() {
	root.add(1)
	root.add(2)
	root.add(3)
	root.add(4)
	root.printAll()
	root.printAll()
	root.reverse()
	root.printAll()
	root.reverse()
	root.printAll()
}
