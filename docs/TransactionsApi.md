# \TransactionsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransactions**](TransactionsApi.md#GetTransactions) | **Get** /transactions | Returns transactions
[**GetTransactionsFundingsourceFundingsourcetoken**](TransactionsApi.md#GetTransactionsFundingsourceFundingsourcetoken) | **Get** /transactions/fundingsource/{funding_source_token} | Returns transactions for a specific funding account
[**GetTransactionsToken**](TransactionsApi.md#GetTransactionsToken) | **Get** /transactions/{token} | Returns a transaction
[**GetTransactionsTokenRelated**](TransactionsApi.md#GetTransactionsTokenRelated) | **Get** /transactions/{token}/related | Returns related transactions


# **GetTransactions**
> TransactionModelListResponse GetTransactions(ctx, optional)
Returns transactions



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransactionsApiGetTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionsApiGetTransactionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of transactions to retrieve | [default to 10]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -user_transaction_time]
 **startDate** | **optional.String**| Start date (yyyy-MM-dd | yyyy-MM-ddTHH:mm:ss.SS) | 
 **endDate** | **optional.String**| End date (yyyy-MM-dd | yyyy-MM-ddTHH:mm:ss.SS) | 
 **type_** | **optional.String**| Comma-delimited list of transaction types to include | 
 **userToken** | **optional.String**| User token | 
 **businessToken** | **optional.String**| Business token | 
 **actingUserToken** | **optional.String**| Acting user token | 
 **cardToken** | **optional.String**| Card token | 
 **merchantToken** | **optional.String**| Merchant token | 
 **campaignToken** | **optional.String**| Campaign token | 
 **state** | **optional.String**| Comma-delimited list of transaction states to display e.g. PENDING | CLEARED | COMPLETION | DECLINED | ERROR | ALL | [default to PENDING,COMPLETION]
 **version** | **optional.String**|  | 
 **verbose** | **optional.Bool**|  | 

### Return type

[**TransactionModelListResponse**](TransactionModelListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionsFundingsourceFundingsourcetoken**
> TransactionModelListResponse GetTransactionsFundingsourceFundingsourcetoken(ctx, fundingSourceToken, optional)
Returns transactions for a specific funding account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fundingSourceToken** | **string**| Funding account token | 
 **optional** | ***TransactionsApiGetTransactionsFundingsourceFundingsourcetokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionsApiGetTransactionsFundingsourceFundingsourcetokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of transactions to retrieve | [default to 10]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -user_transaction_time]
 **startDate** | **optional.String**| Start date (yyyy-MM-dd) | 
 **endDate** | **optional.String**| End date (yyyy-MM-dd) | 
 **type_** | **optional.String**| Comma-delimited list of transaction types to include | 
 **polarity** | **optional.String**| Type of transactions to retrieve: CREDIT or DEBIT | 
 **version** | **optional.String**|  | 
 **verbose** | **optional.Bool**|  | 

### Return type

[**TransactionModelListResponse**](TransactionModelListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionsToken**
> TransactionModel GetTransactionsToken(ctx, token, optional)
Returns a transaction



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Transaction token | 
 **optional** | ***TransactionsApiGetTransactionsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionsApiGetTransactionsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **version** | **optional.String**|  | 
 **verbose** | **optional.Bool**|  | 

### Return type

[**TransactionModel**](transaction_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionsTokenRelated**
> TransactionModelListResponse GetTransactionsTokenRelated(ctx, token, optional)
Returns related transactions



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Transaction token | 
 **optional** | ***TransactionsApiGetTransactionsTokenRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionsApiGetTransactionsTokenRelatedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of transactions to retrieve | [default to 10]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -user_transaction_time]
 **startDate** | **optional.String**| Start date (yyyy-MM-dd | yyyy-MM-ddTHH:mm:ss.SS) | 
 **endDate** | **optional.String**| End date (yyyy-MM-dd | yyyy-MM-ddTHH:mm:ss.SS) | 
 **type_** | **optional.String**| Comma-delimited list of transaction types to include | 
 **state** | **optional.String**| Comma-delimited list of transaction states to display e.g. PENDING | CLEARED | COMPLETION | ALL | [default to ALL]
 **version** | **optional.String**|  | 
 **verbose** | **optional.Bool**|  | 

### Return type

[**TransactionModelListResponse**](TransactionModelListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

