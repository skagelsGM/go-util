package list

import (
	"testing"
)

// AssertElement asserts an element in a linked list contains the expectedValue
func AssertElement(t *testing.T, e *Element, expectedValue int) {
	if e.Value != expectedValue {
		t.Errorf("Element value is incorrect: %d, expected : %d", e.Value, expectedValue)
	}
}

// AssertLinkedList asserts that a linked list contains the values provided in the expected array of integers
func AssertLinkedList(t *testing.T, list *LinkedList, expected []int) {
	var i int
	for e := list.Head; e != nil; e = e.Next {
		AssertElement(t, e, expected[i])
		i++
	}
}
