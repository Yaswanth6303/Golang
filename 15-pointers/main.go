package main

import "fmt"

// By value
// func changeNum(num int) {
//     num = 5
//     fmt.Println("In changeNum", num)
// }
//
// func main() {
//     num := 1
//     // To change this value of num, we need to pass it by reference (pointer), not by value.
//     // By default, Go always passes arguments by value, so the function gets its own copy.
//     changeNum(num)
//     fmt.Println("After changeNum in main function", num)
// }

// By reference
func changeNum(num *int) {
	*num = 5 // We received the address of num, so we dereference it (*) and update the actual value
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1

	fmt.Println("Memory address", &num)         // Prints the memory address of num
	changeNum(&num)                             // Passes the memory address of num (reference) to the function
	fmt.Println("After changeNum in main", num) // Shows that the actual value of num has changed
}
