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

import (
	"net/http"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "https://sandbox.hyperswitch.io",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}