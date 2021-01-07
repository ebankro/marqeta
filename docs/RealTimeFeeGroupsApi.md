# \RealTimeFeeGroupsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRealtimefeegroups**](RealTimeFeeGroupsApi.md#GetRealtimefeegroups) | **Get** /realtimefeegroups | Lists all real-time fee groups
[**GetRealtimefeegroupsToken**](RealTimeFeeGroupsApi.md#GetRealtimefeegroupsToken) | **Get** /realtimefeegroups/{token} | Returns a specific real-time fee group
[**PostRealtimefeegroups**](RealTimeFeeGroupsApi.md#PostRealtimefeegroups) | **Post** /realtimefeegroups | Creates a real-time fee group
[**PutRealtimefeegroupsToken**](RealTimeFeeGroupsApi.md#PutRealtimefeegroupsToken) | **Put** /realtimefeegroups/{token} | Updates a specific real-time fee group


# **GetRealtimefeegroups**
> RealTimeFeeGroupListResponse GetRealtimefeegroups(ctx, optional)
Lists all real-time fee groups



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RealTimeFeeGroupsApiGetRealtimefeegroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealTimeFeeGroupsApiGetRealtimefeegroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of real-time fee groups to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

[**RealTimeFeeGroupListResponse**](RealTimeFeeGroupListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRealtimefeegroupsToken**
> RealTimeFeeGroup GetRealtimefeegroupsToken(ctx, token)
Returns a specific real-time fee group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Real-time fee group token | 

### Return type

[**RealTimeFeeGroup**](real_time_fee_group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRealtimefeegroups**
> RealTimeFeeGroup PostRealtimefeegroups(ctx, optional)
Creates a real-time fee group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RealTimeFeeGroupsApiPostRealtimefeegroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealTimeFeeGroupsApiPostRealtimefeegroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RealTimeFeeGroupCreateRequest**](RealTimeFeeGroupCreateRequest.md)|  | 

### Return type

[**RealTimeFeeGroup**](real_time_fee_group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutRealtimefeegroupsToken**
> RealTimeFeeGroup PutRealtimefeegroupsToken(ctx, token, optional)
Updates a specific real-time fee group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Real-time fee group token | 
 **optional** | ***RealTimeFeeGroupsApiPutRealtimefeegroupsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealTimeFeeGroupsApiPutRealtimefeegroupsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RealTimeFeeGroupRequest**](RealTimeFeeGroupRequest.md)|  | 

### Return type

[**RealTimeFeeGroup**](real_time_fee_group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

