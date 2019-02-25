package list

// Element represents a element in a linked list. This struct contains a Value field for identifying
// the element (int value) and a Next field, which provides a link to the next element in the list.
type Element struct {
	Value int
	Next  *Element
}

// CreateElement creates a new Element for the given element n. Element.Next is left unitialized.
// Returns a pointer to the new Element.
func CreateElement(n int) *Element {
	element := new(Element)
	element.Value = n
	return element
}

// LinkedList is a simple linked list of Elements.
// The list maintains a reference to the Head element. Traversal of the list is
// performed by referencing the current Element's Next field until Next is nil (the Zero-Value)
type LinkedList struct {
	Head, Tail *Element
}

// addElement adds a element to the LinkedList. The element is inserted into the list as the new tail, maintaining a FIFO list (a queue).
func (l *LinkedList) addElement(element *Element) {
	if l.Head == nil {
		l.Head = element
	}

	if l.Tail != nil {
		l.Tail.Next = element
	}

	l.Tail = element
}

// Add adds a new element to the LinkedList for the given vertex v. The element is inserted into the
// list as the new tail.
func (l *LinkedList) Add(v int) {
	element := CreateElement(v)
	l.addElement(element)
}

// Remove removes the head element from the LinkedList and returns the element
func (l *LinkedList) Remove() *Element {
	element := l.Head
	l.Head = element.Next
	return element
}
