# {{classname}}

All URIs are relative to *https://sandbox.hyperswitch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMerchantAccount**](QuickStartApi.md#CreateMerchantAccount) | **Post** /accounts | Merchant Account - Create
[**CreatePayment**](QuickStartApi.md#CreatePayment) | **Post** /payments | Payments - Create
[**CreatePaymentConnector**](QuickStartApi.md#CreatePaymentConnector) | **Post** /account/{account_id}/connectors | Payment Connector - Create
[**CreateRefunds**](QuickStartApi.md#CreateRefunds) | **Post** /refunds | Refunds - Create
[**RetrievePayment**](QuickStartApi.md#RetrievePayment) | **Get** /payments/{id} | Payments - Retrieve
[**RetrieveRefund**](QuickStartApi.md#RetrieveRefund) | **Get** /refunds/{id} | Refunds - Retrieve

# **CreateMerchantAccount**
> MerchantAccountResponse CreateMerchantAccount(ctx, optional)
Merchant Account - Create

Create a new account for a merchant. The merchant could be a seller or retailer or client who likes to receive and send payments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuickStartApiCreateMerchantAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuickStartApiCreateMerchantAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MerchantAccountCreateRequest**](MerchantAccountCreateRequest.md)|  | 

### Return type

[**MerchantAccountResponse**](MerchantAccountResponse.md)

### Authorization

[AdminSecretKey](../README.md#AdminSecretKey)

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
 **optional** | ***QuickStartApiCreatePaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuickStartApiCreatePaymentOpts struct
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

# **CreatePaymentConnector**
> PaymentConnectorResponse CreatePaymentConnector(ctx, accountId, optional)
Payment Connector - Create

Create a new Payment Connector for the merchant account. The connector could be a payment processor / facilitator / acquirer or specialized services like Fraud / Accounting etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| The unique identifier for the merchant account | 
 **optional** | ***QuickStartApiCreatePaymentConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuickStartApiCreatePaymentConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PaymentConnectorCreateRequest**](PaymentConnectorCreateRequest.md)|  | 

### Return type

[**PaymentConnectorResponse**](PaymentConnectorResponse.md)

### Authorization

[AdminSecretKey](../README.md#AdminSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRefunds**
> RefundsObject CreateRefunds(ctx, optional)
Refunds - Create

To create a refund against an already processed payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuickStartApiCreateRefundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuickStartApiCreateRefundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RefundsCreateRequest**](RefundsCreateRequest.md)|  | 

### Return type

[**RefundsObject**](RefundsObject.md)

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

# **RetrieveRefund**
> RefundsObject RetrieveRefund(ctx, id)
Refunds - Retrieve

To retrieve the properties of a Refund. This may be used to get the status of a previously initiated payment or next action for an ongoing payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| unique refund id | 

### Return type

[**RefundsObject**](RefundsObject.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

