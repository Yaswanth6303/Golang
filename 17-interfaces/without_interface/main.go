package main

import "fmt"

// ------------------------------
// The Problem Setup
// ------------------------------

// payment struct represents a payment processing system.
// Currently, it is tightly coupled with a specific payment gateway (stripe).
// This violates the Open/Closed Principle (OCP).
// OCP states: "Software entities should be open for extension,
// but closed for modification."
//
// Meaning → We should be able to add new functionality (like new payment gateways)
// without modifying existing code.
//
// In this case, if we want to switch from Stripe to Razorpay (or add PayPal, etc.),
// we need to change this struct and also update all the places where it's used.
type payment struct {
	gateway stripe // ❌ Hardcoding: Payment is tied only to Stripe
}

// makePayment is the function responsible for processing payments.
//
// ⚠️ Problem:
// Right now, this method is tightly coupled to a concrete type (stripe).
// If we want to use another gateway (e.g., Razorpay), we need to modify this function.
// That breaks OCP — since we're "modifying" instead of just "extending."
//
// Example of what happens when we modify:
//   - Uncomment below code, and we hardcode Razorpay again
//   - Every time a new gateway is added, we modify this method → BAD design
func (p payment) makePayment(amount float32) {
	// Old approach (BAD):
	// razorpayPaymentGW := razorpay{}
	// razorpayPaymentGW.pay(amount)

	// stripePaymentGW := stripe{}
	// stripePaymentGW.pay(amount)

	// Current hardcoded approach (still BAD):
	p.gateway.pay(amount)
}

// ------------------------------
// Payment Gateways Implementations
// ------------------------------

// razorpay struct represents Razorpay's payment gateway
type razorpay struct{}

// pay is Razorpay's implementation of processing a payment
func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using Razorpay:", amount)
}

// stripe struct represents Stripe's payment gateway
type stripe struct{}

// pay is Stripe's implementation of processing a payment
func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using Stripe:", amount)
}

// ------------------------------
// Main Function
// ------------------------------

func main() {
	// ⚠️ Issue:
	// Right now, if we want to use Razorpay, we need to go back and
	// modify the "payment" struct to change the type.
	//
	// This means:
	// - We must change the code in multiple places (struct, test files, etc.)
	// - Our system becomes rigid and hard to maintain
	//
	// This clearly violates the Open/Closed Principle.

	// Example of creating a payment with Stripe (currently hardcoded):
	stripePaymentGW := stripe{}
	newPayment := payment{
		gateway: stripePaymentGW,
	}

	newPayment.makePayment(1000) // Output: Making payment using Stripe: 1000

	// ✅ Proper Fix (not shown yet here):
	// Instead of tying payment to "stripe", we should use an "interface".
	// That way, payment struct depends only on the abstraction,
	// not the concrete implementation.
	//
	// With an interface:
	//   type gateway interface { pay(amount float32) }
	//   type payment struct { gateway gateway }
	//
	// Then any new gateway (Stripe, Razorpay, PayPal, etc.)
	// can be added WITHOUT modifying existing code.
}
