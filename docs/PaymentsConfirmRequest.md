# PaymentsConfirmRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnUrl** | **string** | The URL to redirect after the completion of the operation | [optional] [default to null]
**SetupFutureUsage** | **string** | Indicates that you intend to make future payments with this Payment’s payment method. Possible values are: (i) REQUIRED: The payment will be processed only with payment methods eligible for recurring payments (ii) OPTIONAL: The payment may/ may not be processed with payment methods eligible for recurring payments | [optional] [default to null]
**AuthenticationType** | **string** | The transaction authentication can be set to undergo payer authentication. Possible values are: (i) THREE_DS: If the card is enrolled for 3DS authentication, the 3DS based authentication will be activated. The liability of chargeback shift to the issuer, (ii) NO_THREE_DS: 3DS based authentication will not be activated. The liability of chargeback stays with the merchant. By default, the authentication will be marked as NO_THREE_DS | [optional] [default to AUTHENTICATION_TYPE.NO_THREE_DS]
**PaymentMethod** | **string** | The payment method | [optional] [default to three_ds]
**PaymentMethodData** | [***OneOfPaymentsConfirmRequestPaymentMethodData**](OneOfPaymentsConfirmRequestPaymentMethodData.md) | The payment method information | [optional] [default to null]
**SavePaymentMethod** | **bool** | Enable this flag as true, if the user has consented for saving the payment method information | [optional] [default to false]
**Billing** | [***Address**](Address.md) |  | [optional] [default to null]
**Shipping** | [***Address**](Address.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

