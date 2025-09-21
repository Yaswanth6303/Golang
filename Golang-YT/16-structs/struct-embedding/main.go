package main

import (
	"fmt"
	"time"
)

// Defining a struct named 'customer'.
// A struct in Go is a collection of fields (variables) grouped together.
// Here, the 'customer' struct represents customer details.
type customer struct {
	name   string // Name of the customer
	mobile string // Mobile number of the customer
}

// Defining another struct named 'order'.
// This struct represents an order made by a customer.
//
// NOTE: In Go, you can embed one struct inside another.
// This is called **composition**, and it allows us to reuse existing structs
// instead of redefining the same fields again.
//
// Here, 'customer' is embedded in 'order'.
// That means 'order' has all the fields of 'customer' inside it without explicitly redeclaring them.
type order struct {
	id        string    // Order ID
	amount    float32   // Order amount (using float32 for decimal values)
	status    string    // Current status of the order
	createdAt time.Time // Time when the order was created
	customer            // Embedded struct 'customer' (composition)
}

func main() {
	// Creating a new order without providing customer details initially.
	// Notice that we only assign values for 'id', 'amount', and 'status'.
	// Fields that we don't assign explicitly (like 'createdAt' and 'customer')
	// take their **zero values** in Go.
	//
	// Zero values:
	// - For string → ""
	// - For numbers → 0
	// - For structs → empty struct {}
	// - For time.Time → 0001-01-01 00:00:00 +0000 UTC
	newOrder := order{
		id:     "1",
		amount: 40,
		status: "Recieved",
	}

	// Printing the 'newOrder' struct.
	// Since we didn’t fill 'createdAt' and 'customer', they show zero values.
	fmt.Println(newOrder)

	// Printing only the 'customer' part of the 'newOrder'.
	// At this point, it will display: { } (empty struct),
	// because no values were assigned yet.
	fmt.Println(newOrder.customer)
	fmt.Println()

	// Now, let's create a customer separately.
	newCustomer := customer{
		name:   "Jhon",
		mobile: "+91 7493957674",
	}

	// Assigning the new customer to the order.
	// Now 'newOrder' has a customer associated with it.
	newOrder.customer = newCustomer

	// Printing the updated 'newOrder'.
	// This time, the 'customer' field inside 'order' will not be empty,
	// because we assigned it above.
	fmt.Println(newOrder)
}
