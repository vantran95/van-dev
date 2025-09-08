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

func PrintList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	for list2 != nil && list1 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return result.Next
}

func main() {
	headA := &ListNode{Val: 1}
	headA.Append(2)
	headA.Append(4)

	headB := &ListNode{Val: 1}
	headB.Append(3)
	headB.Append(4)

	rs := mergeTwoLists(headA, headB)
	PrintList(rs)
}
