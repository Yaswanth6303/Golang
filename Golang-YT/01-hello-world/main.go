// In Go, every source file begins with a 'package' declaration.
// 'main' is a special package name in Go.
// Any program with 'package main' is treated as an executable program (not a library).
// The Go compiler looks for a 'main' package with a 'main()' function to start execution.
package main

// The 'import' statement is used to include other packages in our program.
// Here, we are importing the 'fmt' package.
// 'fmt' stands for "format" and provides functions for formatted I/O 
// (Input/Output), such as printing to the console and scanning user input.
import "fmt"

// The 'main' function is the entry point of a Go program.
// Execution starts from here when the program runs.
func main() {
	// 'fmt.Println()' is a function from the 'fmt' package.
	// It prints the given arguments to the standard output (console)
	// and automatically adds a newline at the end.
	fmt.Println("Hello, World!")
}
