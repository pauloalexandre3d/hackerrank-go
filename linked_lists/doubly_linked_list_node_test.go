package linked_lists

import (
	"fmt"
	"testing"
)

type DoublyLinkedListNode struct {
	key  int
	val  int
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

type LRUCache struct {
	capacity int
	hashmap  map[int]*DoublyLinkedListNode
	head     *DoublyLinkedListNode
	tail     *DoublyLinkedListNode
}

func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		hashmap:  make(map[int]*DoublyLinkedListNode),
		head:     &DoublyLinkedListNode{key: -1, val: -1},
		tail:     &DoublyLinkedListNode{key: -1, val: -1},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (l *LRUCache) get(key int) int {
	if _, ok := l.hashmap[key]; !ok {
		return -1
	}
	l.remove(l.hashmap[key])
	l.addToTail(l.hashmap[key])
	return l.hashmap[key].val
}

func (l *LRUCache) addToTail(node *DoublyLinkedListNode) {
	prevNode := l.tail.prev
	node.prev = prevNode
	node.next = l.tail
	prevNode.next = node
	l.tail.prev = node
}

func (l *LRUCache) remove(node *DoublyLinkedListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) put(key int, value int) {
	if _, ok := l.hashmap[key]; ok {
		l.remove(l.hashmap[key])
	}
	node := &DoublyLinkedListNode{key: key, val: value}
	l.hashmap[key] = node
	if len(l.hashmap) >= l.capacity {
		delete(l.hashmap, l.head.next.key)
		l.remove(l.head.next)
	}
	l.addToTail(node)
}

func TestAddToTail(t *testing.T) {
	lru := NewLRUCache(3)
	lru.put(1, 100)
	lru.put(2, 250)
	fmt.Println(lru.get(2))
	lru.put(4, 300)
	lru.put(3, 200)
	fmt.Println(lru.get(4))
	fmt.Println(lru.get(1))
}
