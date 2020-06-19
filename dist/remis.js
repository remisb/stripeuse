function fetchSecret() {
    var response = fetch('/secret?serviceId=ProAccountMonthlyPayment').then(function(response) {
        return response.json()
    }).then(function(responseJson) {
        var clientSecret = responseJson.client_secret;
        console.log("ClientSecret:", clientSecret)

        stripe.confirmCardPayment(clientSecret)
    })
}
