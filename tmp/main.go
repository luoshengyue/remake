package main

func main() {

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
