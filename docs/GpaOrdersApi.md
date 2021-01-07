# \GpaOrdersApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGpaordersToken**](GpaOrdersApi.md#GetGpaordersToken) | **Get** /gpaorders/{token} | Returns a GPA order
[**GetGpaordersUnloads**](GpaOrdersApi.md#GetGpaordersUnloads) | **Get** /gpaorders/unloads | Lists all GPA returns
[**GetGpaordersUnloadsUnloadtoken**](GpaOrdersApi.md#GetGpaordersUnloadsUnloadtoken) | **Get** /gpaorders/unloads/{unload_token} | Returns a specific GPA return
[**PostGpaorders**](GpaOrdersApi.md#PostGpaorders) | **Post** /gpaorders | Funds a user&#39;s GPA account
[**PostGpaordersUnloads**](GpaOrdersApi.md#PostGpaordersUnloads) | **Post** /gpaorders/unloads | Returns a GPA order


# **GetGpaordersToken**
> GpaResponse GetGpaordersToken(ctx, token)
Returns a GPA order



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**GpaResponse**](gpa_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGpaordersUnloads**
> GpaUnloadListResponse GetGpaordersUnloads(ctx, optional)
Lists all GPA returns



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GpaOrdersApiGetGpaordersUnloadsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GpaOrdersApiGetGpaordersUnloadsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of GPA unloads to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]
 **userToken** | **optional.String**| User token | 
 **businessToken** | **optional.String**| Business token | 
 **originalOrderToken** | **optional.String**| Original order token | 

### Return type

[**GpaUnloadListResponse**](GPAUnloadListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGpaordersUnloadsUnloadtoken**
> GpaReturns GetGpaordersUnloadsUnloadtoken(ctx, unloadToken)
Returns a specific GPA return



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unloadToken** | **string**| Unload token | 

### Return type

[**GpaReturns**](gpa_returns.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGpaorders**
> GpaResponse PostGpaorders(ctx, optional)
Funds a user's GPA account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GpaOrdersApiPostGpaordersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GpaOrdersApiPostGpaordersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GpaRequest**](GpaRequest.md)|  | 

### Return type

[**GpaResponse**](gpa_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGpaordersUnloads**
> GpaReturns PostGpaordersUnloads(ctx, optional)
Returns a GPA order



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GpaOrdersApiPostGpaordersUnloadsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GpaOrdersApiPostGpaordersUnloadsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UnloadRequestModel**](UnloadRequestModel.md)|  | 

### Return type

[**GpaReturns**](gpa_returns.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

