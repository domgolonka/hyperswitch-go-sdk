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

type ArrayOfRoutingRulesInner struct {
	// The List of payment methods to include for this routing rule
	PaymentMethodsIncl []PaymentMethod `json:"payment_methods_incl,omitempty"`
	// The List of payment methods to exclude for this routing rule. If there is conflict between include and exclude lists, include list overrides the exclude list.
	PaymentMethodsExcl []PaymentMethod `json:"payment_methods_excl,omitempty"`
	// The List of payment method types to include for this routing rule
	PaymentMethodTypesIncl []PaymentMethodType `json:"payment_method_types_incl,omitempty"`
	// The List of payment method types to exclude for this routing rule. If there is conflict between include and exclude lists, include list overrides the exclude list.
	PaymentMethodTypesExcl []PaymentMethodType `json:"payment_method_types_excl,omitempty"`
	// The List of payment method issuers to include for this routing rule
	PaymentMethodIssuersIncl []string `json:"payment_method_issuers_incl,omitempty"`
	// The List of payment method issuers to exclude for this routing rule. If there is conflict between include and exclude lists, include list overrides the exclude list.
	PaymentMethodIssuersExcl []string `json:"payment_method_issuers_excl,omitempty"`
	// The List of countries to include for this routing rule
	CountriesIncl []string `json:"countries_incl,omitempty"`
	// The List of countries to exclude for this routing rule. If there is conflict between include and exclude lists, include list overrides the exclude list.
	CountriesExcl []string `json:"countries_excl,omitempty"`
	// The List of currencies to include for this routing rule
	CurrenciesIncl []string `json:"currencies_incl,omitempty"`
	// The List of currencies to exclude for this routing rule. If there is conflict between include and exclude lists, include list overrides the exclude list.
	CurrenciesExcl []string `json:"currencies_excl,omitempty"`
	// List of Metadata Filters to apply for the Routing Rule. The filters are presented as 2 arrays of keys and value. This property contains all the keys.
	MetadataFiltersKeys []string `json:"metadata_filters_keys,omitempty"`
	// List of Metadata Filters to apply for the Routing Rule. The filters are presented as 2 arrays of keys and value. This property contains all the keys.
	MetadataFiltersValues []string `json:"metadata_filters_values,omitempty"`
	// The pecking order of payment connectors (or processors) to be used for routing. The first connector in the array will be attempted for routing. If it fails, the second connector will be used till the list is exhausted.
	ConnectorsPeckingOrder []PaymentConnector `json:"connectors_pecking_order,omitempty"`
	// An Array of Connectors (as Keys) with the associated percentage of traffic to be routed through the given connector (Expressed as an array of values)
	ConnectorsTrafficWeightageKey []PaymentConnector `json:"connectors_traffic_weightage_key,omitempty"`
	// An Array of Weightage (expressed in percentage) that needs to be associated with the respective connectors (Expressed as an array of keys)
	ConnectorsTrafficWeightageValue []int32 `json:"connectors_traffic_weightage_value,omitempty"`
}