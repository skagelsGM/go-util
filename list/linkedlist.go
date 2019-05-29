package list

// Element represents an element in a linked list. This struct contains a Value field for
// identifying the element (int value) and a Next field, which provides a link to the next element
// in the list.
type Element struct {
	Value int
	Next  *Element
}

// CreateElement creates a new Element for the given int n. Element.Next is left unitialized.
// Returns a pointer to the new Element.
func CreateElement(n int) *Element {
	element := new(Element)
	element.Value = n
	return element
}

// LinkedList is a simple linked list of Elements. The list maintains references to the Head and
// Tail elements. Traversal of the list is performed by referencing the current Element's Next field
// until the current element is nil:
//
//  for e:= list.Tail; e != nil; e = e.Next {
//    ...
//  }
//
type LinkedList struct {
	Head, Tail *Element
}

// CreateLinkedList is a factory function for creating a Linked List
func CreateLinkedList() *LinkedList {
	return new(LinkedList)
}

// addElement adds an element to the LinkedList. The element is inserted into the list as the new tail, maintaining a FIFO list (a queue).
func (l *LinkedList) addElement(element *Element) {
	if l.Head == nil {
		l.Head = element
	}

	if l.Tail != nil {
		l.Tail.Next = element
	}

	l.Tail = element
}

// Add adds a new element to the LinkedList for the given int n. The element is inserted into the
// list as the new tail.
func (l *LinkedList) Add(n int) {
	element := CreateElement(n)
	l.addElement(element)
}

// Remove removes the head element from the LinkedList and returns a pointer to the removed Element, maintaing a FIFO data structure (first in, first out). The removed element's reference to the next element in the LinkedList is set to nil.
func (l *LinkedList) Remove() *Element {
	element := l.Head
	l.Head = element.Next
	element.Next = nil
	return element
}
