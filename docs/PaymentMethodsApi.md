# {{classname}}

All URIs are relative to *https://sandbox.hyperswitch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentMethodForCustomer**](PaymentMethodsApi.md#CreatePaymentMethodForCustomer) | **Post** /payment_methods | PaymentMethods - Create
[**DeletePaymentMethods**](PaymentMethodsApi.md#DeletePaymentMethods) | **Post** /payment_methods/{id}/detach | Delete PaymentMethods
[**ListPaymentMethodsByCustomerId**](PaymentMethodsApi.md#ListPaymentMethodsByCustomerId) | **Get** /customers/{customer_id}/payment_methods | List payment methods for a Customer
[**ListPaymentMethodsByMerchantId**](PaymentMethodsApi.md#ListPaymentMethodsByMerchantId) | **Get** /payment_methods/{merchant_id} | List payment methods for a Merchant
[**UpdatePaymentMethodForCustomer**](PaymentMethodsApi.md#UpdatePaymentMethodForCustomer) | **Post** /payment_methods/{id} | PaymentMethods - Update

# **CreatePaymentMethodForCustomer**
> PaymentMethodsResponseObject CreatePaymentMethodForCustomer(ctx, optional)
PaymentMethods - Create

To create a payment method against a customer object. In case of cards, this API could be used only by PCI compliant merchants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentMethodsApiCreatePaymentMethodForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentMethodsApiCreatePaymentMethodForCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PaymentMethodsCreateRequest**](PaymentMethodsCreateRequest.md)|  | 

### Return type

[**PaymentMethodsResponseObject**](PaymentMethodsResponseObject.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePaymentMethods**
> PaymentMethodDeleteResponse DeletePaymentMethods(ctx, id)
Delete PaymentMethods

Detaches a PaymentMethod object from a Customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The unique identifier for the payment method | 

### Return type

[**PaymentMethodDeleteResponse**](PaymentMethodDeleteResponse.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPaymentMethodsByCustomerId**
> CustomerPaymentMethods ListPaymentMethodsByCustomerId(ctx, optional)
List payment methods for a Customer

To filter and list the applicable payment methods for a particular Customer ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentMethodsApiListPaymentMethodsByCustomerIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentMethodsApiListPaymentMethodsByCustomerIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptedCountry** | [**optional.Interface of []string**](string.md)|  | 
 **acceptedCurrency** | [**optional.Interface of []string**](string.md)|  | 
 **minimumAmount** | **optional.Int32**|  | 
 **maximumAmount** | **optional.Int32**|  | 
 **recurringPaymentEnabled** | **optional.Bool**|  | 
 **installmentPaymentEnabled** | **optional.Bool**|  | 

### Return type

[**CustomerPaymentMethods**](CustomerPaymentMethods.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPaymentMethodsByMerchantId**
> []ArrayOfMerchantPaymentMethodsInner ListPaymentMethodsByMerchantId(ctx, optional)
List payment methods for a Merchant

To filter and list the applicable payment methods for a particular merchant id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentMethodsApiListPaymentMethodsByMerchantIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentMethodsApiListPaymentMethodsByMerchantIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptedCountry** | [**optional.Interface of []string**](string.md)|  | 
 **acceptedCurrency** | [**optional.Interface of []string**](string.md)|  | 
 **minimumAmount** | **optional.Int32**|  | 
 **maximumAmount** | **optional.Int32**|  | 
 **recurringPaymentEnabled** | **optional.Bool**|  | 
 **installmentPaymentEnabled** | **optional.Bool**|  | 

### Return type

[**[]ArrayOfMerchantPaymentMethodsInner**](array.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePaymentMethodForCustomer**
> PaymentMethodsResponseObject UpdatePaymentMethodForCustomer(ctx, id, optional)
PaymentMethods - Update

To update an existing a payment method attached to a customer object. This API is useful for use cases such as updating the card number for expired cards, to prevent discontinuity in recurring payments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The unique identifier for the payment method | 
 **optional** | ***PaymentMethodsApiUpdatePaymentMethodForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentMethodsApiUpdatePaymentMethodForCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PaymentMethodsUpdateRequest**](PaymentMethodsUpdateRequest.md)|  | 

### Return type

[**PaymentMethodsResponseObject**](PaymentMethodsResponseObject.md)

### Authorization

[ApiSecretKey](../README.md#ApiSecretKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

