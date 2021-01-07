# \CommandoModesApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommandomodes**](CommandoModesApi.md#GetCommandomodes) | **Get** /commandomodes | Lists all commando mode control sets
[**GetCommandomodesCommandomodetokenTransitions**](CommandoModesApi.md#GetCommandomodesCommandomodetokenTransitions) | **Get** /commandomodes/{commandomode_token}/transitions | Lists all commando mode transitions related to a commando mode control set
[**GetCommandomodesToken**](CommandoModesApi.md#GetCommandomodesToken) | **Get** /commandomodes/{token} | Returns a specific commando mode control set
[**GetCommandomodesTransitionsToken**](CommandoModesApi.md#GetCommandomodesTransitionsToken) | **Get** /commandomodes/transitions/{token} | Returns a specific commando mode transition


# **GetCommandomodes**
> CommandoModeListResponse GetCommandomodes(ctx, optional)
Lists all commando mode control sets



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommandoModesApiGetCommandomodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommandoModesApiGetCommandomodesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of commando modes to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**CommandoModeListResponse**](CommandoModeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommandomodesCommandomodetokenTransitions**
> CommandoModeTransitionListResponse GetCommandomodesCommandomodetokenTransitions(ctx, commandomodeToken, optional)
Lists all commando mode transitions related to a commando mode control set



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commandomodeToken** | **string**| Commando mode token | 
 **optional** | ***CommandoModesApiGetCommandomodesCommandomodetokenTransitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommandoModesApiGetCommandomodesCommandomodetokenTransitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of transitions to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

[**CommandoModeTransitionListResponse**](CommandoModeTransitionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommandomodesToken**
> CommandoModeResponse GetCommandomodesToken(ctx, token)
Returns a specific commando mode control set



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**CommandoModeResponse**](commando_mode_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommandomodesTransitionsToken**
> CommandoModeTransitionResponse GetCommandomodesTransitionsToken(ctx, token)
Returns a specific commando mode transition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Commando mode transition token | 

### Return type

[**CommandoModeTransitionResponse**](commando_mode_transition_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

