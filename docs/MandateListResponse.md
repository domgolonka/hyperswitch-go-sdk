# MandateListResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MandateId** | **string** | The unique id corresponding to the mandate.  | [default to null]
**Status** | **string** | The status of the mandate, which indicates whether it can be used to initiate a payment. | [default to null]
**Type_** | **string** | The type of the mandate. (i) single_use refers to one-time mandates and (ii) multi-user refers to multiple payments. | [optional] [default to TYPE_.MULTI_USE]
**PaymentMethodId** | **string** | The id corresponding to the payment method. | [default to null]
**PaymentMethod** | **string** | The type of payment method use for the payment.  | [optional] [default to null]
**Card** | [***CardPaymentMethodResponse**](CardPaymentMethodResponse.md) |  | [optional] [default to null]
**CustomerAcceptance** | [***CustomerAcceptance**](CustomerAcceptance.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

