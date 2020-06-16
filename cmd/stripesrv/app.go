package main

import (
	"encoding/json"
	"log"
	"github.com/stripe/stripe-go/paymentintent"
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
)

const (
	port = "3000"
	stripeKey = "sk_test_n287Bossh5n9yH7zyC6mXZIf00guWEHwUL"
)

type CheckoutData struct {
	ClientSecret string `json:"client_secret"`
}

func main() {
	log.Println("---------------------------------------------------------------------------")
	log.Println("--")
	log.Println("--  Started server on port " + port)
	log.Println("--")
	log.Println("---------------------------------------------------------------------------")
	initRoutes()
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initRoutes() {
	http.HandleFunc("/secret", func(w http.ResponseWriter, r *http.Request) {
		intent, err := createPaymentIntent() // ... Fetch or create the PaymentIntent
		if err != nil {
			log.Fatal(err)
		}

		data := CheckoutData{
			ClientSecret: intent.ClientSecret,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	})

	//http.HandlerFunc("/", server.Init)
}

func createPaymentIntent() (*stripe.PaymentIntent, error) {
	stripe.Key = stripeKey
	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(1299),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
	}
	// Verify your integration in this guide by including this parameter
	params.AddMetadata("integration_check", "accept_a_payment")

	return paymentintent.New(params)
}
