package main // 'main' package is the starting point of execution in Go.

import "fmt" // Importing the 'fmt' package for printing to the console.

// ----------------------------------------------------------
// ARRAYS IN GO
// ----------------------------------------------------------
// An array in Go is a numbered sequence of elements of a specific type
// with a fixed length. The length is part of the array’s type,
// meaning [4]int and [5]int are different types.
//
// - Array elements are stored in contiguous memory.
// - The size of an array cannot be changed once declared.
// - Default (zero) values are automatically assigned for all elements.
// ----------------------------------------------------------

func main() {
	// ----------------------------------------------------------
	// 1. DECLARING AN ARRAY
	// ----------------------------------------------------------
	var nums [4]int // Declares an array of 4 integers. Default value for int = 0.
	// Syntax: var <name> [<length>]<type>

	// Printing the length of the array using len()
	fmt.Println(len(nums)) // Output: 4

	// ----------------------------------------------------------
	// 2. ASSIGNING VALUES TO ARRAY ELEMENTS
	// ----------------------------------------------------------
	nums[0] = 1 // Index starts from 0.
	nums[1] = 2

	// Accessing and printing a single element
	fmt.Println(nums[0]) // Output: 1

	// Printing the entire array
	fmt.Println(nums) // Output: [1 2 0 0] — unassigned elements have default values.

	// ----------------------------------------------------------
	// 3. ARRAYS WITH BOOLEAN TYPE
	// ----------------------------------------------------------
	var vals [4]bool // Boolean array (default value = false).
	vals[2] = true
	fmt.Println(vals) // Output: [false false true false]

	// ----------------------------------------------------------
	// 4. ARRAYS WITH STRING TYPE
	// ----------------------------------------------------------
	var names [3]string // String array (default value = "").
	names[1] = "Yaswanth"
	fmt.Println(names) // Output: ["" "Yaswanth" ""] — empty values for unassigned elements.

	// ----------------------------------------------------------
	// NOTE:
	// Arrays in Go are always filled with the zero-value of their type:
	// - int → 0
	// - bool → false
	// - string → ""
	// - float → 0.0
	// - pointer → nil
	// ----------------------------------------------------------

	// ----------------------------------------------------------
	// 5. DECLARING AND INITIALIZING AN ARRAY AT THE SAME TIME
	// ----------------------------------------------------------
	array := [3]int{1, 2, 3}
	fmt.Println(array) // Output: [1 2 3]

	// ----------------------------------------------------------
	// 6. MULTIDIMENSIONAL ARRAYS
	// ----------------------------------------------------------
	twoDNums := [2][2]int{{3, 4}, {5, 6}} // 2x2 integer array.
	fmt.Println(twoDNums)                 // Output: [[3 4] [5 6]]

	// ----------------------------------------------------------
	// 7. ARRAYS vs SLICES IN GO
	// ----------------------------------------------------------
	// - Arrays have a fixed size, determined at compile time.
	// - Slices are more commonly used because their size can grow or shrink dynamically.
	// - If the size is known and fixed, arrays are more memory-efficient.
	// - Array element access is O(1) (constant time).
}
