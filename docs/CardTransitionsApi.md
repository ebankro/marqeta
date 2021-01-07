# \CardTransitionsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCardtransitionsCardToken**](CardTransitionsApi.md#GetCardtransitionsCardToken) | **Get** /cardtransitions/card/{token} | Lists all card transitions
[**GetCardtransitionsToken**](CardTransitionsApi.md#GetCardtransitionsToken) | **Get** /cardtransitions/{token} | Returns a card transition object
[**PostCardtransitions**](CardTransitionsApi.md#PostCardtransitions) | **Post** /cardtransitions | Creates a card transition object


# **GetCardtransitionsCardToken**
> CardTransitionListResponse GetCardtransitionsCardToken(ctx, token, optional)
Lists all card transitions



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Card token | 
 **optional** | ***CardTransitionsApiGetCardtransitionsCardTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardTransitionsApiGetCardtransitionsCardTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of card transitions to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

[**CardTransitionListResponse**](CardTransitionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardtransitionsToken**
> CardTransitionResponse GetCardtransitionsToken(ctx, token, optional)
Returns a card transition object



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Card transition token | 
 **optional** | ***CardTransitionsApiGetCardtransitionsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardTransitionsApiGetCardtransitionsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**CardTransitionResponse**](card_transition_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCardtransitions**
> CardTransitionResponse PostCardtransitions(ctx, optional)
Creates a card transition object



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CardTransitionsApiPostCardtransitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardTransitionsApiPostCardtransitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CardTransitionRequest**](CardTransitionRequest.md)|  | 

### Return type

[**CardTransitionResponse**](card_transition_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

