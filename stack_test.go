package main

import (
	"sync"
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
	assertEqual(t, got, want)
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
	} else {
		t.Fatalf("Unexpected error, stack is empty")
	}
	got := len(stack.items)

	// Assert
	assertEqual(t, got, want)
}

func Test_Pop_Until_Empty(t *testing.T) {
	// Arrange
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	t.Logf("Pushed 3 elementes in Stack: %v", stack.items)
	want := 0
	l := len(stack.items)

	// Act
	for i := 0; i < l; i++ {
		item, ok := stack.Pop()
		if ok {
			t.Logf("Removed top element: %v", item)
		} else {
			t.Fatalf("Unexpected error, stack is empty")
		}
	}
	got := len(stack.items)

	// Assert
	assertEqual(t, got, want)
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
	} else {
		t.Fatalf("Unexpected error, stack is empty")
	}
	got := len(stack.items)

	// Assert
	assertEqual(t, got, want)
}

func Test_Peek_When_Empty(t *testing.T) {
	// Arrange
	stack := Stack{}
	t.Logf("Stack is empty: %v", stack.items)
	want := 0

	// Act
	item, ok := stack.Peek()
	if !ok {
		t.Logf("Unable to peek, stack is empty: %v", item)
	}
	got := len(stack.items)

	// Assert
	assertEqual(t, got, want)
}

func Test_If_Empty(t *testing.T) {
	// Arrange
	stack := Stack{}
	stack.Push(1)
	t.Logf("Pushed 1 elementes in Stack: %v", stack.items)
	want := 0

	// Act
	item, ok := stack.Pop()
	if ok {
		t.Logf("Peeked last element: %v", item)
	} else {
		t.Fatalf("Unexpected error, stack is empty")
	}
	_, got := stack.IsEmpty()

	// Assert
	assertEqual(t, got, want)
}

func Test_Concurrently_Push(t *testing.T) {
	// Arrange
	stack := Stack{}
	want := 12
	wg := sync.WaitGroup{}

	// Act
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func() {
			stack.Push(i)
			wg.Done()
		}()
	}
	wg.Wait()

	t.Logf("Pushed 12 elementes in Stack: %v", stack.items)
	got := len(stack.items)

	// Assert
	assertEqual(t, got, want)
}

func Test_Pop_Concurrently(t *testing.T) {
	// Arrange
	stack := Stack{}
	want := 0
	wg := sync.WaitGroup{}
	max := 12
	for i := 0; i < max; i++ {
		wg.Add(1)
		go func() {
			stack.Push(i)
			wg.Done()
		}()
	}
	wg.Wait()

	// Act
	for j := 0; j < max; j++ {
		wg.Add(1)
		go func() {
			stack.Pop()
			wg.Done()
		}()
	}
	wg.Wait()

	t.Logf("Popped 12 elementes, in stack: %v", stack.items)
	got := len(stack.items)

	// Assert
	assertEqual(t, got, want)
}

func assertEqual(t *testing.T, got, want interface{}) {
	if got != want {
		t.Errorf("Expected %v but got %v", want, got)
		return
	}

	t.Logf("Expected %v, got %v", want, got)
}
