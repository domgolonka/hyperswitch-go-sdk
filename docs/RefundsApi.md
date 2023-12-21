# {{classname}}

All URIs are relative to *https://sandbox.hyperswitch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRefunds**](RefundsApi.md#CreateRefunds) | **Post** /refunds | Refunds - Create
[**RetrieveRefund**](RefundsApi.md#RetrieveRefund) | **Get** /refunds/{id} | Refunds - Retrieve
[**UpdateRefund**](RefundsApi.md#UpdateRefund) | **Post** /refunds/{id} | Refunds - Update

# **CreateRefunds**
> RefundsObject CreateRefunds(ctx, optional)
Refunds - Create

To create a refund against an already processed payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RefundsApiCreateRefundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefundsApiCreateRefundsOpts struct
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

# **UpdateRefund**
> RefundsObject UpdateRefund(ctx, id, optional)
Refunds - Update

To update the properties of a Refund object. This may include attaching a reason for the refund or metadata fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| unique refund id | 
 **optional** | ***RefundsApiUpdateRefundOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefundsApiUpdateRefundOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RefundsUpdateRequest**](RefundsUpdateRequest.md)|  | 

### Return type

[**RefundsObject**](RefundsObject.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

