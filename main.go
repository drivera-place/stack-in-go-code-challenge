package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	items []interface{}
	rwmu  sync.RWMutex
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() (bool, int) {

	l := 0
	s.rwmu.RLock()
	l = len(s.items)
	defer s.rwmu.RUnlock()

	if l == 0 {
		return true, 0
	}

	return false, l
}

// Push adds an item to the stack
func (s *Stack) Push(item interface{}) {
	s.rwmu.Lock()
	s.items = append(s.items, item)
	s.rwmu.Unlock()
}

// Pop removes the top item from the stack and returns it
func (s *Stack) Pop() (interface{}, bool) {
	item := interface{}(nil)

	if e, _ := s.IsEmpty(); e {
		return item, false
	}

	s.rwmu.Lock()

	last := len(s.items) - 1
	if last >= 0 {
		item = s.items[last]
		s.items = s.items[:last]
	}

	defer s.rwmu.Unlock()

	return item, true
}

// Peek returns the top item without removing it
func (s *Stack) Peek() (interface{}, bool) {
	item := interface{}(nil)
	e, last := s.IsEmpty()

	if e {
		return item, false
	}

	last = last - 1
	if last >= 0 {
		s.rwmu.RLock()
		item = s.items[last]
		s.rwmu.RUnlock()
	}

	return item, true
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
		fmt.Printf("Peeked item: %v \n", item)
	}

	fmt.Printf("stack: %v", myStack)
}
