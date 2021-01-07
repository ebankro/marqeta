# \MsaOrdersApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsaordersToken**](MsaOrdersApi.md#GetMsaordersToken) | **Get** /msaorders/{token} | Returns an MSA order 
[**GetMsaordersUnloads**](MsaOrdersApi.md#GetMsaordersUnloads) | **Get** /msaorders/unloads | Lists all MSA unloads
[**GetMsaordersUnloadsUnloadtoken**](MsaOrdersApi.md#GetMsaordersUnloadsUnloadtoken) | **Get** /msaorders/unloads/{unload_token} | Returns a specific MSA unload
[**PostMsaorders**](MsaOrdersApi.md#PostMsaorders) | **Post** /msaorders | Creates a merchant-specific account order
[**PostMsaordersUnloads**](MsaOrdersApi.md#PostMsaordersUnloads) | **Post** /msaorders/unloads | Returns an MSA order
[**PutMsaordersToken**](MsaOrdersApi.md#PutMsaordersToken) | **Put** /msaorders/{token} | Updates a specific merchant-specific account order


# **GetMsaordersToken**
> MsaOrderResponse GetMsaordersToken(ctx, token)
Returns an MSA order 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Order token | 

### Return type

[**MsaOrderResponse**](msa_order_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsaordersUnloads**
> MsaUnloadListResponse GetMsaordersUnloads(ctx, optional)
Lists all MSA unloads



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MsaOrdersApiGetMsaordersUnloadsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsaOrdersApiGetMsaordersUnloadsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of MSA unloads to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]
 **userToken** | **optional.String**| User token | 
 **businessToken** | **optional.String**| Business token | 
 **originalOrderToken** | **optional.String**| Original order token | 

### Return type

[**MsaUnloadListResponse**](MSAUnloadListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsaordersUnloadsUnloadtoken**
> MsaReturns GetMsaordersUnloadsUnloadtoken(ctx, unloadToken)
Returns a specific MSA unload



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unloadToken** | **string**| Unload token | 

### Return type

[**MsaReturns**](msa_returns.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMsaorders**
> MsaOrderResponse PostMsaorders(ctx, optional)
Creates a merchant-specific account order



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MsaOrdersApiPostMsaordersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsaOrdersApiPostMsaordersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MsaOrderRequest**](MsaOrderRequest.md)|  | 

### Return type

[**MsaOrderResponse**](msa_order_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMsaordersUnloads**
> MsaReturns PostMsaordersUnloads(ctx, optional)
Returns an MSA order



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MsaOrdersApiPostMsaordersUnloadsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsaOrdersApiPostMsaordersUnloadsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MsaUnloadRequestModel**](MsaUnloadRequestModel.md)|  | 

### Return type

[**MsaReturns**](msa_returns.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMsaordersToken**
> MsaOrderResponse PutMsaordersToken(ctx, token, optional)
Updates a specific merchant-specific account order



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Order token | 
 **optional** | ***MsaOrdersApiPutMsaordersTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsaOrdersApiPutMsaordersTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MsaOrderUpdateRequest**](MsaOrderUpdateRequest.md)|  | 

### Return type

[**MsaOrderResponse**](msa_order_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

