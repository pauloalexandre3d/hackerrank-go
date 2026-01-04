package fast_and_slow_pointer

import (
	"fmt"
	"testing"
)

type ListNode struct {
	val  int
	prev *ListNode
	next *ListNode
}

func TestLinkedListLoop(t *testing.T) {

	head := &ListNode{val: 1}
	node2 := &ListNode{val: 2}
	node3 := &ListNode{val: 3}
	node4 := &ListNode{val: 4}
	node5 := &ListNode{val: 5}

	head.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node2

	isLoop := linkedListLoop(head)
	fmt.Println(isLoop)
}

func linkedListLoop(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			fmt.Println(slow)
			return true
		}
	}
	return false
}
