# \CardholderBalancesApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalancesToken**](CardholderBalancesApi.md#GetBalancesToken) | **Get** /balances/{token} | Returns account balances for a cardholder
[**GetBalancesTokenMsas**](CardholderBalancesApi.md#GetBalancesTokenMsas) | **Get** /balances/{token}/msas | Returns a merchant-specific account balance
[**UpdateBalance**](CardholderBalancesApi.md#UpdateBalance) | **Put** /balances/{token} | Updates the cached_balance of cardholder


# **GetBalancesToken**
> CardholderBalances GetBalancesToken(ctx, token)
Returns account balances for a cardholder



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| User or Business token | 

### Return type

[**CardholderBalances**](cardholder_balances.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBalancesTokenMsas**
> CardholderMsaBalance GetBalancesTokenMsas(ctx, token, optional)
Returns a merchant-specific account balance



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| User or Business token | 
 **optional** | ***CardholderBalancesApiGetBalancesTokenMsasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardholderBalancesApiGetBalancesTokenMsasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of restrictions to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | 

### Return type

[**CardholderMsaBalance**](cardholder_msa_balance.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBalance**
> CardholderBalances UpdateBalance(ctx, token, body)
Updates the cached_balance of cardholder



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| User or Business token | 
  **body** | [**CardholderUpdateCachedBalances**](CardholderUpdateCachedBalances.md)| Cardholder Update Cache Balance | 

### Return type

[**CardholderBalances**](cardholder_balances.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

