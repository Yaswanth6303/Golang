package main // The 'main' package is the starting point of any Go program.

import "fmt" // Importing the 'fmt' package to use functions like Println for output.

// ------------------------------------------------------------
// NOTE ABOUT LOOPS IN GOLANG:
// Go does not have a 'while' loop or 'do-while' loop like some
// other languages (C, Java, Python). Instead, Go uses only
// the 'for' loop for all looping purposes.
//
// 'for' in Go can work in three main ways:
// 1. Like a 'while' loop.
// 2. As a classic 'for' loop (with init, condition, increment).
// 3. As an infinite loop.
// ------------------------------------------------------------

func main() {

	// ------------------------------
	// 1. Using 'for' like a while loop
	// ------------------------------
	i := 1 // Initialize a counter variable.

	// This loop will run while 'i <= 5'.
	for i <= 5 {
		fmt.Println(i)
		i = i + 1 // Increment 'i' by 1 in each iteration.
	}

	fmt.Println() // Print an empty line for separation.

	// ------------------------------
	// 2. Infinite loop in Go
	// ------------------------------
	/*
	   Syntax for an infinite loop in Go:
	   for {
	       // Code here will run forever unless you break out manually
	   }
	*/
	// Example (commented out to prevent infinite execution):
	/*
		for {
			fmt.Println("This will run forever")
		}
	*/

	// ------------------------------
	// 3. Classic 'for' loop
	// ------------------------------
	// Syntax: for initialization; condition; post-statement { ... }
	for i := 0; i <= 5; i++ {

		// 'continue' skips the current iteration and moves to the next one.
		if i == 2 {
			continue
		}

		// 'break' exits the loop entirely.
		if i == 4 {
			break
		}

		fmt.Println(i)
	}

	fmt.Println() // Separation line.

	// ------------------------------
	// 4. New 'range' feature in Go 1.22
	// ------------------------------
	// Before Go 1.22, 'range' was mainly used with slices, arrays, maps, and strings.
	// From Go 1.22 onward, you can 'range' over an integer value directly.
	// The loop will run from 0 to k-1.
	k := 3
	for i := range k {
		fmt.Println(i) // Output: 0, 1, 2
	}
}
