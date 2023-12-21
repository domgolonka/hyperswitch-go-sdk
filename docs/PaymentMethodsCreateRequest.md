# PaymentMethodsCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethod** | **string** | The type of payment method use for the payment.  | [default to null]
**PaymentMethodType** | **string** | This is a sub-category of payment method.  | [optional] [default to null]
**PaymentMethodIssuer** | **string** | The name of the bank/ provider issuing the payment method to the end user  | [optional] [default to null]
**PaymentMethodIssuerCode** | **string** | A standard code representing the issuer of payment method  | [optional] [default to null]
**Card** | [***CardPaymentMethodRequest**](CardPaymentMethodRequest.md) |  | [optional] [default to null]
**CustomerId** | **string** | The unique identifier of the Customer.  | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | You can specify up to 50 keys, with key names up to 40 characters long and values up to 500 characters long. Metadata is useful for storing additional, structured information on an object. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

