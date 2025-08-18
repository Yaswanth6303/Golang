package main

import (
	"fmt"
)

// ------------------------------------------------------------
// RANGE IN GO
// ------------------------------------------------------------
// The `range` keyword in Go is used to iterate over different
// data structures such as slices, arrays, maps, and strings.
//
// It works like a `for-each` loop in other programming languages.
//
// For each iteration, `range` provides:
// - An index/key
// - The corresponding value
//
// If you don’t need the index/key, you can replace it with `_`.
// ------------------------------------------------------------
func main() {
	// --------------------------------------------------------
	// 1. ITERATING OVER A SLICE USING NORMAL FOR LOOP
	// --------------------------------------------------------
	nums := []int{1, 2, 4, 5, 6}

	// Using classic for loop with len()
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	fmt.Println()

	// --------------------------------------------------------
	// 2. ITERATING OVER A SLICE USING range
	// --------------------------------------------------------
	sum := 0
	// Here, range returns (index, value).
	// If we only need the value, we can ignore the index with `_`.
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum of the array is:", sum)

	// If we want both index and value:
	for i, num := range nums {
		fmt.Println("Value:", num, "| Index:", i)
	}

	// --------------------------------------------------------
	// 3. ITERATING OVER A MAP USING range
	// --------------------------------------------------------
	m := map[string]string{"fname": "Yaswanth", "lname": "Gudivada"}

	// For maps, range returns (key, value).
	for k, v := range m {
		fmt.Println("Key:", k, "| Value:", v)
	}

	fmt.Println()

	// If we only need the keys:
	fmt.Println("Printing Keys only:")
	for k := range m {
		fmt.Println(k)
	}
	fmt.Println()

	// If we only need the values:
	fmt.Println("Printing Values only:")
	for _, v := range m {
		fmt.Println(v)
	}

	// --------------------------------------------------------
	// 4. ITERATING OVER A STRING USING range
	// --------------------------------------------------------
	// When we use range on a string, it returns:
	// - The index (byte position of the rune in the string)
	// - The rune (Unicode code point of the character)
	//
	// ASCII characters (0–255) take 1 byte,
	// but Unicode characters may take 2, 3, or 4 bytes.
	// This is why Go uses "runes" to represent characters.
	for i, c := range "golang" {
		fmt.Println("Index:", i, "| Rune Value:", c)
		// To print the actual character, convert rune → string.
		fmt.Println("Index:", i, "| Character:", string(c))
	}
}
