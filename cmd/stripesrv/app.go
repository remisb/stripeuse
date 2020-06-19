package main

import (
	"encoding/json"
	"errors"
	"log"
	"github.com/stripe/stripe-go/paymentintent"
	serve "github.com/remisb/stripeuse/internal/server"
	"net/http"

	//stripe "github.com/stripe/stripe-go/v71"
	stripe "github.com/stripe/stripe-go"
)

const (
	port           = "3000"
	publishableKey = "pk_test_48VYROIMG5lQXPPvE7lFGdOg007hVI6jfG"
	stripeKey      = "sk_test_n287Bossh5n9yH7zyC6mXZIf00guWEHwUL"
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
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type ErrorResponse struct {
	Code    string `json:code`
	Message string `json:message`
}

func byteError(code, message string) ([]byte, error) {
	errResp := ErrorResponse{
		Code:    code,
		Message: message,
	}
	return json.Marshal(errResp)
}

func respondError(w http.ResponseWriter, statusCode int, code, message string) {
	response, err := byteError(code, message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error. We are working on it."))
		return
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
}

type ServiceCost struct {
	Amount int64
	Currency string
}

func getServiceCost(serviceId string) (ServiceCost, error) {
	if serviceId == "ProAccountMonthlyPayment" {
		return ServiceCost{
			Amount: int64(1000),
			Currency: string(stripe.CurrencyUSD),
		}, nil
	}

	return ServiceCost{}, errors.New("no such service")
}

// get /secret?serviceId=Subscription100
func secretHandler(w http.ResponseWriter, r *http.Request) {
	serviceId := r.URL.Query().Get("serviceId")
	if serviceId == "" {
		respondError(w, http.StatusBadRequest, "E101", "missing required field serviceId")
		return
	}

	cost, err := getServiceCost(serviceId)
	if err != nil {
		respondError(w, http.StatusBadRequest, "E102", "unknown service")
	}

	stripe.Key = stripeKey
	params := &stripe.PaymentIntentParams{
		Amount:  &cost.Amount,
		Currency: stripe.String(cost.Currency),
	}
	params.AddMetadata("integration_check", "accept_a_payment")
	pIntent, err := paymentintent.New(params)
	if err != nil {
		log.Print(err)
		respondError(w, http.StatusInternalServerError, "E103", "error on payment intent creation")
	}

	data := CheckoutData{
		ClientSecret: pIntent.ClientSecret,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func initRoutes() {
	http.HandleFunc("/secret", secretHandler)

	http.HandleFunc("/stripe-key", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/create-payment-intent", func(w http.ResponseWriter, r *http.Request) {
		intent, err := createPaymentIntent() // ... Fetch or create the PaymentIntent
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(intent)
	})

	http.Handle("/", serve.InitRoutes())
}

func createPaymentIntent() (*stripe.PaymentIntent, error) {
	stripe.Key = stripeKey
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1299),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
	}
	// Verify your integration in this guide by including this parameter
	params.AddMetadata("integration_check", "accept_a_payment")

	return paymentintent.New(params)
}
