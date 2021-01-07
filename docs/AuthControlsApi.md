# \AuthControlsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthcontrols**](AuthControlsApi.md#GetAuthcontrols) | **Get** /authcontrols | Lists all global auth control exceptions for the program
[**GetAuthcontrolsExemptmids**](AuthControlsApi.md#GetAuthcontrolsExemptmids) | **Get** /authcontrols/exemptmids | Lists all auth control exempted MIDs for the program
[**GetAuthcontrolsExemptmidsToken**](AuthControlsApi.md#GetAuthcontrolsExemptmidsToken) | **Get** /authcontrols/exemptmids/{token} | Returns a specific auth control exemptmids
[**GetAuthcontrolsToken**](AuthControlsApi.md#GetAuthcontrolsToken) | **Get** /authcontrols/{token} | Returns a specific auth control exception
[**PostAuthcontrols**](AuthControlsApi.md#PostAuthcontrols) | **Post** /authcontrols | Creates an auth control exception
[**PostAuthcontrolsExemptmids**](AuthControlsApi.md#PostAuthcontrolsExemptmids) | **Post** /authcontrols/exemptmids | Creates an auth control for exempting MIDs
[**PutAuthcontrolsExemptmidsToken**](AuthControlsApi.md#PutAuthcontrolsExemptmidsToken) | **Put** /authcontrols/exemptmids/{token} | Updates the status an auth control exemptmids
[**PutAuthcontrolsToken**](AuthControlsApi.md#PutAuthcontrolsToken) | **Put** /authcontrols/{token} | Updates an auth control exception


# **GetAuthcontrols**
> AuthControlListResponse GetAuthcontrols(ctx, optional)
Lists all global auth control exceptions for the program



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthControlsApiGetAuthcontrolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthControlsApiGetAuthcontrolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardProduct** | **optional.String**| Card product token. Use \&quot;null\&quot; to get auth controls that are not associated with any card product. | 
 **user** | **optional.String**| User token. Use \&quot;null\&quot; to get auth controls that are not associated with any user. | 
 **count** | **optional.Int32**| Number of items to retrieve. Count can be between 1 - 10 items. | [default to 5]
 **startIndex** | **optional.Int32**| Indicates from what row to start returning data. | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Field by which to sort the returned items. Use any field in the model, or system fields lastModifiedTime or createdTime. | [default to -lastModifiedTime]

### Return type

[**AuthControlListResponse**](AuthControlListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthcontrolsExemptmids**
> AuthControlExemptMidsListResponse GetAuthcontrolsExemptmids(ctx, optional)
Lists all auth control exempted MIDs for the program



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthControlsApiGetAuthcontrolsExemptmidsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthControlsApiGetAuthcontrolsExemptmidsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardProduct** | **optional.String**| Card product token. Use \&quot;null\&quot; to get auth controls that are not associated with any card product. | 
 **user** | **optional.String**| User token. Use \&quot;null\&quot; to get auth controls that are not associated with any user. | 
 **count** | **optional.Int32**| Number of items to retrieve. Count can be between 1 - 10 items. | [default to 5]
 **startIndex** | **optional.Int32**| Indicates from what row to start returning data. | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Field by which to sort the returned items. Use any field in the model, or system fields lastModifiedTime or createdTime. | [default to -lastModifiedTime]

### Return type

[**AuthControlExemptMidsListResponse**](AuthControlExemptMidsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthcontrolsExemptmidsToken**
> AuthControlExemptMidsResponse GetAuthcontrolsExemptmidsToken(ctx, token)
Returns a specific auth control exemptmids



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Auth control token | 

### Return type

[**AuthControlExemptMidsResponse**](auth_control_exempt_mids_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthcontrolsToken**
> AuthControlResponse GetAuthcontrolsToken(ctx, token, optional)
Returns a specific auth control exception



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Auth control token | 
 **optional** | ***AuthControlsApiGetAuthcontrolsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthControlsApiGetAuthcontrolsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**AuthControlResponse**](auth_control_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAuthcontrols**
> AuthControlResponse PostAuthcontrols(ctx, body)
Creates an auth control exception



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthControlRequest**](AuthControlRequest.md)| Auth control object | 

### Return type

[**AuthControlResponse**](auth_control_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAuthcontrolsExemptmids**
> AuthControlExemptMidsResponse PostAuthcontrolsExemptmids(ctx, body)
Creates an auth control for exempting MIDs



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthControlExemptMidsRequest**](AuthControlExemptMidsRequest.md)| Auth control exempt MID object | 

### Return type

[**AuthControlExemptMidsResponse**](auth_control_exempt_mids_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAuthcontrolsExemptmidsToken**
> PutAuthcontrolsExemptmidsToken(ctx, token, optional)
Updates the status an auth control exemptmids



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Auth control token | 
 **optional** | ***AuthControlsApiPutAuthcontrolsExemptmidsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthControlsApiPutAuthcontrolsExemptmidsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AuthControlExemptMidsUpdateRequest**](AuthControlExemptMidsUpdateRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAuthcontrolsToken**
> AuthControlResponse PutAuthcontrolsToken(ctx, token, body)
Updates an auth control exception



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Auth control token | 
  **body** | [**AuthControlUpdateRequest**](AuthControlUpdateRequest.md)| Auth control object | 

### Return type

[**AuthControlResponse**](auth_control_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

