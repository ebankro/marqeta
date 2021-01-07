# \AccountHolderGroupsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountholdergroups**](AccountHolderGroupsApi.md#GetAccountholdergroups) | **Get** /accountholdergroups | Lists account holder groups
[**GetAccountholdergroupsToken**](AccountHolderGroupsApi.md#GetAccountholdergroupsToken) | **Get** /accountholdergroups/{token} | Returns a specific account holder group object
[**PostAccountholdergroups**](AccountHolderGroupsApi.md#PostAccountholdergroups) | **Post** /accountholdergroups | Creates an account holder group object
[**PutAccountholdergroupsToken**](AccountHolderGroupsApi.md#PutAccountholdergroupsToken) | **Put** /accountholdergroups/{token} | Updates an account holder group object


# **GetAccountholdergroups**
> AccountHolderGroupListResponse GetAccountholdergroups(ctx, optional)
Lists account holder groups



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountHolderGroupsApiGetAccountholdergroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountHolderGroupsApiGetAccountholdergroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of items to retrieve. Count can be between 1 - 10 items. | [default to 10]
 **startIndex** | **optional.Int32**| Indicates from what row to start returning data. | [default to 0]
 **sortBy** | **optional.String**| Field by which to sort the returned items. Use any field in the model, or system fields lastModifiedTime or createdTime. | [default to -lastModifiedTime]

### Return type

[**AccountHolderGroupListResponse**](AccountHolderGroupListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountholdergroupsToken**
> AccountHolderGroupResponse GetAccountholdergroupsToken(ctx, token)
Returns a specific account holder group object



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Account holder group token | 

### Return type

[**AccountHolderGroupResponse**](account_holder_group_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAccountholdergroups**
> AccountHolderGroupResponse PostAccountholdergroups(ctx, body)
Creates an account holder group object



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountHolderGroupRequest**](AccountHolderGroupRequest.md)| Account holder group object | 

### Return type

[**AccountHolderGroupResponse**](account_holder_group_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAccountholdergroupsToken**
> AccountHolderGroupResponse PutAccountholdergroupsToken(ctx, body, token)
Updates an account holder group object



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountHolderGroupUpdateRequest**](AccountHolderGroupUpdateRequest.md)| Account holder group update object | 
  **token** | **string**|  | 

### Return type

[**AccountHolderGroupResponse**](account_holder_group_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

