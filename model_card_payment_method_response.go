/*
 * Juspay Router - API Documentation
 *
 * ## Get started Juspay Router provides a collection of APIs that enable you to process and manage payments. Our APIs accept and return JSON in the HTTP body, and return standard HTTP response codes</a>. You can consume the APIs directly using your favorite HTTP/REST library. We have a testing environment referred to \"sandbox\", which you can setup to test API calls without affecting production data. ### Base URLs Use the following base URLs when making requests to the APIs:    | Environment   |      Base URL                                        |   |---------------|------------------------------------------------------|   | Sandbox       | <https://sandbox.hyperswitch.io>                   |   | Production    | <https://router.juspay.io>                           |  # Authentication When you sign up on our [dashboard](https://app.hyperswitch.io) and create a merchant account, you are given a secret key (also referred as api-key). You may authenticate all API requests with Juspay server by providing the appropriate key in the request Authorization header. Never share your secret api keys. Keep them guarded and secure. 
 *
 * API version: 0.2
 * Contact: support@juspay.in
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Card Payment Method object
type CardPaymentMethodResponse struct {
	// The last four digits of the case which could be displayed to the end user for identification.
	Last4Digits string `json:"last4_digits,omitempty"`
	// The expiry month for the card
	CardExpMonth string `json:"card_exp_month,omitempty"`
	// Expiry year for the card
	CardExpYear string `json:"card_exp_year,omitempty"`
	// The name of card holder
	CardHolderName string `json:"card_holder_name,omitempty"`
	// The token provided against a user's saved card. The token would be valid for 15 minutes.
	CardToken string `json:"card_token,omitempty"`
	// The card scheme network for the particular card
	Scheme string `json:"scheme,omitempty"`
	// The country code in in which the card was issued
	IssuerCountry string `json:"issuer_country,omitempty"`
	// A unique identifier alias to identify a particular card.
	CardFingerprint string `json:"card_fingerprint,omitempty"`
}