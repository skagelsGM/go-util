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
	list := new(LinkedList)
	list.Add(1)
	list.Add(2)
	list.Add(3)

	AssertLinkedList(t, list, expected)
}

func TestRemove(t *testing.T) {
	list := new(LinkedList)
	list.Add(1)
	list.Add(2)
	list.Add(3)

	removedElement := list.Remove()
	AssertElement(t, removedElement, 1)
	expected := []int{2, 3}
	AssertLinkedList(t, list, expected)
}

func AssertElement(t *testing.T, e *Element, expectedValue int) {
	if e.Value != expectedValue {
		t.Errorf("Element value is incorrect: %d, expected : %d", e.Value, expectedValue)
	}
}

func AssertLinkedList(t *testing.T, list *LinkedList, expected []int) {
	var i int
	for e := list.Head; e != nil; e = e.Next {
		AssertElement(t, e, expected[i])
		i++
	}
}
