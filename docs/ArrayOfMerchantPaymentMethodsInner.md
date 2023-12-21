# ArrayOfMerchantPaymentMethodsInner

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethod** | [***PaymentMethod**](PaymentMethod.md) |  | [optional] [default to null]
**PaymentMethodTypes** | [**[]PaymentMethodType**](PaymentMethodType.md) |  | [optional] [default to null]
**PaymentMethodIssuers** | **[]string** | List of payment method issuers to be enabled for this payment method | [optional] [default to null]
**PaymentMethodIssuerCodes** | **string** | A standard code representing the issuer of payment method  | [optional] [default to null]
**PaymentSchemes** | [**[]PaymentSchemes**](PaymentSchemes.md) | List of payment schemes accepted or has the processing capabilities of the processor | [optional] [default to null]
**AcceptedCurrencies** | **[]string** | List of currencies accepted or has the processing capabilities of the processor | [optional] [default to null]
**AcceptedCountries** | **[]string** | List of Countries accepted or has the processing capabilities of the processor | [optional] [default to null]
**MinimumAmount** | **int32** | Minimum amount supported by the processor. To be represented in the lowest denomination of the target currency (For example, for USD it should be in cents) | [optional] [default to null]
**MaximumAmount** | **int32** | Maximum amount supported by the processor. To be represented in the lowest denomination of the target currency (For example, for USD it should be in cents) | [optional] [default to null]
**RecurringEnabled** | **bool** | Boolean to enable recurring payments / mandates. Default is true. | [optional] [default to null]
**InstallmentPaymentEnabled** | **bool** | Boolean to enable installment / EMI / BNPL payments. Default is true. | [optional] [default to null]
**PaymentExperience** | [**[]NextAction**](NextAction.md) | Type of payment experience enabled with the connector | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

