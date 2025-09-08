package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Append(val int) *ListNode {
	newNode := &ListNode{Val: val}
	if l == nil {
		return newNode // Trường hợp danh sách rỗng
	}

	current := l
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return l
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	curN := head
	preN := dummy
	for curN != nil {
		if curN.Val == val {
			preN.Next = curN.Next
		} else {
			preN = preN.Next
		}

		curN = curN.Next
	}

	PrintList(dummy.Next)

	return dummy.Next
}

func PrintList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	head := &ListNode{Val: 1}
	head.Append(4)
	head.Append(6)
	head.Append(10)

	removeElements(head, 6)

	//PrintList(head)
}
