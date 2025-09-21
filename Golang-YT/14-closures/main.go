package main

import "fmt"

// counter is a function that does not take any parameters
// but it returns another function as its output.
// The returned function itself has no parameters, and it returns an integer.
//
// Concept: This demonstrates a "closure" in Go.
// A closure is a function that captures variables from its surrounding scope.
func counter() func() int {
	// Define a variable 'count' inside counter function.
	// This variable is scoped to the counter function but will still
	// be accessible to the inner function (closure).
	var count int = 0

	// Return an anonymous function.
	// Every time this function is called:
	//   1. It increments 'count' by 1.
	//   2. Returns the updated value of 'count'.
	// NOTE: The inner function "remembers" the value of 'count' even
	// after counter() has finished executing. This is the essence of a closure.
	return func() int {
		count += 1
		return count
	}
}

func main() {
	// Call counter() which returns a function.
	// Assign that returned function to the variable 'increment'.
	// Now 'increment' itself is a function we can call multiple times.
	increment := counter()

	// First call: 'count' starts at 0, increments to 1, and returns 1.
	fmt.Println(increment()) // Output: 1

	// Second call: 'count' was already 1, increments to 2, and returns 2.
	fmt.Println(increment()) // Output: 2

	// Notice that 'increment' maintains its own state of 'count'
	// across multiple calls, thanks to closure.
}
