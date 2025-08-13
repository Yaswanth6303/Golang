// package main is required for an executable Go program.
// In Go, code is organized into packages. The `main` package is special —
// it tells the compiler this file should be compiled into a standalone executable.
// The Go runtime will start execution by calling the `main()` function from this package.
package main

// fmt is part of Go’s standard library.
// It stands for "format" and provides functions for formatted I/O,
// such as printing to the console or reading from input.
import "fmt"

func main() {

	// ===== Type Inference with var =====
	// In Go, if you don't explicitly specify a variable's type when using `var`,
	// the compiler automatically infers the type based on the assigned value.
	// Example: here the value "John" is a string, so `name` becomes a string.
	var name = "John"
	fmt.Println(name)

	// ===== Short Variable Declaration (:=) =====
	// Go also allows you to declare and initialize variables using `:=` (short declaration operator).
	// This can only be used inside functions, not at the package level.
	// Go will automatically infer the type based on the value.
	age := 20
	fmt.Println(age)

	// ===== Explicit Type Declaration with var =====
	// You can also explicitly specify the type when declaring a variable.
	// This is useful if you want to make the type clear or if the initial value is not assigned immediately.
	var firstName string = "Jane"
	fmt.Println(firstName)

	var marks int = 100
	fmt.Println(marks)

	// ===== Declaring without Initial Value =====
	// You can declare a variable without assigning a value initially.
	// In this case, the variable will have the "zero value" for its type.
	// Later, you can assign a value.
	var programmingLanguage string
	programmingLanguage = "Golang"
	fmt.Println(programmingLanguage)
}
