package main // The 'main' package is the entry point of a Go program.
// When we run 'go run', the Go compiler looks for the 'main' package and starts executing from the 'main()' function.

import "fmt" // The 'fmt' package is part of Go's standard library and is used for formatted I/O operations.
// Here we use it for printing output to the console using fmt.Println().

// ---------------------------------------------------------------
// IMPORTANT: Avoid using the same variable name for both global
//            and local variables inside functions. This can lead
//            to confusion and unexpected behavior.
// ---------------------------------------------------------------

// Declaring a global constant named 'number'.
// Constants are variables whose values cannot be changed after initialization.
// Syntax: const <name> <type> = <value>
const number int = 50 // Constants can be declared globally as well, outside any function.

// Declaring a global variable named 'collegeName'.
// 'var' is used for variables whose values can be changed later.
var collegeName = "RVCE"

// The following is an example of invalid short variable declaration at the global level:
// name := "Yaswanth" // ❌ This is not possible because ':=' (short-hand declaration) is only allowed inside functions.

func main() { // The 'main()' function is the starting point of execution.

	// Declaring a local constant named 'name'.
	// Once a constant is assigned a value, it cannot be reassigned.
	const name string = "Yaswanth"

	// Attempting to reassign will result in a compilation error:
	// name = "GoLang" // ❌ Not allowed.

	// Declaring another constant named 'age'.
	const age int = 25

	// Printing the constant values.
	fmt.Println(name) // Output: Yaswanth
	fmt.Println(age)  // Output: 25

	// ---------------------------------------------------------------
	// Multiple constant declaration:
	// Go allows grouping constants together using parentheses '()'.
	// This is useful when you have multiple related constants.
	// ---------------------------------------------------------------
	const (
		port = 5000        // A numeric constant representing a port number.
		host = "localhost" // A string constant representing a host name.
	)
	// These constants are immutable (cannot be changed once assigned).

	// Printing multiple constants in one line:
	fmt.Println(port, host) // Output: 5000 localhost
}
