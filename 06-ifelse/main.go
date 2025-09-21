package main // The 'main' package is the entry point for execution in Go.

import "fmt" // Importing the 'fmt' package to print output to the console.

func main() {
	// ----------------------------------------------------------
	// IF-ELSE IN GO:
	// Compared to many other programming languages (like C, Java, Python),
	// Go's if-else syntax does not require parentheses around the condition.
	// Example:
	// if condition {
	//     // code
	// } else if anotherCondition {
	//     // code
	// } else {
	//     // code
	// }
	// ----------------------------------------------------------

	age := 18 // Declaring a variable to store age.

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is a teenager")
	} else {
		fmt.Println("Person is a child") // Changed to "child" for correctness.
	}

	fmt.Println() // Empty line for better output separation.

	// ----------------------------------------------------------
	// LOGICAL OPERATORS IN IF-ELSE
	// &&  → Logical AND (both conditions must be true)
	// ||  → Logical OR  (at least one condition must be true)
	// !   → Logical NOT (inverts a boolean value)
	// ----------------------------------------------------------

	var role = "admin"
	var hasPermissions = true

	// Checking if the person is an admin AND has permissions.
	if role == "admin" && hasPermissions {
		fmt.Println("Yes, this person is an admin")
	} else {
		fmt.Println("No, this person is not an admin")
	}

	fmt.Println()

	// ----------------------------------------------------------
	// DECLARING VARIABLES DIRECTLY INSIDE AN IF STATEMENT
	// ----------------------------------------------------------
	// Go allows declaring and initializing a variable directly inside
	// the if condition. This variable will only be accessible inside
	// the if-else block where it is declared.
	if age := 15; age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is a teenager")
	} else {
		fmt.Println("Person is a child")
	}
	// The variable 'age' declared here is separate from the one declared earlier.

	// ----------------------------------------------------------
	// TERNARY OPERATOR IN GO
	// ----------------------------------------------------------
	// Go does NOT support the ternary operator (e.g., condition ? value1 : value2).
	// Instead, you must use the standard if-else structure.
	// Example (other languages):
	//   result = (age >= 18) ? "Adult" : "Minor"
	// In Go, this must be written as:
	//   if age >= 18 {
	//       result = "Adult"
	//   } else {
	//       result = "Minor"
	//   }
	// ----------------------------------------------------------
}
