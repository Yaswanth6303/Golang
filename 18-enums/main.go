package main

import "fmt"

// ---------------------- ENUMS IN GOLANG ----------------------
// In many programming languages like Java, C#, etc. we have a direct concept
// of "enums" (enumerated types). They are useful when we want to represent
// a fixed set of related constants (like statuses, days of the week, etc.).
//
// Example: An "Order" in an e-commerce app might have different statuses such as:
//   - Received
//   - Confirmed
//   - Prepared
//   - Delivered
//
// If we use raw strings like "received", "confirmed", etc., then we might face issues:
//   1. Inconsistency → Sometimes developers may write them in lowercase, sometimes uppercase,
//      e.g. "Received", "received", "RECEIVED". This leads to bugs.
//   2. Repetition → If we use string literals everywhere, we need to update them in multiple places.
//   3. No compiler check → The compiler cannot restrict values to only valid statuses.
//      Someone might pass "Cancelled" even if we don’t support it.
//
// Go does not have built-in "enum" support, but we can simulate enums
// using **const + iota** and define a **custom type** for better readability.
//
// ---------------------- IMPLEMENTATION ----------------------

// Define a custom type `OrderStatus` which is based on an integer.
// This ensures that we only use this type for order status values.
type OrderStatus int

// Define constants for each status using iota.
// iota is a special Go identifier used inside constant blocks.
// It starts at 0 and increments automatically by 1 for each constant.
//
// So here, values will be assigned as:
//
//	Recieved = 0
//	Confirmed = 1
//	Prepared = 2
//	Delivered = 3
const (
	Recieved  OrderStatus = iota // 0
	Confirmed                    // 1
	Prepared                     // 2
	Delivered                    // 3
)

// Function that takes OrderStatus as a parameter instead of a plain string.
// This enforces type safety → you cannot accidentally pass an unrelated string.
// Only defined constants of type `OrderStatus` can be passed here.
func changeOrderStatus(status OrderStatus) {
	// Prints the numeric value of status (0,1,2,3).
	// By default, printing an enum constant in Go prints its integer value.
	fmt.Println("Change Order Status to:", status)
}

func main() {
	// Calling function with the enum constant.
	// Instead of writing strings like "Received", we now use Recieved (typed constant).
	changeOrderStatus(Recieved)
}
