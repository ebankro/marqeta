# \DirectDepositAccountsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](DirectDepositAccountsApi.md#CreateAccount) | **Post** /depositaccounts | Creates new direct deposit account for cardholder.
[**CreateTransition**](DirectDepositAccountsApi.md#CreateTransition) | **Post** /depositaccounts/transitions | Creates new transition for a direct deposit account.
[**GetCDDInfo**](DirectDepositAccountsApi.md#GetCDDInfo) | **Get** /depositaccounts/{token}/cdd | Get direct deposit account transition list for card holder.
[**GetDirectDepositAccount**](DirectDepositAccountsApi.md#GetDirectDepositAccount) | **Get** /depositaccounts/{token} | Get direct deposit account.
[**GetDirectDepositAccountTransition**](DirectDepositAccountsApi.md#GetDirectDepositAccountTransition) | **Get** /depositaccounts/transitions/{token} | Get direct deposit account transition.
[**GetTransitionList**](DirectDepositAccountsApi.md#GetTransitionList) | **Get** /depositaccounts/{user_token}/transitions | Get direct deposit account transition list for card holder.
[**GetUserDirectDepositAccounts**](DirectDepositAccountsApi.md#GetUserDirectDepositAccounts) | **Get** /depositaccounts/user/{token} | List all specific direct deposit accounts.
[**GetUserForDirectDepositAccount**](DirectDepositAccountsApi.md#GetUserForDirectDepositAccount) | **Get** /depositaccounts/account/{account_number}/user | Get User for Plain Text Account Number
[**Update**](DirectDepositAccountsApi.md#Update) | **Put** /depositaccounts/{token} | Update direct deposit account.
[**UpdateCDDInfo**](DirectDepositAccountsApi.md#UpdateCDDInfo) | **Put** /depositaccounts/{token}/cdd/{cddtoken} | Update CDD answers for Direct Deposit Account


# **CreateAccount**
> DirectDepositAccountResponse CreateAccount(ctx, body)
Creates new direct deposit account for cardholder.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DirectDepositAccountRequest**](DirectDepositAccountRequest.md)| Create direct deposit account for cardholder | 

### Return type

[**DirectDepositAccountResponse**](direct_deposit_account_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTransition**
> DirectDepositAccountTransitionResponse CreateTransition(ctx, body)
Creates new transition for a direct deposit account.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DirectDepositAccountTransitionRequest**](DirectDepositAccountTransitionRequest.md)| Create transition for direct deposit account | 

### Return type

[**DirectDepositAccountTransitionResponse**](direct_deposit_account_transition_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCDDInfo**
> CustomerDueDiligenceResponse GetCDDInfo(ctx, token)
Get direct deposit account transition list for card holder.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Get CDD info for a specific DDA token | 

### Return type

[**CustomerDueDiligenceResponse**](customer_due_diligence_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDirectDepositAccount**
> DirectDepositAccountResponse GetDirectDepositAccount(ctx, token)
Get direct deposit account.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Get specific direct deposit account | 

### Return type

[**DirectDepositAccountResponse**](direct_deposit_account_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDirectDepositAccountTransition**
> DirectDepositAccountTransitionResponse GetDirectDepositAccountTransition(ctx, token)
Get direct deposit account transition.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Get specific direct deposit account transition | 

### Return type

[**DirectDepositAccountTransitionResponse**](direct_deposit_account_transition_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransitionList**
> DirectDepositAccountTransitionResponse GetTransitionList(ctx, userToken, optional)
Get direct deposit account transition list for card holder.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userToken** | **string**| Get direct deposit account transition list for user | 
 **optional** | ***DirectDepositAccountsApiGetTransitionListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DirectDepositAccountsApiGetTransitionListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of users to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

[**DirectDepositAccountTransitionResponse**](direct_deposit_account_transition_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserDirectDepositAccounts**
> DirectDepositAccountListResponse GetUserDirectDepositAccounts(ctx, token, optional)
List all specific direct deposit accounts.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Get specific direct deposit account | 
 **optional** | ***DirectDepositAccountsApiGetUserDirectDepositAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DirectDepositAccountsApiGetUserDirectDepositAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of users to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]
 **state** | **optional.String**| Direct deposit account status | 

### Return type

[**DirectDepositAccountListResponse**](DirectDepositAccountListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserForDirectDepositAccount**
> DirectDepositAccountResponse GetUserForDirectDepositAccount(ctx, accountNumber)
Get User for Plain Text Account Number



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountNumber** | **string**| Get user associated with direct deposit account number | 

### Return type

[**DirectDepositAccountResponse**](direct_deposit_account_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> DirectDepositAccountResponse Update(ctx, body, token)
Update direct deposit account.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DepositAccountUpdateRequest**](DepositAccountUpdateRequest.md)| Update direct deposit account | 
  **token** | **string**|  | 

### Return type

[**DirectDepositAccountResponse**](direct_deposit_account_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCDDInfo**
> CustomerDueDiligenceResponse UpdateCDDInfo(ctx, body, token, cddtoken)
Update CDD answers for Direct Deposit Account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomerDueDiligenceUpdateResponse**](CustomerDueDiligenceUpdateResponse.md)| Update CDD answers | 
  **token** | **string**|  | 
  **cddtoken** | **string**|  | 

### Return type

[**CustomerDueDiligenceResponse**](customer_due_diligence_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

