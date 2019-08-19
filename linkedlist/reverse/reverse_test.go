package reverse

import (
	"testing"
)

var l *LinkedList

func init() {
	n5 := &Node{value: 5}
	n4 := &Node{value: 4, next: n5}
	n3 := &Node{value: 3, next: n4}
	n2 := &Node{value: 2, next: n3}
	n1 := &Node{value: 1, next: n2}
	l = &LinkedList{head: &Node{next: n1}}
}

func TestReverse(t *testing.T) {
	l.Print()
	l.Reverse()
	l.Print()
}

func TestHasCycle(t *testing.T) {
	t.Log(l.HasCycle())
	l.head.next.next.next.next.next.next = l.head.next.next.next
	t.Log(l.HasCycle())
}
