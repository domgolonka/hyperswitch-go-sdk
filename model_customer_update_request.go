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

// The customer details to be updated
type CustomerUpdateRequest struct {
	// The customer's email address
	Email string `json:"email,omitempty"`
	// The customer's name
	Name string `json:"name,omitempty"`
	// The country code for the customer phone number
	PhoneCountryCode string `json:"phone_country_code,omitempty"`
	// The customer's phone number
	Phone string `json:"phone,omitempty"`
	Address *Address `json:"address,omitempty"`
	// An arbitrary string that you can attach to a customer object.
	Description string `json:"description,omitempty"`
	// You can specify up to 50 keys, with key names up to 40 characters long and values up to 500 characters long. Metadata is useful for storing additional, structured information on an object.
	Metadata *interface{} `json:"metadata,omitempty"`
}
