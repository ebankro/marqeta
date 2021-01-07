# \DirectDepositsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDirectdeposits**](DirectDepositsApi.md#GetDirectdeposits) | **Get** /directdeposits | Lists all direct deposits
[**GetDirectdepositsAccountsUserorbusinesstoken**](DirectDepositsApi.md#GetDirectdepositsAccountsUserorbusinesstoken) | **Get** /directdeposits/accounts/{user_or_business_token} | Returns an account and routing number which can be used for direct deposit
[**GetDirectdepositsToken**](DirectDepositsApi.md#GetDirectdepositsToken) | **Get** /directdeposits/{token} | Returns a direct deposit entry
[**GetDirectdepositsTransitions**](DirectDepositsApi.md#GetDirectdepositsTransitions) | **Get** /directdeposits/transitions | Returns a list of direct deposit transitions
[**GetDirectdepositsTransitionsToken**](DirectDepositsApi.md#GetDirectdepositsTransitionsToken) | **Get** /directdeposits/transitions/{token} | Returns a direct deposit transition
[**PostDirectdepositsTransitions**](DirectDepositsApi.md#PostDirectdepositsTransitions) | **Post** /directdeposits/transitions | Creates a direct deposit transition
[**PutDirectdepositsAccountsUserorbusinesstoken**](DirectDepositsApi.md#PutDirectdepositsAccountsUserorbusinesstoken) | **Put** /directdeposits/accounts/{user_or_business_token} | Updates a specific direct deposit account


# **GetDirectdeposits**
> DirectDepositListResponse GetDirectdeposits(ctx, optional)
Lists all direct deposits



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectDepositsApiGetDirectdepositsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DirectDepositsApiGetDirectdepositsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of direct deposits to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **reversedAfterGracePeriod** | **optional.Bool**| Reversed after grace period | 
 **userToken** | **optional.String**| User token | 
 **businessToken** | **optional.String**| Business token | 
 **directDepositState** | **optional.String**| Direct deposit state | 
 **startSettlementDate** | **optional.String**| Start settlement date | 
 **endSettlementDate** | **optional.String**| End settlement date | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**DirectDepositListResponse**](DirectDepositListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDirectdepositsAccountsUserorbusinesstoken**
> DepositAccount GetDirectdepositsAccountsUserorbusinesstoken(ctx, userOrBusinessToken)
Returns an account and routing number which can be used for direct deposit



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userOrBusinessToken** | **string**|  | 

### Return type

[**DepositAccount**](deposit_account.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDirectdepositsToken**
> DepositDepositResponse GetDirectdepositsToken(ctx, token)
Returns a direct deposit entry



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**DepositDepositResponse**](DepositDepositResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDirectdepositsTransitions**
> DirectDepositTransitionListResponse GetDirectdepositsTransitions(ctx, optional)
Returns a list of direct deposit transitions



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectDepositsApiGetDirectdepositsTransitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DirectDepositsApiGetDirectdepositsTransitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of direct deposit transitions to retrieve | [default to 5]
 **userToken** | **optional.String**| User token | 
 **businessToken** | **optional.String**| Business token | 
 **directDepositToken** | **optional.String**| Direct deposit token | 
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]
 **states** | **optional.String**| Comma-delimited list of direct deposit states to display e.g. PENDING | REVERSED | APPLIED | REJECTED  | 

### Return type

[**DirectDepositTransitionListResponse**](DirectDepositTransitionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDirectdepositsTransitionsToken**
> DirectDepositTransitionResponse GetDirectdepositsTransitionsToken(ctx, token)
Returns a direct deposit transition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**DirectDepositTransitionResponse**](DirectDepositTransitionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDirectdepositsTransitions**
> DirectDepositTransitionResponse PostDirectdepositsTransitions(ctx, optional)
Creates a direct deposit transition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectDepositsApiPostDirectdepositsTransitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DirectDepositsApiPostDirectdepositsTransitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DirectDepositTransitionRequest**](DirectDepositTransitionRequest.md)|  | 

### Return type

[**DirectDepositTransitionResponse**](DirectDepositTransitionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDirectdepositsAccountsUserorbusinesstoken**
> DepositAccount PutDirectdepositsAccountsUserorbusinesstoken(ctx, userOrBusinessToken, body)
Updates a specific direct deposit account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userOrBusinessToken** | **string**| User or business token | 
  **body** | [**DepositAccountUpdateRequest**](DepositAccountUpdateRequest.md)| Deposit account update request | 

### Return type

[**DepositAccount**](deposit_account.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

