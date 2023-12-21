# MerchantAccountResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The identifier for the Merchant Account. | [default to null]
**ApiKey** | **string** | The public security key for accessing the APIs | [optional] [default to null]
**MerchantName** | **string** | Name of the merchant | [optional] [default to null]
**MerchantDetails** | [***MerchantDetails**](MerchantDetails.md) |  | [optional] [default to null]
**WebhookDetails** | [***WebhookDetails**](WebhookDetails.md) |  | [optional] [default to null]
**RoutingAlgorithm** | **string** | The routing algorithm to be used to process the incoming request from merchant to outgoing payment processor or payment method. The default is CUSTOM if | [optional] [default to null]
**CustomRoutingRules** | [***[]ArrayOfRoutingRulesInner**](array.md) |  | [optional] [default to null]
**SubMerchantsEnabled** | **bool** | A boolean value to indicate if the merchant is a sub-merchant under a master or a parent merchant. By default, its value is false. | [optional] [default to null]
**ParentMerchantId** | **string** | Refers to the Parent Merchant ID if the merchant being created is a sub-merchant | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | You can specify up to 50 keys, with key names up to 40 characters long and values up to 500 characters long. Metadata is useful for storing additional, structured information on an object. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

