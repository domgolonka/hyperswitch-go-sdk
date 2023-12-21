# {{classname}}

All URIs are relative to *https://sandbox.hyperswitch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomer**](CustomersApi.md#CreateCustomer) | **Post** /customers | Create Customer
[**DeleteCustomer**](CustomersApi.md#DeleteCustomer) | **Delete** /customers/{id} | Delete Customer
[**RetrieveCustomer**](CustomersApi.md#RetrieveCustomer) | **Get** /customers/{id} | Retrieve Customer
[**UpdateCustomer**](CustomersApi.md#UpdateCustomer) | **Post** /customers/{id} | Update Customer

# **CreateCustomer**
> CustomerCreateResponse CreateCustomer(ctx, body)
Create Customer

Create a customer object and store the customer details to be reused for future payments. Incase the customer already exists in the system, this API will respond with the customer details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomerCreateRequest**](CustomerCreateRequest.md)|  | 

### Return type

[**CustomerCreateResponse**](CustomerCreateResponse.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomer**
> CustomerDeleteResponse DeleteCustomer(ctx, id)
Delete Customer

Delete a customer record.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| unique customer id | 

### Return type

[**CustomerDeleteResponse**](CustomerDeleteResponse.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCustomer**
> CustomerCreateResponse RetrieveCustomer(ctx, id)
Retrieve Customer

Retrieve a customer's details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| unique customer id | 

### Return type

[**CustomerCreateResponse**](CustomerCreateResponse.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomer**
> CustomerCreateResponse UpdateCustomer(ctx, body, id)
Update Customer

Updates the customer's details in a customer object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomerUpdateRequest**](CustomerUpdateRequest.md)|  | 
  **id** | **string**| unique customer id | 

### Return type

[**CustomerCreateResponse**](CustomerCreateResponse.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

