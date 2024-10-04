package main

import (
	"testing"
)

func Test_Push(t *testing.T) {
	// Arrange
	stack := Stack{}
	want := 3

	// Act
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	t.Logf("Pushed 3 elementes in Stack: %v", stack.items)
	got := len(stack.items)

	// Assert
	assert(t, got, want)
}

func Test_Pop(t *testing.T) {
	// Arrange
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	t.Logf("Pushed 3 elementes in Stack: %v", stack.items)
	want := 2

	// Act
	item, ok := stack.Pop()
	if ok {
		t.Logf("Removed top element: %v", item)
	}else{
		t.Fatalf("Unexpected error, stack is empty")
	}
	got := len(stack.items)

	// Assert
	assert(t, got, want)
}

func Test_Peek(t *testing.T) {
	// Arrange
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	t.Logf("Pushed 3 elementes in Stack: %v", stack.items)
	want := 3

	// Act
	item, ok := stack.Peek()
	if ok {
		t.Logf("Peeked last element: %v", item)
	}else{
		t.Fatalf("Unexpected error, stack is empty")
	}
	got := len(stack.items)

	// Assert
	assert(t, got, want)
}

func Test_If_Empty(t *testing.T) {
	// Arrange
	stack := Stack{}
	stack.Push(1)
	t.Logf("Pushed 1 elementes in Stack: %v", stack.items)
	want := true

	// Act
	item, ok := stack.Pop()
	if ok {
		t.Logf("Peeked last element: %v", item)
	}else{
		t.Fatalf("Unexpected error, stack is empty")
	}
	_, got := stack.IsEmpty()

	// Assert
	assert(t, got, want)
}

func assert(t *testing.T, got, want interface{}) {
	if got != want {
		t.Errorf("Expected %v but got %v", want, got)
	}

	t.Logf("Expected %v, got %v", want, got)
}
