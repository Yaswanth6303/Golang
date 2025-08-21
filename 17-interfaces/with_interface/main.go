package main

import "fmt"

//
// ------------------------- INTERFACES IN GO --------------------------
// In Go, an interface defines a set of method signatures (behaviors).
// Any type (struct, custom type, etc.) that provides implementations for
// those methods AUTOMATICALLY "implements" the interface.
// There is no explicit "implements" keyword in Go (unlike Java or C#).
//
// By convention in Go, interface names usually end with "er"
// (e.g., Reader, Writer, Formatter, PaymentGateway).
// ---------------------------------------------------------------------
//

// PaymentGateway represents any payment processing system.
// It declares a single method `Pay` that accepts an amount.
type PaymentGateway interface {
	Pay(amount float32)
}

//
// ------------------- CONCRETE IMPLEMENTATIONS ------------------------
// We now create specific payment systems (Stripe, Razorpay) that satisfy
// the PaymentGateway interface by providing a `Pay` method.
// ---------------------------------------------------------------------
//

// Stripe struct represents a payment system (Stripe).
// This struct does not have any fields; it just serves as a type.
type Stripe struct{}

// Pay method for Stripe.
// Since the method signature matches the interface requirement
// (i.e., `Pay(amount float32)`), Stripe automatically implements
// PaymentGateway.
func (s Stripe) Pay(amount float32) {
	fmt.Println("Making payment using Stripe:", amount)
}

// Razorpay struct represents another payment system (Razorpay).
type Razorpay struct{}

// Pay method for Razorpay.
// Same reasoning as Stripe: Razorpay provides the `Pay` method,
// so it also automatically implements PaymentGateway.
// No explicit "implements PaymentGateway" is needed in Go.
func (r Razorpay) Pay(amount float32) {
	fmt.Println("Making payment using Razorpay:", amount)
}

//
// ---------------- DEPENDENCY INJECTION USING INTERFACES ----------------
// Instead of hardcoding Stripe or Razorpay inside the Payment struct,
// we depend on the interface (PaymentGateway).
// This allows flexibility: any type implementing PaymentGateway
// can be injected into Payment at runtime.
// -----------------------------------------------------------------------
//

// Payment struct represents a payment request handler.
// It depends on the interface `PaymentGateway`, not on any specific provider.
// This is an example of **Dependency Inversion Principle (DIP)**.
type Payment struct {
	Gateway PaymentGateway
}

// MakePayment calls the `Pay` method of whichever concrete payment gateway
// was injected into the Payment struct.
// This keeps the code flexible and extensible.
func (p Payment) MakePayment(amount float32) {
	p.Gateway.Pay(amount)
}

// ----------------------------- MAIN -----------------------------------
// In main(), we demonstrate using Payment struct with two different
// payment providers (Stripe and Razorpay).
//
// Notice how we don’t need to change the Payment struct at all.
// We just pass in a different Gateway (interface implementation).
// -----------------------------------------------------------------------
func main() {
	// Using Stripe as the payment provider
	stripePayment := Payment{
		Gateway: Stripe{}, // Inject Stripe implementation
	}
	stripePayment.MakePayment(1000) // Output: Making payment using Stripe: 1000

	// Using Razorpay as the payment provider
	razorpayPayment := Payment{
		Gateway: Razorpay{}, // Inject Razorpay implementation
	}
	razorpayPayment.MakePayment(2000) // Output: Making payment using Razorpay: 2000
}

//
// ---------------------------- SUMMARY ---------------------------------
// 1. Interfaces in Go define behavior, not data.
// 2. Any type that implements the required methods automatically satisfies
//    the interface (no "implements" keyword needed).
// 3. PaymentGateway interface has one method: Pay(amount float32).
// 4. Stripe and Razorpay structs both implement the Pay method,
//    so both satisfy PaymentGateway.
// 5. Payment struct depends only on the interface (not concrete types),
//    allowing flexibility and easier testing/mock implementations.
// 6. This demonstrates Dependency Injection: behavior is chosen at runtime
//    by passing the desired implementation into Payment.
// -----------------------------------------------------------------------
//
