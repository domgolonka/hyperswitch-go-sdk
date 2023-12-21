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

type Address struct {
	// The first line of the address
	Line1 string `json:"line1,omitempty"`
	// The second line of the address
	Line2 string `json:"line2,omitempty"`
	// The second line of the address
	Line3 string `json:"line3,omitempty"`
	// The address city
	City string `json:"city,omitempty"`
	// The address state
	State string `json:"state,omitempty"`
	// The address zip/postal code
	Zip string `json:"zip,omitempty"`
	// The two-letter ISO country code
	Country string `json:"country,omitempty"`
}