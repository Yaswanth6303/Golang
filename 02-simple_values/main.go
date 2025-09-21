// package main is the starting point of a Go program.
// Every standalone executable Go program must have `package main`.
// The Go compiler looks for the main package to find the entry point (main function).
package main

// fmt is a package in Go's standard library.
// It stands for "format" and provides I/O formatting functions like Println, Printf, etc.
import "fmt"

// The main function is the entry point of the program.
// Execution starts from here in a Go program.
func main() {

	// ===== Integers =====
	// Printing an integer value
	fmt.Println(1)

	// Printing the sum of two integers
	fmt.Println(1 + 1)

	// ===== Strings =====
	// Strings are sequences of characters enclosed in double quotes
	fmt.Println("Hello, World!")

	// ===== Booleans =====
	// Boolean values can be either true or false
	fmt.Println(true)
	fmt.Println(false)

	// ===== Floating Point Numbers =====
	// Printing a floating-point number
	fmt.Println(10.5)

	// Performing floating-point arithmetic
	fmt.Println(150.5 + 10.8) // Addition
	fmt.Println(150.5 - 10.8) // Subtraction
	fmt.Println(150.5 * 10.8) // Multiplication
	fmt.Println(150.5 / 10.8) // Division

	// ===== Complex Numbers =====
	// Complex numbers have a real and imaginary part
	// complex(real, imaginary)
	fmt.Println(complex(1.0, 2.0))

	// ===== Rune =====
	// A rune represents a Unicode code point (character)
	// In single quotes, Go treats it as a rune and prints its ASCII/Unicode value
	fmt.Println('a') // Prints 97 (ASCII value of 'a')
}
