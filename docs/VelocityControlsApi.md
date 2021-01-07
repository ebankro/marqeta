# \VelocityControlsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVelocitycontrols**](VelocityControlsApi.md#GetVelocitycontrols) | **Get** /velocitycontrols | Queries velocity controls
[**GetVelocitycontrolsToken**](VelocityControlsApi.md#GetVelocitycontrolsToken) | **Get** /velocitycontrols/{token} | Returns a specific velocity control
[**GetVelocitycontrolsUserUsertokenAvailable**](VelocityControlsApi.md#GetVelocitycontrolsUserUsertokenAvailable) | **Get** /velocitycontrols/user/{user_token}/available | Queries a user&#39;s velocity control balances
[**PostVelocitycontrols**](VelocityControlsApi.md#PostVelocitycontrols) | **Post** /velocitycontrols | Creates a velocity control
[**PutVelocitycontrolsToken**](VelocityControlsApi.md#PutVelocitycontrolsToken) | **Put** /velocitycontrols/{token} | Updates a specific velocity control


# **GetVelocitycontrols**
> VelocityControlListResponse GetVelocitycontrols(ctx, optional)
Queries velocity controls



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VelocityControlsApiGetVelocitycontrolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VelocityControlsApiGetVelocitycontrolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardProduct** | **optional.String**| Card product token. Use \&quot;null\&quot; to get velocity controls that are not associated with any card product. | 
 **user** | **optional.String**| User token. Use \&quot;null\&quot; to get velocity controls that are not associated with any user. | 
 **count** | **optional.Int32**| Number of items to retrieve. Count can be between 1 - 10 items. | [default to 5]
 **startIndex** | **optional.Int32**| Indicates from what row to start returning data. | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Field by which to sort the returned items. Use any field in the model, or system fields lastModifiedTime or createdTime. | [default to -lastModifiedTime]

### Return type

[**VelocityControlListResponse**](VelocityControlListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVelocitycontrolsToken**
> VelocityControlResponse GetVelocitycontrolsToken(ctx, token, optional)
Returns a specific velocity control



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Velocity control token | 
 **optional** | ***VelocityControlsApiGetVelocitycontrolsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VelocityControlsApiGetVelocitycontrolsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**VelocityControlResponse**](velocity_control_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVelocitycontrolsUserUsertokenAvailable**
> VelocityControlBalanceListResponse GetVelocitycontrolsUserUsertokenAvailable(ctx, userToken, optional)
Queries a user's velocity control balances



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userToken** | **string**| User token | 
 **optional** | ***VelocityControlsApiGetVelocitycontrolsUserUsertokenAvailableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VelocityControlsApiGetVelocitycontrolsUserUsertokenAvailableOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of items to retrieve. Count can be between 1 - 10 items. | [default to 5]
 **startIndex** | **optional.Int32**| Indicates from what row to start returning data. | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Field by which to sort the returned items. Use any field in the model, or system fields lastModifiedTime or createdTime. | [default to -lastModifiedTime]
 **forceDto** | **optional.String**|  | 

### Return type

[**VelocityControlBalanceListResponse**](VelocityControlBalanceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVelocitycontrols**
> VelocityControlResponse PostVelocitycontrols(ctx, body)
Creates a velocity control



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VelocityControlRequest**](VelocityControlRequest.md)| Velocity control object | 

### Return type

[**VelocityControlResponse**](velocity_control_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutVelocitycontrolsToken**
> VelocityControlResponse PutVelocitycontrolsToken(ctx, token, body)
Updates a specific velocity control



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Velocity control token | 
  **body** | [**VelocityControlUpdateRequest**](VelocityControlUpdateRequest.md)| Velocity control object | 

### Return type

[**VelocityControlResponse**](velocity_control_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

