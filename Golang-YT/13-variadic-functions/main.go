package main

import "fmt"

// Variadic Function Example in Go

// sum is a variadic function which can take unlimited integer inputs
// Syntax: func functionName(params ...datatype)
// Here `nums ...int` means `nums` will be a slice of integers ([]int).
// Inside the function, we can treat `nums` as a normal slice and loop over it.
func sum(nums ...int) int {
	total := 0
	// Looping over the slice using range
	// `_` is used to ignore the index, `val` gives the actual value
	for _, val := range nums {
		total += val // Accumulate the sum
	}

	// Returning the final total sum of all the integers
	return total
}

func main() {
	// fmt.Println is also a variadic function.
	// It can accept multiple arguments of different data types.
	// That’s why we can pass integers, strings, booleans, runes, etc., all together.
	fmt.Println(1, 2, 3, "Hello", true, 'K')
	// Output: 1 2 3 Hello true 75
	// (Note: 'K' is a rune, so it prints its Unicode value 75)
	fmt.Println()

	// Calling the custom variadic function `sum`
	// Passing multiple integers directly
	sumResult := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Sum of numbers is: ", sumResult)
	// Output: Sum of numbers is: 55
	fmt.Println()

	// Variadic function also accepts a slice,
	// but we cannot pass a slice directly (like sum(nums))
	// because Go will treat it as a single slice argument, not multiple integers.
	// To fix this, we use the `...` operator (called the spread operator in other languages).
	// This expands the slice into individual elements.
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(nums...))
	// Equivalent to calling: sum(1, 2, 3, 4, 5)
	// Output: 15
}
