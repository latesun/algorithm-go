package linkedlist

import (
	"fmt"
)

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head   *Node
	length uint
}

// Initialize a node.
func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}

// Get next node.
func (n *Node) GetNext() *Node {
	return n.next
}

// Get next node value.
func (n *Node) GetValue() interface{} {
	return n.value
}

// Initialize a linked list.
func NewLinkedList() *LinkedList {
	return &LinkedList{NewNode(0), 0}
}

// Insert a node after a specified node.
func (l *LinkedList) InsertAfter(n *Node, v interface{}) bool {
	if n == nil {
		return false
	}
	newNode := NewNode(v)
	oldNext := n.next
	n.next = newNode
	newNode.next = oldNext
	l.length++
	return true
}

// Insert a node before a specified node.
func (l *LinkedList) InsertBefore(n *Node, v interface{}) bool {
	if n == nil || n == l.head {
		return false
	}

	// Find the prenode of the specified ndoe.
	cursor := l.head.next
	preNode := l.head

	for cursor != nil {
		if cursor == preNode {
			break
		}
		preNode = cursor
		cursor = cursor.next
	}
	if cursor == nil {
		return false
	}

	newNode := NewNode(v)
	preNode.next = newNode
	newNode.next = n
	l.length++
	return true
}

// Insert a node in the head of the linked list.
func (l *LinkedList) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}

// Insert a node in the tail of the linked list.
func (l *LinkedList) InsertToTail(v interface{}) bool {
	// Find the prenode of the tail node.
	cursor := l.head
	for cursor.next != nil {
		cursor = cursor.next
	}

	return l.InsertAfter(cursor, v)
}

// Find a node by index.
func (l *LinkedList) FindByIndex(index uint) *Node {
	if index >= l.length {
		return nil
	}

	cursor := l.head
	for i := uint(0); i < index; i++ {
		cursor = cursor.next
	}
	return cursor
}

// Delete a node.
func (l *LinkedList) DeleteNode(n *Node) bool {
	if n == nil {
		return false
	}

	cursor := l.head.next
	preNode := l.head
	for cursor != nil {
		if n == cursor {
			break
		}
		preNode = cursor
		cursor = cursor.next
	}

	if cursor == nil {
		return false
	}

	preNode.next = n.next
	n = nil
	l.length--

	return true
}

// Print the linked list.
func (l *LinkedList) Print() {
	cursor := l.head.next
	format := ""
	for cursor != nil {
		format += fmt.Sprintf("%+v", cursor.GetValue())
		cursor = cursor.next
		if cursor != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
