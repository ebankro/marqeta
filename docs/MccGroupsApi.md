# \MccGroupsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMccgroups**](MccGroupsApi.md#GetMccgroups) | **Get** /mccgroups | Lists all MCC groups
[**GetMccgroupsToken**](MccGroupsApi.md#GetMccgroupsToken) | **Get** /mccgroups/{token} | Returns a specific MCC group
[**PostMccgroups**](MccGroupsApi.md#PostMccgroups) | **Post** /mccgroups | Creates an MCC group
[**PutMccgroupsToken**](MccGroupsApi.md#PutMccgroupsToken) | **Put** /mccgroups/{token} | Updates an MCC group


# **GetMccgroups**
> MccGroupListResponse GetMccgroups(ctx, optional)
Lists all MCC groups



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MccGroupsApiGetMccgroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MccGroupsApiGetMccgroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mcc** | **optional.String**| MCC | 
 **count** | **optional.Int32**| Number of items to retrieve. Count can be between 1 - 10 items. | [default to 10]
 **startIndex** | **optional.Int32**| Indicates from what row to start returning data. | [default to 0]
 **sortBy** | **optional.String**| Field by which to sort the returned items. Use any field in the model, or system fields lastModifiedTime or createdTime. | [default to -lastModifiedTime]

### Return type

[**MccGroupListResponse**](MCCGroupListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMccgroupsToken**
> MccGroupModel GetMccgroupsToken(ctx, token)
Returns a specific MCC group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| MCC group token | 

### Return type

[**MccGroupModel**](mcc_group_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMccgroups**
> MccGroupModel PostMccgroups(ctx, body)
Creates an MCC group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MccGroupModel**](MccGroupModel.md)| MCC group | 

### Return type

[**MccGroupModel**](mcc_group_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMccgroupsToken**
> MccGroupUpdateModel PutMccgroupsToken(ctx, body, token)
Updates an MCC group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MccGroupUpdateModel**](MccGroupUpdateModel.md)| MCC group | 
  **token** | **string**|  | 

### Return type

[**MccGroupUpdateModel**](mcc_group_update_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

