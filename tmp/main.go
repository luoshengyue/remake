package main

import (
	"fmt"
)

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	d := &ListNode{Val: 4}
	e := &ListNode{Val: 5}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	ans := reverseKGroup(a, 2)
	for ans != nil {
		fmt.Printf("%d\t", ans.Val)
		ans = ans.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) (ans *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			slow = head
			for fast != nil && fast.Next != nil {
				fast = fast.Next.Next
				slow = slow.Next
				if fast == slow {
					return fast
				}
			}
		}
	}
	return
}
