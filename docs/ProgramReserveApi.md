# \ProgramReserveApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deposit**](ProgramReserveApi.md#Deposit) | **Post** /programreserve/deposits | 
[**GetProgramReserveDeposits**](ProgramReserveApi.md#GetProgramReserveDeposits) | **Get** /programreserve/deposits | 
[**GetProgramreserveBalances**](ProgramReserveApi.md#GetProgramreserveBalances) | **Get** /programreserve/balances | Returns the latest balance in the program reserve account
[**GetProgramreserveTransactions**](ProgramReserveApi.md#GetProgramreserveTransactions) | **Get** /programreserve/transactions | Returns a list of program reserve transactions (credits and debits)
[**PostProgramreserveTransactions**](ProgramReserveApi.md#PostProgramreserveTransactions) | **Post** /programreserve/transactions | Credits or debits the program reserve account


# **Deposit**
> Deposit(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProgramReserveApiDepositOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgramReserveApiDepositOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ProgramReserveDepositRequest**](ProgramReserveDepositRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramReserveDeposits**
> GetProgramReserveDeposits(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProgramReserveApiGetProgramReserveDepositsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgramReserveApiGetProgramReserveDepositsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of items to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramreserveBalances**
> ProgramReserveAccountBalance GetProgramreserveBalances(ctx, )
Returns the latest balance in the program reserve account



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProgramReserveAccountBalance**](program_reserve_account_balance.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramreserveTransactions**
> ProgramReserveTransactionListResponse GetProgramreserveTransactions(ctx, optional)
Returns a list of program reserve transactions (credits and debits)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProgramReserveApiGetProgramreserveTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgramReserveApiGetProgramreserveTransactionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of items to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]
 **body** | **optional.String**| Type | 

### Return type

[**ProgramReserveTransactionListResponse**](ProgramReserveTransactionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostProgramreserveTransactions**
> ProgramReserveTransactionResponse PostProgramreserveTransactions(ctx, optional)
Credits or debits the program reserve account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProgramReserveApiPostProgramreserveTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgramReserveApiPostProgramreserveTransactionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ProgramReserveTransactionRequest**](ProgramReserveTransactionRequest.md)|  | 

### Return type

[**ProgramReserveTransactionResponse**](program_reserve_transaction_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

