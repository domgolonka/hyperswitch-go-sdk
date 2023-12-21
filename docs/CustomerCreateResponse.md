# CustomerCreateResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The identifier for the customer object. If not provided the customer ID will be autogenerated. | [default to null]
**Email** | **string** | The customer&#x27;s email address | [optional] [default to null]
**Name** | **string** | The customer&#x27;s name | [optional] [default to null]
**PhoneCountryCode** | [***interface{}**](interface{}.md) | The country code for the customer phone number | [optional] [default to null]
**Phone** | [***interface{}**](interface{}.md) | The customer&#x27;s phone number | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**Description** | [***interface{}**](interface{}.md) | An arbitrary string that you can attach to a customer object. | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | You can specify up to 50 keys, with key names up to 40 characters long and values up to 500 characters long. Metadata is useful for storing additional, structured information on an object. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
