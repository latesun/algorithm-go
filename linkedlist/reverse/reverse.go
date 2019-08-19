package reverse

import (
	"fmt"
)

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Print() {
	cursor := l.head.next
	format := ""
	for cursor != nil {
		format += fmt.Sprintf("%+v", cursor.value)
		cursor = cursor.next
		if cursor != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

/*
Single linked list reserve.
Time complexity: O(n)
*/
func (l *LinkedList) Reverse() {
	if l.head == nil || l.head.next == nil || l.head.next.next == nil {
		return
	}

	var preNode *Node = nil
	cursor := l.head.next
	for cursor != nil {
		temp := cursor.next
		cursor.next = preNode
		preNode = cursor
		cursor = temp
	}

	l.head.next = preNode
}

// Determine if the single linked list has a cycle.
func (l *LinkedList) HasCycle() bool {
	if l.head != nil {
		slow := l.head
		fast := l.head
		if fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}
