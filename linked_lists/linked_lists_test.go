package linked_lists

import (
	"fmt"
	"testing"
)

type ListNode struct {
	val  int
	prev *ListNode
	next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("%+v, -> %+v", l.val, l.next)
}

func TestLinkedListReversal(t *testing.T) {
	ln4 := &ListNode{val: 4}
	ln3 := &ListNode{val: 3, next: ln4}
	ln2 := &ListNode{val: 2, next: ln3}
	head := &ListNode{val: 1, next: ln2}

	nr := linkedListReversal(head)
	fmt.Printf("Linked list reversal: %+v", nr)
}

func linkedListReversal(head *ListNode) *ListNode {
	var currNode, prevNode *ListNode = head, nil
	for currNode != nil {
		nextNode := currNode.next
		currNode.next = prevNode
		prevNode = currNode
		currNode = nextNode
	}
	return prevNode
}

func TestRemoveKthLastNode(t *testing.T) {
	ln3 := &ListNode{val: 3}
	ln7 := &ListNode{val: 7, next: ln3}
	ln4 := &ListNode{val: 4, next: ln7}
	ln2 := &ListNode{val: 2, next: ln4}
	head := &ListNode{val: 1, next: ln2}

	nr := removeKthLastNode(head, 2)
	fmt.Printf("Remove kth last node: %+v", nr)
}

func removeKthLastNode(head *ListNode, k int) *ListNode {
	dummy := &ListNode{val: -1, next: head}
	trailer, leader := dummy, dummy

	for range k {
		leader = leader.next
		if leader == nil {
			return head
		}
	}

	for leader.next != nil {
		trailer = trailer.next
		leader = leader.next
	}

	trailer.next = trailer.next.next
	return dummy.next
}

func TestLinkedListIntersection(t *testing.T) {
	ln2 := &ListNode{val: 2}
	ln7 := &ListNode{val: 7, next: ln2}
	ln8 := &ListNode{val: 8, next: ln7}

	ln4 := &ListNode{val: 4, next: ln8}
	ln6 := &ListNode{val: 6, next: ln4}
	ln3 := &ListNode{val: 3, next: ln4}
	ln1 := &ListNode{val: 1, next: ln3}

	intersection := linkedListIntersection(ln1, ln6)
	fmt.Printf("Linked list intersection: %+v", intersection)
}

func linkedListIntersection(headA *ListNode, headB *ListNode) *ListNode {
	ptrA, ptrB := headA, headB
	for ptrA != ptrB {
		if ptrA != nil {
			ptrA = ptrA.next
		} else {
			ptrA = headB
		}
		if ptrB != nil {
			ptrB = ptrB.next
		} else {
			ptrB = headA
		}
	}
	return ptrA.next
}
