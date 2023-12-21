# {{classname}}

All URIs are relative to *https://sandbox.hyperswitch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMerchantAccount**](MerchantAccountsApi.md#CreateMerchantAccount) | **Post** /accounts | Merchant Account - Create
[**DeleteMerchantAccount**](MerchantAccountsApi.md#DeleteMerchantAccount) | **Delete** /accounts/{id} | Merchant Account - Delete
[**RetrieveMerchantAccount**](MerchantAccountsApi.md#RetrieveMerchantAccount) | **Get** /accounts/{id} | Merchant Account - Retrieve
[**UpdateMerchantAccount**](MerchantAccountsApi.md#UpdateMerchantAccount) | **Post** /accounts/{id} | Merchant Account - Update

# **CreateMerchantAccount**
> MerchantAccountResponse CreateMerchantAccount(ctx, optional)
Merchant Account - Create

Create a new account for a merchant. The merchant could be a seller or retailer or client who likes to receive and send payments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MerchantAccountsApiCreateMerchantAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantAccountsApiCreateMerchantAccountOpts struct
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

# **DeleteMerchantAccount**
> MerchantAccountDeleteResponse DeleteMerchantAccount(ctx, id)
Merchant Account - Delete

Delete a Merchant Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The unique identifier for the merchant account | 

### Return type

[**MerchantAccountDeleteResponse**](MerchantAccountDeleteResponse.md)

### Authorization

[AdminSecretKey](../README.md#AdminSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveMerchantAccount**
> MerchantAccountResponse RetrieveMerchantAccount(ctx, id)
Merchant Account - Retrieve

Retrieve a merchant account details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The unique identifier for the merchant account | 

### Return type

[**MerchantAccountResponse**](MerchantAccountResponse.md)

### Authorization

[AdminSecretKey](../README.md#AdminSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMerchantAccount**
> MerchantAccountResponse UpdateMerchantAccount(ctx, id, optional)
Merchant Account - Update

To update an existing merchant account. Helpful in updating merchant details such as email, contact details, or other configuration details like webhook, routing algorithm etc

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The unique identifier for the merchant account | 
 **optional** | ***MerchantAccountsApiUpdateMerchantAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantAccountsApiUpdateMerchantAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MerchantAccountUpdateRequest**](MerchantAccountUpdateRequest.md)|  | 

### Return type

[**MerchantAccountResponse**](MerchantAccountResponse.md)

### Authorization

[AdminSecretKey](../README.md#AdminSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

