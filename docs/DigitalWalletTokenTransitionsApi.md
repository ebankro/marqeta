# \DigitalWalletTokenTransitionsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDigitalwallettokentransitionsDigitalwallettokenToken**](DigitalWalletTokenTransitionsApi.md#GetDigitalwallettokentransitionsDigitalwallettokenToken) | **Get** /digitalwallettokentransitions/digitalwallettoken/{token} | Lists all digital wallet token transitions
[**GetDigitalwallettokentransitionsToken**](DigitalWalletTokenTransitionsApi.md#GetDigitalwallettokentransitionsToken) | **Get** /digitalwallettokentransitions/{token} | Returns a digital wallet transition object
[**PostDigitalwallettokentransitions**](DigitalWalletTokenTransitionsApi.md#PostDigitalwallettokentransitions) | **Post** /digitalwallettokentransitions | Creates a digital wallet token transition


# **GetDigitalwallettokentransitionsDigitalwallettokenToken**
> DigitalWalletTokenTransitionListResponse GetDigitalwallettokentransitionsDigitalwallettokenToken(ctx, token, optional)
Lists all digital wallet token transitions



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Digital wallet token | 
 **optional** | ***DigitalWalletTokenTransitionsApiGetDigitalwallettokentransitionsDigitalwallettokenTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DigitalWalletTokenTransitionsApiGetDigitalwallettokentransitionsDigitalwallettokenTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of digital wallet transitions to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -id]

### Return type

[**DigitalWalletTokenTransitionListResponse**](DigitalWalletTokenTransitionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDigitalwallettokentransitionsToken**
> DigitalWalletTokenTransitionResponse GetDigitalwallettokentransitionsToken(ctx, token, optional)
Returns a digital wallet transition object



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Digital wallet transition token | 
 **optional** | ***DigitalWalletTokenTransitionsApiGetDigitalwallettokentransitionsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DigitalWalletTokenTransitionsApiGetDigitalwallettokentransitionsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**DigitalWalletTokenTransitionResponse**](digital_wallet_token_transition_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDigitalwallettokentransitions**
> DigitalWalletTokenTransitionResponse PostDigitalwallettokentransitions(ctx, optional)
Creates a digital wallet token transition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DigitalWalletTokenTransitionsApiPostDigitalwallettokentransitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DigitalWalletTokenTransitionsApiPostDigitalwallettokentransitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DigitalWalletTokenTransitionRequest**](DigitalWalletTokenTransitionRequest.md)|  | 

### Return type

[**DigitalWalletTokenTransitionResponse**](digital_wallet_token_transition_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

