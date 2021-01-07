# \UserTransitionsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsertransitionsToken**](UserTransitionsApi.md#GetUsertransitionsToken) | **Get** /usertransitions/{token} | Returns a user transition
[**GetUsertransitionsUserUsertoken**](UserTransitionsApi.md#GetUsertransitionsUserUsertoken) | **Get** /usertransitions/user/{user_token} | Returns transitions for a specific user
[**PostUsertransitions**](UserTransitionsApi.md#PostUsertransitions) | **Post** /usertransitions | Creates a user transition


# **GetUsertransitionsToken**
> UserTransitionResponse GetUsertransitionsToken(ctx, token, optional)
Returns a user transition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Transition token | 
 **optional** | ***UserTransitionsApiGetUsertransitionsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserTransitionsApiGetUsertransitionsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**UserTransitionResponse**](UserTransitionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsertransitionsUserUsertoken**
> UserTransitionListResponse GetUsertransitionsUserUsertoken(ctx, userToken, optional)
Returns transitions for a specific user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userToken** | **string**| User token | 
 **optional** | ***UserTransitionsApiGetUsertransitionsUserUsertokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserTransitionsApiGetUsertransitionsUserUsertokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of user transitions to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -id]

### Return type

[**UserTransitionListResponse**](UserTransitionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsertransitions**
> UserTransitionResponse PostUsertransitions(ctx, optional)
Creates a user transition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserTransitionsApiPostUsertransitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserTransitionsApiPostUsertransitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UserTransitionRequest**](UserTransitionRequest.md)|  | 

### Return type

[**UserTransitionResponse**](UserTransitionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

