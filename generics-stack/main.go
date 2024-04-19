package main

import "fmt"

// A stack is an array of items, note the generic [T any] type parameter
type Stack[T any] struct {
	// a slice of any items
	items []T
}

type Person struct {
	name string
}

// Adds any item (T) to Stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Removes last item and returns it, not the return type [T]
func (s *Stack[T]) Pop() T {
	// panic if there is no item on stack
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	// get last item from stack
	item := s.items[len(s.items)-1]
	// shorten stack via removing last item
	s.items = s.items[:len(s.items)-1]
	// return generic item
	return item
}

func main() {
	// Initalize our stacks
	// we can limit the input type via [] as we can see below
	var IntStack Stack[int]
	var StringStack Stack[string]
	var PersonStack Stack[Person]

	// Add value to IntStack
	IntStack.Push(1)
	IntStack.Push(2)

	// Add value to StringStack
	StringStack.Push("Alberto")

	// Add value to PersonStack
	PersonStack.Push(Person{name: "bob"})

	// Pop Values
	fmt.Println(IntStack.Pop())
	fmt.Println(StringStack.Pop())
	fmt.Println(PersonStack.Pop())
}
