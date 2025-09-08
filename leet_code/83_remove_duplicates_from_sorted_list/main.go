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

func deleteDuplicates(head *ListNode) *ListNode {
	if head != nil {
		dummy := &ListNode{Val: head.Val - 1, Next: head}
		curN := head
		preN := dummy

		for curN != nil {
			if curN.Val == preN.Val {
				preN.Next = curN.Next
			} else {
				preN = preN.Next
			}

			curN = curN.Next
		}

		return dummy.Next
	}

	return nil
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
	head.Append(1)
	head.Append(2)
	head.Append(3)
	head.Append(3)

	rs := deleteDuplicates(head)
	PrintList(rs)
}
