# PaymentConnectorResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The identifier for the Merchant Account. | [default to null]
**ConnectorType** | [***ConnectorType**](ConnectorType.md) |  | [default to null]
**ConnectorName** | [***PaymentConnector**](PaymentConnector.md) |  | [default to null]
**ConnectorId** | **int32** | Unique ID of the connector | [default to null]
**ConnectorAccountDetails** | [***interface{}**](interface{}.md) | Account details of the Connector. You can specify up to 50 keys, with key names up to 40 characters long and values up to 500 characters long. Useful for storing additional, structured information on an object. | [optional] [default to null]
**TestMode** | **bool** | A boolean value to indicate if the connector is in Test mode. By default, its value is false. | [optional] [default to null]
**Disabled** | **bool** | A boolean value to indicate if the connector is disabled. By default, its value is false. | [optional] [default to null]
**PaymentMethods** | [***[]PaymentMethodsEnabledInner**](array.md) |  | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | You can specify up to 50 keys, with key names up to 40 characters long and values up to 500 characters long. Metadata is useful for storing additional, structured information on an object. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

