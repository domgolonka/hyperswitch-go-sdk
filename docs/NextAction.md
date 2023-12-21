# NextAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectToUrl** | **string** | The URL to which the customer needs to be redirected for completing the payment. | [optional] [default to null]
**DisplayQrCode** | **string** | The QR code data to be displayed to the customer. | [optional] [default to null]
**InvokePaymentApp** | [***interface{}**](interface{}.md) | Contains the data for invoking the sdk client for completing the payment. | [optional] [default to null]
**InvokeSdkClient** | [***interface{}**](interface{}.md) | Contains the data for invoking the sdk client for completing the payment. | [optional] [default to null]
**TriggerApi** | [***interface{}**](interface{}.md) | Provides the instructions on the next API to be triggered to complete the payment. This is applicable in cases wherein the merchant has to display a UI to the user for collecting information such as OTP, 2factor authentication details. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

