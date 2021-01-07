# \BusinessTransitionsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBusinesstransitionsBusinessBusinesstoken**](BusinessTransitionsApi.md#GetBusinesstransitionsBusinessBusinesstoken) | **Get** /businesstransitions/business/{business_token} | Returns transitions for a given business
[**GetBusinesstransitionsToken**](BusinessTransitionsApi.md#GetBusinesstransitionsToken) | **Get** /businesstransitions/{token} | Returns a business transition
[**PostBusinesstransitions**](BusinessTransitionsApi.md#PostBusinesstransitions) | **Post** /businesstransitions | Creates a business transition


# **GetBusinesstransitionsBusinessBusinesstoken**
> BusinessTransitionListResponse GetBusinesstransitionsBusinessBusinesstoken(ctx, businessToken, optional)
Returns transitions for a given business



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **businessToken** | **string**| Business token | 
 **optional** | ***BusinessTransitionsApiGetBusinesstransitionsBusinessBusinesstokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessTransitionsApiGetBusinesstransitionsBusinessBusinesstokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of business transitions to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -id]

### Return type

[**BusinessTransitionListResponse**](BusinessTransitionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBusinesstransitionsToken**
> BusinessTransitionResponse GetBusinesstransitionsToken(ctx, token, optional)
Returns a business transition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Transition token | 
 **optional** | ***BusinessTransitionsApiGetBusinesstransitionsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessTransitionsApiGetBusinesstransitionsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**BusinessTransitionResponse**](BusinessTransitionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBusinesstransitions**
> BusinessTransitionResponse PostBusinesstransitions(ctx, optional)
Creates a business transition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BusinessTransitionsApiPostBusinesstransitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessTransitionsApiPostBusinesstransitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BusinessTransitionRequest**](BusinessTransitionRequest.md)|  | 

### Return type

[**BusinessTransitionResponse**](BusinessTransitionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

