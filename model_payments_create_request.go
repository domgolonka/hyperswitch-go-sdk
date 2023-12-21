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

type PaymentsCreateRequest struct {
	// The payment amount. Amount for the payment in lowest denomination of the currency. (i.e) in cents for USD denomination, in paisa for INR denomination etc., 
	Amount int32 `json:"amount"`
	// The three-letter ISO currency code 
	Currency string `json:"currency"`
	// Unique Identifier for the Payment. It is always recommended to provide this ID while creating a payment. If the identifiers in not provided in the Payment Request, this filed shall be auto generated and provide in the API response. It is suggested to keep the payment_id length as a maximum of 30 alphanumeric characters irrespective of payment methods and gateways. Sequential and only numeric characters are not recommended. 
	PaymentId string `json:"payment_id,omitempty"`
	// Whether to confirm the payment (if applicable)
	Confirm bool `json:"confirm,omitempty"`
	// This is the instruction for capture/ debit the money from the users' card. On the other hand authorization refers to blocking the amount on the users' payment method. Capture request may happen in three types: (1) AUTOMATIC: Post the payment authorization, the capture will be executed on the full amount immediately, (2) MANUAL: The capture will happen only if the merchant triggers a Capture API request, (3) SCHEDULED: The capture can be scheduled to automatically get triggered at a specific date & time
	CaptureMethod string `json:"capture_method,omitempty"`
	// A timestamp (ISO 8601 code) that determines when the payment should be captured. Providing this field will automatically set `capture` to true 
	CaptureOn *AllOfPaymentsCreateRequestCaptureOn `json:"capture_on,omitempty"`
	// The Amount to be captured/ debited from the users payment method. It shall be in lowest denomination of the currency. (i.e) in cents for USD denomination, in paisa for INR denomination etc., If not provided, the default amount_to_capture will be the payment amount. 
	AmountToCapture int32 `json:"amount_to_capture,omitempty"`
	// The identifier for the customer object. If not provided the customer ID will be autogenerated.
	CustomerId string `json:"customer_id,omitempty"`
	// A description of the payment
	Description string `json:"description,omitempty"`
	// The customer's email address
	Email string `json:"email,omitempty"`
	// The customer's name
	Name string `json:"name,omitempty"`
	// The customer's phone number
	Phone string `json:"phone,omitempty"`
	// The country code for the customer phone number
	PhoneCountryCode string `json:"phone_country_code,omitempty"`
	// The URL to redirect after the completion of the operation
	ReturnUrl string `json:"return_url,omitempty"`
	// Indicates that you intend to make future payments with this Payment’s payment method. Providing this parameter will attach the payment method to the Customer, if present, after the Payment is confirmed and any required actions from the user are complete.
	SetupFutureUsage string `json:"setup_future_usage,omitempty"`
	// Set to true to indicate that the customer is not in your checkout flow during this payment, and therefore is unable to authenticate. This parameter is intended for scenarios where you collect card details and charge them later. This parameter can only be used with confirm=true.
	OffSession bool `json:"off_session,omitempty"`
	MandateData *PaymentsCreateRequestMandateData `json:"mandate_data,omitempty"`
	// ID of the mandate to be used for this payment. This parameter can only be used with confirm = true. This parameter should be passed for Merchant Initiated Transaction scenarios where the user has already registered for a mandate and raw payment method details are not available.
	MandateId string `json:"mandate_id,omitempty"`
	// The transaction authentication can be set to undergo payer authentication. Possible values are: (i) THREE_DS: If the card is enrolled for 3DS authentication, the 3DS based authentication will be activated. The liability of chargeback shift to the issuer, (ii) NO_THREE_DS: 3DS based authentication will not be activated. The liability of chargeback stays with the merchant. By default, the authentication will be marked as NO_THREE_DS
	AuthenticationType string `json:"authentication_type,omitempty"`
	// The payment method
	PaymentMethod string `json:"payment_method,omitempty"`
	PaymentMethodData *PaymentsCreateRequestPaymentMethodData `json:"payment_method_data,omitempty"`
	Billing *Address `json:"billing,omitempty"`
	Shipping *Address `json:"shipping,omitempty"`
	// For non-card charges, you can use this value as the complete description that appears on your customers’ statements. Must contain at least one letter, maximum 22 characters.
	StatementDescriptorName string `json:"statement_descriptor_name,omitempty"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the prefix (shortened descriptor) or statement descriptor that’s set on the account to form the complete statement descriptor. Maximum 22 characters for the concatenated descriptor.
	StatementDescriptorSuffix string `json:"statement_descriptor_suffix,omitempty"`
	// You can specify up to 50 keys, with key names up to 40 characters long and values up to 500 characters long. Metadata is useful for storing additional, structured information on an object.
	Metadata *interface{} `json:"metadata,omitempty"`
}
