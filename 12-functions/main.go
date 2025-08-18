package main

import (
	"fmt"
)

func main() {
	// --------------------------------------------------------
	// 1. SIMPLE FUNCTION CALL
	// --------------------------------------------------------
	// Call the function `sum` with arguments (3,5).
	// The function returns an integer result which is printed.
	result := sum(3, 5)
	fmt.Println(result)

	// --------------------------------------------------------
	// 2. MULTIPLE RETURN VALUES
	// --------------------------------------------------------
	// The function `getLanguages` returns 4 values:
	// -> three strings and one integer.
	lang1, lang2, lang3, num := getLanguages()

	// Print all return values directly.
	fmt.Println(getLanguages())

	// Print each return value separately.
	fmt.Println(lang1, lang2, lang3, num)

	// If we only want SOME of the return values, we can ignore
	// unwanted ones using `_` (blank identifier).
	// Here we only care about `lang2` and `num`.
	_, lang2, _, num = getLanguages()
	fmt.Println(lang2, num)

	fmt.Println()

	// --------------------------------------------------------
	// 3. FUNCTIONS AS FIRST-CLASS CITIZENS
	// --------------------------------------------------------
	// In Go, functions are "first-class citizens":
	// - We can assign functions to variables
	// - We can pass functions as arguments
	// - We can return functions from other functions
	//
	// Example: Assign an anonymous function (lambda) to a variable.
	fn := func(a int) int {
		// This ignores the input and always returns 21.
		return 21
	}

	// Pass `fn` into another function `processIt`.
	// processIt expects a function that takes an int and returns an int.
	fmt.Println(processIt(fn))

	// --------------------------------------------------------
	// 4. FUNCTION RETURNING ANOTHER FUNCTION
	// --------------------------------------------------------
	// Here, `returnsFunc` returns a function.
	fn1 := returnsFunc()

	// Now `fn1` is itself a function which we can call with an argument.
	fmt.Println(fn1(16)) // Prints 16 because the returned function returns its input.
}

// ------------------------------------------------------------
// FUNCTION DEFINITIONS
// ------------------------------------------------------------

// sum -> Simple function that takes two integers as input
// and returns their sum as an integer.
//
// NOTE: If multiple parameters are of the same type,
// we can declare the type once at the end (like `a, b int`).
func sum(a, b int) int {
	return a + b
}

// getLanguages -> Function returning multiple values.
//
// Go allows functions to return multiple values.
// Here it returns:
//  1. string  -> "GoLang"
//  2. string  -> "JS"
//  3. string  -> "C"
//  4. int     -> 3
func getLanguages() (string, string, string, int) {
	return "GoLang", "JS", "C", 3
}

// processIt -> Higher-order function.
// Takes a function as input (parameter fn).
// The input function must:
//   - accept an int parameter
//   - return an int
//
// Then it calls that function with argument `1` and returns the result.
func processIt(fn func(a int) int) int {
	return fn(1)
}

// returnsFunc -> Function that RETURNS another function.
//
// It returns a function that takes an int as input and
// returns the same int (identity function).
func returnsFunc() func(a int) int {
	return func(a int) int {
		return a
	}
}
