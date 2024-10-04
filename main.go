package main

import "fmt"

type Stack struct {
	items []interface{}
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() (bool, int) {
	if len(s.items) == 0 {
		return true, 0
	}

	return false, len(s.items)
}

// Push adds an item to the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes the top item from the stack and returns it
func (s *Stack) Pop() (interface{}, bool) {
	b, l := s.IsEmpty()
	l = l - 1

	if !b {
		item := s.items[l]
		s.items = s.items[:l]
		return item, true
	}

	return 0, false

}

// Peek returns the top item without removing it
func (s *Stack) Peek() (interface{}, bool) {
	b, l := s.IsEmpty()

	if !b {
		item := s.items[l-1]
		return item, true
	}

	return 0, false
}

func main() {
	myStack := Stack{}
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)

	item, ok := myStack.Pop()

	if ok {
		fmt.Printf("Removed: %v \n", item)
	}

	item, ok = myStack.Peek()
	if ok {
		fmt.Printf("Peek item: %v \n", item)
	}

	fmt.Printf("stack: %v", myStack)
}
