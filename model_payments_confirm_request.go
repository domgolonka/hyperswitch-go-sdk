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

type PaymentsConfirmRequest struct {
	// The URL to redirect after the completion of the operation
	ReturnUrl string `json:"return_url,omitempty"`
	// Indicates that you intend to make future payments with this Payment’s payment method. Possible values are: (i) REQUIRED: The payment will be processed only with payment methods eligible for recurring payments (ii) OPTIONAL: The payment may/ may not be processed with payment methods eligible for recurring payments
	SetupFutureUsage string `json:"setup_future_usage,omitempty"`
	// The transaction authentication can be set to undergo payer authentication. Possible values are: (i) THREE_DS: If the card is enrolled for 3DS authentication, the 3DS based authentication will be activated. The liability of chargeback shift to the issuer, (ii) NO_THREE_DS: 3DS based authentication will not be activated. The liability of chargeback stays with the merchant. By default, the authentication will be marked as NO_THREE_DS
	AuthenticationType string `json:"authentication_type,omitempty"`
	// The payment method
	PaymentMethod string `json:"payment_method,omitempty"`
	// The payment method information
	PaymentMethodData *OneOfPaymentsConfirmRequestPaymentMethodData `json:"payment_method_data,omitempty"`
	// Enable this flag as true, if the user has consented for saving the payment method information
	SavePaymentMethod bool `json:"save_payment_method,omitempty"`
	Billing *Address `json:"billing,omitempty"`
	Shipping *Address `json:"shipping,omitempty"`
}
