package main

import (
	"fmt"
	"time"
)

// In Go, there are **no classes** (like in Java, C++).
// Instead, Go provides **structs** to create complex data types.
// Structs allow us to group multiple fields together into a single unit.
// Example use case: In an e-commerce application, we can use a struct to represent an Order.

type order struct { // Structs act like a blueprint, and we can create multiple instances from them.
	id        string
	amount    float32
	status    string
	createdAt time.Time // time.Time is a built-in type that stores timestamps with nanosecond precision.
}

// Unlike other languages, Go does not have constructors.
// But we can create a **constructor-like function** to initialize and return a struct.

// This function creates a new order and returns a pointer to it.
func newOrder(id string, amount float32, status string) *order {
	// Initialize the struct using field:value syntax
	customerOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	// Return a pointer to the struct
	return &customerOrder
}

// In Go, we can associate methods with structs. This allows us to achieve some principles of OOP.
// To do that, we use a **receiver**. The receiver is like `this` in other languages.

// This method changes the status of an order.
// We use a pointer receiver (*order), so the method updates the original struct instead of a copy.
func (o *order) changeStatus(status string) {
	// In Go, structs are automatically dereferenced, so we can write o.status instead of (*o).status.
	o.status = status
}

// This method returns the amount of an order.
// Since we are not modifying the struct, we can use a value receiver (order instead of *order).
func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// Create an order instance (struct literal initialization)
	order1 := order{
		id:     "1",
		amount: 500000,
		status: "Recieved",
	}

	// Assign current time to createdAt field
	order1.createdAt = time.Now()

	fmt.Println("Status of the order: ", order1.status)
	fmt.Println("Order - 1 Struct", order1)

	fmt.Println()
	fmt.Println("After Changing Data")

	// Another order instance
	order2 := order{
		id:        "2",
		amount:    10000,
		status:    "Delivered",
		createdAt: time.Now(),
	}

	// Update status of order1
	order1.status = "Paid"

	fmt.Println("Order - 1 Struct", order1)
	fmt.Println("Order - 2 Struct", order2)
	fmt.Println()

	// Another order instance
	order3 := order{
		id:     "3",
		amount: 100000,
		status: "Recieved",
	}
	fmt.Println("Order - 3 Struct", order3)

	// Call method with pointer receiver → modifies the struct
	order3.changeStatus("Confirmed")
	fmt.Println("Order - 3 (After Changing) Struct", order3)

	// Call method with value receiver → retrieves amount
	fmt.Println("Order 3 Amount: ", order3.getAmount())
	fmt.Println()

	// If we do not assign values to fields, Go uses **zero values**:
	// - 0 for int/float, "" for string, nil for pointer types, and so on.
	order4 := order{
		id: "4", // other fields remain zero values
	}
	fmt.Println("Order - 4 Struct: ", order4)
	fmt.Println()

	// Using constructor-like function newOrder()
	order5 := newOrder("5", 100.49, "Recieved")
	fmt.Println("Order - 5 Struct: ", order5)
	fmt.Println()

	// **Anonymous Structs (Inline Structs)**
	// If we want to use a struct only once and don’t plan to reuse it,
	// we can define it without giving a name.
	languages := struct {
		name   string
		isGood bool
	}{
		"Golang", // first value assigned to name
		true,     // second value assigned to isGood
	}
	fmt.Println("Language Struct: ", languages)
}
