package main

import (
	"fmt"
	"maps"
)

// ----------------------------------------------------------
// MAPS IN GO
// ----------------------------------------------------------
// A map in Go is an **associative data structure** that stores
// key-value pairs.
//
// - Similar to "Map" in Java or "Dictionary" in Python.
// - Keys must be unique within a map.
// - Values can be duplicated.
// - Zero values are returned if a key does not exist.
// - Maps are reference types (passed by reference).
// ----------------------------------------------------------

func main() {
	// ----------------------------------------------------------
	// 1. CREATING A MAP USING make()
	// ----------------------------------------------------------
	// map[keyType]valueType
	m := make(map[string]string)

	// Adding key-value pairs
	m["name"] = "Golang"
	m["area"] = "Backend"

	// Accessing values by key
	fmt.Println(m["name"], m["area"]) // Output: Golang Backend

	// If a key does not exist → returns the zero value of that type.
	// For strings, the zero value is "" (empty string).
	fmt.Println(m["city"]) // Output: "" (empty string)

	// ----------------------------------------------------------
	// 2. MAP WITH DIFFERENT VALUE TYPES
	// ----------------------------------------------------------
	m1 := make(map[string]int)
	m1["Age"] = 21
	m1["Price"] = 20000
	fmt.Println(m1["Age"], m1["Price"]) // Output: 21 20000
	fmt.Println(m1)                     // Output: map[Age:21 Price:20000]

	// Get the number of key-value pairs using len()
	fmt.Println("Length of the map is:", len(m1)) // Output: 2

	// Deleting a key-value pair
	delete(m1, "Price")
	fmt.Println(m1) // Output: map[Age:21]

	// ----------------------------------------------------------
	// 3. DECLARING AND INITIALIZING A MAP DIRECTLY
	// ----------------------------------------------------------
	m2 := map[string]int{"price": 20000, "age": 21}
	fmt.Println(m2) // Output: map[age:21 price:20000]

	// ----------------------------------------------------------
	// 4. CHECKING IF A KEY EXISTS IN THE MAP
	// ----------------------------------------------------------
	// In Go, the "comma ok" idiom is used to check key existence.
	k, ok := m2["price"]
	fmt.Println(k, ok) // Output: 20000 true

	if ok {
		fmt.Println("Everything is fine")
	} else {
		fmt.Println("Not ok")
	}

	// If the key doesn’t exist, ok = false and k = zero value.

	// ----------------------------------------------------------
	// 5. COMPARING MAPS
	// ----------------------------------------------------------
	// Direct comparison (==) is not allowed for maps, except for nil check.
	// Instead, use the "maps" package (Go 1.21+).
	m3 := map[string]int{"Yaswanth": 100, "Rahul": 200}
	m4 := map[string]int{"Yaswanth": 100, "Rahul": 200}

	fmt.Println(maps.Equal(m3, m4)) // Output: true
}
