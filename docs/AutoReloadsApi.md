# \AutoReloadsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutoreloads**](AutoReloadsApi.md#GetAutoreloads) | **Get** /autoreloads | Lists all auto reloads for the program
[**GetAutoreloadsToken**](AutoReloadsApi.md#GetAutoreloadsToken) | **Get** /autoreloads/{token} | Returns a specific auto reload object
[**PostAutoreloads**](AutoReloadsApi.md#PostAutoreloads) | **Post** /autoreloads | Creates an auto reload object
[**PutAutoreloadsToken**](AutoReloadsApi.md#PutAutoreloadsToken) | **Put** /autoreloads/{token} | Updates a specific auto reload object


# **GetAutoreloads**
> AutoReloadListResponse GetAutoreloads(ctx, optional)
Lists all auto reloads for the program



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutoReloadsApiGetAutoreloadsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoReloadsApiGetAutoreloadsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardProduct** | **optional.String**| Card product token | 
 **userToken** | **optional.String**| User token | 
 **businessToken** | **optional.String**| Business token | 
 **count** | **optional.Int32**| Number of items to retrieve. Count can be between 1 - 10 items. | [default to 10]
 **startIndex** | **optional.Int32**| Indicates from what row to start returning data. | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Field by which to sort the returned items. Use any field in the model, or system fields lastModifiedTime or createdTime. | [default to -lastModifiedTime]

### Return type

[**AutoReloadListResponse**](AutoReloadListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutoreloadsToken**
> AutoReloadResponseModel GetAutoreloadsToken(ctx, token, optional)
Returns a specific auto reload object



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Auto reload token | 
 **optional** | ***AutoReloadsApiGetAutoreloadsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoReloadsApiGetAutoreloadsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**AutoReloadResponseModel**](auto_reload_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutoreloads**
> AutoReloadResponseModel PostAutoreloads(ctx, body)
Creates an auto reload object



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AutoReloadModel**](AutoReloadModel.md)| Auto reload object | 

### Return type

[**AutoReloadResponseModel**](auto_reload_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAutoreloadsToken**
> AutoReloadResponseModel PutAutoreloadsToken(ctx, body, token)
Updates a specific auto reload object



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AutoReloadUpdateModel**](AutoReloadUpdateModel.md)| Auto reload object | 
  **token** | **string**|  | 

### Return type

[**AutoReloadResponseModel**](auto_reload_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

