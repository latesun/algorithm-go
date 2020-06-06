package stack

import (
	"testing"
)

func TestArrayPush(t *testing.T) {
	a := NewArrayStack()
	t.Log(a.IsEmpty())
	a.Push(1)
	t.Log(a.IsEmpty())

	a.Push(2)
	a.Push(3)
	a.Push("hello")
	t.Logf("%+v\n", a)
}
