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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	return nil
}

func reverseInt(n int) int {
	reversed := 0
	for n != 0 {
		remainder := n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}
	return reversed
}

func main() {
	headA := &ListNode{Val: 2}
	headA.Append(4)
	headA.Append(3)

	headB := &ListNode{Val: 5}
	headB.Append(6)
	headB.Append(4)

	rs := addTwoNumbers(headA, headB)

	PrintList(rs)
}
