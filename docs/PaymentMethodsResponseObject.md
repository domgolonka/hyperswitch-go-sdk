# PaymentMethodsResponseObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodId** | **string** | The identifier for the payment method object. | [default to null]
**PaymentMethod** | **string** | The type of payment method use for the payment.  | [default to null]
**PaymentMethodType** | **string** | This is a sub-category of payment method.  | [optional] [default to null]
**PaymentMethodIssuer** | **string** | The name of the bank/ provider issuing the payment method to the end user  | [optional] [default to null]
**PaymentMethodIssuerCode** | **string** | A standard code representing the issuer of payment method  | [optional] [default to null]
**Card** | [***CardPaymentMethodResponse**](CardPaymentMethodResponse.md) |  | [optional] [default to null]
**PaymentScheme** | **[]string** |  | [optional] [default to null]
**AcceptedCountry** | **[]string** |  | [optional] [default to null]
**AcceptedCurrency** | **[]string** |  | [optional] [default to null]
**MinimumAmount** | **int32** | The minimum amount accepted for processing by the particular payment method.  | [optional] [default to null]
**MaximumAmount** | **int32** | The minimum amount accepted for processing by the particular payment method.  | [optional] [default to null]
**RecurringPaymentEnabled** | **bool** | Indicates whether the payment method is eligible for recurring payments | [optional] [default to null]
**InstallmentPaymentEnabled** | **bool** | Indicates whether the payment method is eligible for installment payments | [optional] [default to null]
**PaymentExperience** | **[]string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

