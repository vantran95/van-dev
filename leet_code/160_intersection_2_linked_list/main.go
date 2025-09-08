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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	dummy1 := headA
	dummy2 := headB

	for dummy1 != dummy2 {
		if dummy1 != nil {
			dummy1 = dummy1.Next
		} else {
			dummy1 = headB
		}

		if dummy2 != nil {
			dummy2 = dummy2.Next
		} else {
			dummy2 = headA
		}
	}

	return dummy1
}

func main() {
	headA := &ListNode{Val: 4}
	headA.Append(1)
	headA.Append(8)
	headA.Append(4)

	headB := &ListNode{Val: 5}
	headB.Append(6)
	headB.Append(1)
	headB.Append(8)

	rs := getIntersectionNode(headA, headB)

	PrintList(rs)
}
