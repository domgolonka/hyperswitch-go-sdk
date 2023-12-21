# {{classname}}

All URIs are relative to *https://sandbox.hyperswitch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPayment**](PaymentsApi.md#CancelPayment) | **Post** /payments/{id}/cancel | Payments - Cancel
[**CapturePayment**](PaymentsApi.md#CapturePayment) | **Post** /payments/{id}/capture | Payments - Capture
[**ConfirmPayment**](PaymentsApi.md#ConfirmPayment) | **Post** /payments/{id}/confirm | Payments - Confirm
[**CreatePayment**](PaymentsApi.md#CreatePayment) | **Post** /payments | Payments - Create
[**RetrievePayment**](PaymentsApi.md#RetrievePayment) | **Get** /payments/{id} | Payments - Retrieve
[**UpdatePayment**](PaymentsApi.md#UpdatePayment) | **Post** /payments/{id} | Payments - Update

# **CancelPayment**
> PaymentsCreateResponse CancelPayment(ctx, id, optional)
Payments - Cancel

A Payment could can be cancelled when it is in one of these statuses: requires_payment_method, requires_capture, requires_confirmation, requires_customer_action

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| unique payment id | 
 **optional** | ***PaymentsApiCancelPaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentsApiCancelPaymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PaymentsCancelRequest**](PaymentsCancelRequest.md)|  | 

### Return type

[**PaymentsCreateResponse**](PaymentsCreateResponse.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CapturePayment**
> PaymentsCreateResponse CapturePayment(ctx, id, optional)
Payments - Capture

To capture the funds for an uncaptured payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| unique payment id | 
 **optional** | ***PaymentsApiCapturePaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentsApiCapturePaymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PaymentsCaptureRequest**](PaymentsCaptureRequest.md)|  | 

### Return type

[**PaymentsCreateResponse**](PaymentsCreateResponse.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfirmPayment**
> PaymentsCreateResponse ConfirmPayment(ctx, id, optional)
Payments - Confirm

This API is to confirm the payment request and forward payment to the payment processor. This API provides more granular control upon when the API is forwarded to the payment processor. Alternatively you can confirm the payment within the Payments-Create API

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| unique payment id | 
 **optional** | ***PaymentsApiConfirmPaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentsApiConfirmPaymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PaymentsConfirmRequest**](PaymentsConfirmRequest.md)|  | 

### Return type

[**PaymentsCreateResponse**](PaymentsCreateResponse.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePayment**
> InlineResponse2001 CreatePayment(ctx, optional)
Payments - Create

To process a payment you will have to create a payment, attach a payment method and confirm. Depending on the user journey you wish to achieve, you may opt to all the steps in a single request or in a sequence of API request using following APIs: (i) Payments - Update, (ii) Payments - Confirm, and (iii) Payments - Capture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentsApiCreatePaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentsApiCreatePaymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PaymentsBody**](PaymentsBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrievePayment**
> PaymentsCreateResponse RetrievePayment(ctx, id)
Payments - Retrieve

To retrieve the properties of a Payment. This may be used to get the status of a previously initiated payment or next action for an ongoing payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| unique payment id | 

### Return type

[**PaymentsCreateResponse**](PaymentsCreateResponse.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePayment**
> PaymentsCreateResponse UpdatePayment(ctx, id, optional)
Payments - Update

To update the properties of a PaymentIntent object. This may include attaching a payment method, or attaching customer object or metadata fields after the Payment is created 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| unique payment id | 
 **optional** | ***PaymentsApiUpdatePaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentsApiUpdatePaymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PaymentsUpdateRequest**](PaymentsUpdateRequest.md)|  | 

### Return type

[**PaymentsCreateResponse**](PaymentsCreateResponse.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

