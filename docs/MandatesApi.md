# {{classname}}

All URIs are relative to *https://sandbox.hyperswitch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMandateDetails**](MandatesApi.md#ListMandateDetails) | **Get** /customers/{customer_id}/mandates | Mandate - List all mandates against a customer id
[**ListMandateDetails_0**](MandatesApi.md#ListMandateDetails_0) | **Get** /mandates/{id} | Mandate - List details of a mandate
[**RevokeMandateDetails**](MandatesApi.md#RevokeMandateDetails) | **Post** /mandates/{id}/revoke | Mandate - Revoke a mandate

# **ListMandateDetails**
> []MandateListResponse ListMandateDetails(ctx, customerId)
Mandate - List all mandates against a customer id

To list the all the mandates for a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Unique customer id for which the list of mandates to be retrieved. | 

### Return type

[**[]MandateListResponse**](MandateListResponse.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMandateDetails_0**
> MandateListResponse ListMandateDetails_0(ctx, id)
Mandate - List details of a mandate

To list the details of a mandate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Unique mandate id | 

### Return type

[**MandateListResponse**](MandateListResponse.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeMandateDetails**
> InlineResponse200 RevokeMandateDetails(ctx, id)
Mandate - Revoke a mandate

To revoke a mandate registered against a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Unique mandate id | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

