package main

import "fmt"

type Person struct {
	name string
	age  int
}

// Simple generic example, [T any] is the type parameters, we pass s with type T as the args
func Print[T any](s T) {
	// We then print out the T generic
	fmt.Println(s)
}

func main() {
	bob := Person{name: "bob", age: 25}
	Print("Jammy")
	Print(1)
	Print(false)
	Print(bob)
}
