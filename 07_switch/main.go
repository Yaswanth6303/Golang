package main // The 'main' package is the starting point of execution in Go.

import (
	"fmt"  // The 'fmt' package is used for printing output to the console.
	"time" // The 'time' package provides functionalities for working with dates and times.
)

func main() {
	// ----------------------------------------------------------
	// 1. SIMPLE SWITCH STATEMENT
	// ----------------------------------------------------------
	// In Go, 'switch' is used for conditional branching, similar to other languages.
	// Unlike C, C++, or Java, we do NOT need to explicitly write 'break' after each case.
	// Once a matching case is found, Go automatically exits the switch block.
	// (If you want fall-through behavior, you must explicitly use 'fallthrough'.)

	i := 5 // Example integer value.
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Third")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Another Number") // Runs when no case matches.
	}

	fmt.Println() // Empty line for better output separation.

	// ----------------------------------------------------------
	// 2. MULTIPLE CONDITION SWITCH
	// ----------------------------------------------------------
	// A single case in Go's switch can check for multiple matching values.
	// Example: Using the 'time' package to check the current day of the week.

	switch time.Now().Weekday() { // time.Now().Weekday() returns the current day (Monday, Tuesday, etc.).
	case time.Saturday, time.Sunday:
		fmt.Println("Hooray, It's Weekend") // Matches if today is Saturday or Sunday.
	default:
		fmt.Println("Sad, It's Weekday") // Runs if it's a weekday.
	}

	fmt.Println()

	// ----------------------------------------------------------
	// 3. TYPE SWITCH
	// ----------------------------------------------------------
	// A type switch is used when you have a value of 'interface{}' type
	// and want to find out its actual underlying data type.
	// 'interface{}' in Go is a generic type that can hold values of any type.

	// Example: Assigning a function to a variable.
	// In Go, functions are first-class citizens:
	// - You can store them in variables.
	// - You can pass them as arguments to other functions.
	// - You can return them from functions.

	whoAmI := func(i interface{}) {
		// 'i.(type)' is used in a type switch to determine the actual type of 'i'.
		switch variableType := i.(type) {
		case int:
			fmt.Println("The given value is an integer")
		case string:
			fmt.Println("The given value is a string")
		case bool:
			fmt.Println("The given value is a boolean")
		default:
			// If none of the above cases match, this will run.
			// 'variableType' here actually contains the dynamic type name.
			fmt.Println("Other Datatype", variableType)
		}
	}

	// Testing the function with different types.
	whoAmI("Yaswanth") // string case.
	whoAmI('a')        // rune (alias for int32) — prints ASCII code value in 'default'.
}
