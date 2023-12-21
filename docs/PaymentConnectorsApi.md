# {{classname}}

All URIs are relative to *https://sandbox.hyperswitch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentConnector**](PaymentConnectorsApi.md#CreatePaymentConnector) | **Post** /account/{account_id}/connectors | Payment Connector - Create
[**DeletePaymentConnector**](PaymentConnectorsApi.md#DeletePaymentConnector) | **Delete** /account/{account_id}/connectors/{connector_id} | Payment Connector - Delete
[**RetrievePaymentConnector**](PaymentConnectorsApi.md#RetrievePaymentConnector) | **Get** /account/{account_id}/connectors/{connector_id} | Payment Connector - Retrieve
[**UpdatePaymentConnector**](PaymentConnectorsApi.md#UpdatePaymentConnector) | **Post** /account/{account_id}/connectors/{connector_id} | Payment Connector - Update

# **CreatePaymentConnector**
> PaymentConnectorResponse CreatePaymentConnector(ctx, accountId, optional)
Payment Connector - Create

Create a new Payment Connector for the merchant account. The connector could be a payment processor / facilitator / acquirer or specialized services like Fraud / Accounting etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| The unique identifier for the merchant account | 
 **optional** | ***PaymentConnectorsApiCreatePaymentConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentConnectorsApiCreatePaymentConnectorOpts struct
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

# **DeletePaymentConnector**
> PaymentConnectorDeleteResponse DeletePaymentConnector(ctx, id1, id2)
Payment Connector - Delete

Delete or Detach a Payment Connector from Merchant Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id1** | **string**| The unique identifier for the merchant account | 
  **id2** | **string**| The unique identifier for the payment connector | 

### Return type

[**PaymentConnectorDeleteResponse**](PaymentConnectorDeleteResponse.md)

### Authorization

[AdminSecretKey](../README.md#AdminSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrievePaymentConnector**
> PaymentConnectorResponse RetrievePaymentConnector(ctx, accountId, connectorId)
Payment Connector - Retrieve

Retrieve Payment Connector details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| The unique identifier for the merchant account | 
  **connectorId** | **string**| The unique identifier for the payment connector | 

### Return type

[**PaymentConnectorResponse**](PaymentConnectorResponse.md)

### Authorization

[AdminSecretKey](../README.md#AdminSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePaymentConnector**
> PaymentConnectorResponse UpdatePaymentConnector(ctx, id1, id2, optional)
Payment Connector - Update

To update an existing Payment Connector. Helpful in enabling / disabling different payment methods and other settings for the connector etc

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id1** | **string**| The unique identifier for the merchant account | 
  **id2** | **string**| The unique identifier for the payment connector | 
 **optional** | ***PaymentConnectorsApiUpdatePaymentConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentConnectorsApiUpdatePaymentConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PaymentConnectorUpdateRequest**](PaymentConnectorUpdateRequest.md)|  | 

### Return type

[**PaymentConnectorResponse**](PaymentConnectorResponse.md)

### Authorization

[AdminSecretKey](../README.md#AdminSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

