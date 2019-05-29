package list

import (
	"testing"
)

func TestElement(t *testing.T) {
	e := CreateElement(1)
	AssertElement(t, e, 1)

	e2 := CreateElement(2)
	e.Next = e2
	AssertElement(t, e.Next, 2)
}

func TestLinkedList(t *testing.T) {
	expected := []int{1, 2, 3}
	list := CreateLinkedList()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	AssertLinkedList(t, list, expected)
}

func TestRemove(t *testing.T) {
	list := CreateLinkedList()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	removedElement := list.Remove()
	AssertElement(t, removedElement, 1)
	expected := []int{2, 3}
	AssertLinkedList(t, list, expected)
}
