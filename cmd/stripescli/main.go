package main

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

//const stripeKey = "sk_test_n287Bossh5n9yH7zyC6mXZIf00guWEHwUL"

//const publishableKey = "pk_test_48VYROIMG5lQXPPvE7lFGdOg007hVI6jfG"
const stripeKey = "sk_test_n287Bossh5n9yH7zyC6mXZIf00guWEHwUL"

func main() {
	createPaymentIntent()
}

func a() {
	// Set your secret key. Remember to switch to your live secret key in production!
	// See your keys here: https://dashboard.stripe.com/account/apikeys
	stripe.Key = stripeKey

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(1000),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		ReceiptEmail: stripe.String("jenny.rosen@example.com"),
	}
	paymentintent.New(params)
}

// createPaymentIntent should be performed on the server side to prevent malicious customers
// from being able to choose their own prices.
func createPaymentIntent() {
	// Set your secret key. Remember to switch to your live secret key in production!
	// See your keys here: https://dashboard.stripe.com/account/apikeys
	stripe.Key = stripeKey

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(500),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String("Upgrade to Pro account"),
	}
	// Verify your integration in this guide by including this parameter
	params.AddMetadata("integration_check", "accept_a_payment")

	pi, err := paymentintent.New(params)
	if err != nil {
		return
	}
	fmt.Printf("pi: %#v", *pi)
}
