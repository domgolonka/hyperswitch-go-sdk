# RefundsCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Total amount for which the refund is to be initiated. Amount for the payment in lowest denomination of the currency. (i.e) in cents for USD denomination, in paisa for INR denomination etc., If not provided, this will default to the full payment amount  | [optional] [default to null]
**RefundId** | **string** | Unique Identifier for the Refund. This is to ensure idempotency for multiple partial refund initiated against the same payment. If the identifiers is not defined by the merchant, this filed shall be auto generated and provide in the API response. It is recommended to generate uuid(v4) as the refund_id.  | [optional] [default to null]
**PaymentId** | **string** | Unique Identifier for the Payment. It is always recommended to provide this ID while creating a payment. If the identifiers in not provided in the Payment Request, this filed shall be auto generated and provide in the API response. It is suggested to keep the payment_id length as a maximum of 30 alphanumeric characters irrespective of payment methods and gateways. Sequential and only numeric characters are not recommended.  | [default to null]
**Currency** | **string** | The three-letter ISO currency code  | [optional] [default to null]
**Reason** | **string** | An arbitrary string attached to the object. Often useful for displaying to users and your customer support executive | [optional] [default to null]
**CreatedAt** | [***AllOfRefundsCreateRequestCreatedAt**](AllOfRefundsCreateRequestCreatedAt.md) | A timestamp (ISO 8601 code) that determines when the refund was created.  | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | You can specify up to 50 keys, with key names up to 40 characters long and values up to 500 characters long. Metadata is useful for storing additional, structured information on an object. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

