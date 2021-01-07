# \DigitalWalletTokensApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDigitalwallettokens**](DigitalWalletTokensApi.md#GetDigitalwallettokens) | **Get** /digitalwallettokens | Returns digital wallet tokens
[**GetDigitalwallettokensCardCardtoken**](DigitalWalletTokensApi.md#GetDigitalwallettokensCardCardtoken) | **Get** /digitalwallettokens/card/{card_token} | Returns a list of digital wallet tokens for the specified card
[**GetDigitalwallettokensToken**](DigitalWalletTokensApi.md#GetDigitalwallettokensToken) | **Get** /digitalwallettokens/{token} | Returns a specific digital wallet token
[**GetDigitalwallettokensTokenShowtokenpan**](DigitalWalletTokensApi.md#GetDigitalwallettokensTokenShowtokenpan) | **Get** /digitalwallettokens/{token}/showtokenpan | Returns a specific digital wallet token PAN visible


# **GetDigitalwallettokens**
> DigitalWalletTokenListResponse GetDigitalwallettokens(ctx, optional)
Returns digital wallet tokens



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DigitalWalletTokensApiGetDigitalwallettokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DigitalWalletTokensApiGetDigitalwallettokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of digital wallet tokens to retrieve | [default to 10]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]
 **startDate** | **optional.String**| Start date (yyyy-MM-dd) | 
 **endDate** | **optional.String**| End date (yyyy-MM-dd) | 
 **panReferenceId** | **optional.String**| PAN reference ID | 
 **tokenReferenceId** | **optional.String**| Token reference ID | 
 **correlationId** | **optional.String**| Correlation ID | 
 **tokenType** | **optional.String**| Comma-delimited list of digital wallet token types to display e.g. DEVICE_SECURE_ELEMENT | MERCHANT_CARD_ON_FILE | DEVICE_CLOUD_BASED | ECOMMERCE_DIGITAL_WALLET | PSEUDO_ACCOUNT | 
 **tokenRequestorName** | **optional.String**| Comma-delimited list of digital wallet token wallet providers to display e.g. APPLE_PAY | ANDROID_PAY| SAMSUNG_PAY | MICROSOFT_PAY | VISA_CHECKOUT | FACEBOOK | NETFLIX | UNKNOWN | 
 **state** | **optional.String**| Comma-delimited list of digital wallet token states to display e.g. REQUESTED | REQUEST_DECLINED | TERMINATED | SUSPENDED | ACTIVE | 
 **embed** | **optional.String**| Embed | 

### Return type

[**DigitalWalletTokenListResponse**](DigitalWalletTokenListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDigitalwallettokensCardCardtoken**
> DigitalWalletTokenListResponse GetDigitalwallettokensCardCardtoken(ctx, cardToken, optional)
Returns a list of digital wallet tokens for the specified card



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardToken** | **string**| Card token | 
 **optional** | ***DigitalWalletTokensApiGetDigitalwallettokensCardCardtokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DigitalWalletTokensApiGetDigitalwallettokensCardCardtokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of items to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

[**DigitalWalletTokenListResponse**](DigitalWalletTokenListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDigitalwallettokensToken**
> DigitalWalletToken GetDigitalwallettokensToken(ctx, token)
Returns a specific digital wallet token



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Digital wallet token | 

### Return type

[**DigitalWalletToken**](digital_wallet_token.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDigitalwallettokensTokenShowtokenpan**
> DigitalWalletToken GetDigitalwallettokensTokenShowtokenpan(ctx, token)
Returns a specific digital wallet token PAN visible



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Digital wallet token | 

### Return type

[**DigitalWalletToken**](digital_wallet_token.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

