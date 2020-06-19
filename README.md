API keys

Key publishable: pk_test_48VYROIMG5lQXPPvE7lFGdOg007hVI6jfG`
Key secret: `sk_test_n287Bossh5n9yH7zyC6mXZIf00guWEHwUL`


```
go get -u github.com/stripe/stripe-go

import (
  "github.com/stripe/stripe-go/v71"
)
```

## Notes

* Accept a payment - Payment Intents API
    - steps:
        - creating an object to track a payment | PaymentIntent
        - collecting card information
        - submitting the payment to stripe for processing


Objects

### PaymentIntent

`PaymentIntent` is used to track and handle all the states of the payment until
it's completed. It represents your intent to collect payment from a customer, tracking
charge attempts and payment state changes throughout process.

Steps:

1. Setup a Stripe / Server-side

Use `go get -u github.com/stripe/stripe-go` and `import "github.com/stripe/stripe-go/v71"`.

2. Create a PaymentIntent / Server-side  

Create a PaymentIntent on your server with an amount and currency.

3. Collect card details / Client - side

JavaScript side [documentation stripe.confirmPayment(...)](https://stripe.com/docs/js/payment_intents)
Client Side Stripe UI elements - [Stripe Elements](https://stripe.com/docs/stripe-js)
Stipe Checkout Forms - [Custom Forms](https://stripe.dev/elements-examples/)



## Finding solution

a) required data.publishableKey from the response
b) returned data = {client_secret: "pi_1GvRVvGo53Y1bozf8Vz2dWCk_secret_WgwbYcMXsJGgAE9YDslVh8p16"}
```JavaScript
//on line 42
var setupElements = function(data) {...}
```

## Questions

- When to use ChargeParams?

```go
params := &stripe.ChargeParams{
  Amount: stripe.Int64(2000),
  Currency: stripe.String("eur"),
}
...
charge.New(params)
```
