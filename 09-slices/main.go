package main // 'main' package is the entry point of execution in Go.

import (
	"fmt"    // For printing output.
	"slices" // Go 1.21+ standard library package for slice utilities (equality, sorting, etc.).
)

// ----------------------------------------------------------
// SLICES IN GO
// ----------------------------------------------------------
// - Slices are more powerful and flexible than arrays.
// - They are built on top of arrays but have a dynamic size.
// - Arrays in Go are fixed-size, while slices grow and shrink dynamically.
// - Most of the time, slices are preferred over arrays in Go.
// ----------------------------------------------------------

func main() {
	// ----------------------------------------------------------
	// 1. DECLARING AN UNINITIALIZED SLICE
	// ----------------------------------------------------------
	var nums []int // A nil slice (no underlying array allocated yet).

	fmt.Println(nums)        // Output: [] (empty)
	fmt.Println(nums == nil) // true → an uninitialized slice is 'nil'
	fmt.Println(len(nums))   // 0 → no elements

	// ----------------------------------------------------------
	// 2. USING THE make() FUNCTION
	// ----------------------------------------------------------
	// 'make' is used to create slices with a given length and capacity.
	// Syntax: make([]T, length, capacity)
	// If capacity is not given, it is set equal to the length.

	var array = make([]int, 2) // Slice with length = 2, capacity = 2
	fmt.Println(array == nil)  // false → slice is allocated
	fmt.Println(cap(array))    // 2 → capacity is 2

	// Creating a slice with length = 2, capacity = 5
	var arr = make([]int, 2, 5)
	fmt.Println(cap(arr)) // 5

	// ----------------------------------------------------------
	// 3. APPENDING ELEMENTS TO A SLICE
	// ----------------------------------------------------------
	// Append increases the length, and if the capacity is exceeded,
	// Go automatically allocates a new underlying array with doubled capacity.

	arr = append(arr, 1, 2, 3)
	fmt.Println(arr) // Output: [0 0 1 2 3]

	// Appending another slice using '...'
	addElements := []int{8, 9, 10}
	arr = append(arr, addElements...)
	fmt.Println(arr)      // [0 0 1 2 3 8 9 10]
	fmt.Println(cap(arr)) // Capacity increases dynamically

	// Updating elements at specific indexes
	arr[0] = 17
	arr[1] = 21
	fmt.Println(arr) // [17 21 1 2 3 8 9 10]

	// ----------------------------------------------------------
	// 4. EMPTY SLICE INITIALIZATION
	// ----------------------------------------------------------
	a := []int{}        // Empty slice (not nil, but length = 0)
	fmt.Println(a)      // []
	fmt.Println(cap(a)) // 0
	fmt.Println(len(a)) // 0
	a = append(a, addElements...)
	fmt.Println(a) // [8 9 10] → grows dynamically

	// ----------------------------------------------------------
	// 5. COPYING SLICES
	// ----------------------------------------------------------
	// 'copy(destination, source)' copies elements from source into destination.
	// Only copies up to the length of the destination slice.

	var numsArray1 = make([]int, 0, 5)
	// Copy won't work if destination length = 0, so append one element first
	numsArray1 = append(numsArray1, 2)
	var numsArray2 = make([]int, len(numsArray1))

	fmt.Println(numsArray1, numsArray2) // [2] [0]

	copy(numsArray2, numsArray1)        // Copy contents
	fmt.Println(numsArray1, numsArray2) // [2] [2]

	// ----------------------------------------------------------
	// 6. SLICE OPERATOR
	// ----------------------------------------------------------
	numsArray := []int{1, 2, 3, 4, 5}
	fmt.Println(numsArray[0:2]) // Elements from index 0 to 1 → [1 2]
	fmt.Println(numsArray[:2])  // Same as above → [1 2]
	fmt.Println(numsArray[2:])  // From index 2 to end → [3 4 5]

	// ----------------------------------------------------------
	// 7. USING THE "slices" PACKAGE (Go 1.21+)
	// ----------------------------------------------------------
	var nums1 = []int{1, 2}
	var nums2 = []int{1, 2}

	// slices.Equal() checks if two slices are equal element by element.
	fmt.Println(slices.Equal(nums1, nums2)) // true

	// ----------------------------------------------------------
	// 8. MULTIDIMENSIONAL SLICES
	// ----------------------------------------------------------
	var nums2D = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums2D) // [[1 2 3] [4 5 6]]
}
