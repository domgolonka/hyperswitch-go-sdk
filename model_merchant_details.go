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

type MerchantDetails struct {
	// The merchant's primary contact name
	PrimaryContactPerson string `json:"primary_contact_person,omitempty"`
	// The merchant's primary email address
	PrimaryEmail string `json:"primary_email,omitempty"`
	// The merchant's primary phone number
	PrimaryPhone string `json:"primary_phone,omitempty"`
	// The merchant's secondary contact name
	SecondaryContactPerson string `json:"secondary_contact_person,omitempty"`
	// The merchant's secondary email address
	SecondaryEmail string `json:"secondary_email,omitempty"`
	// The merchant's secondary phone number
	SecondaryPhone string `json:"secondary_phone,omitempty"`
	// The business website of the merchant
	Website string `json:"website,omitempty"`
	// A brief description about merchant's business
	AboutBusiness string `json:"about_business,omitempty"`
	Address *Address `json:"address,omitempty"`
}
