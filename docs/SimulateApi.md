# \SimulateApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSimulateAdvancedClearing**](SimulateApi.md#PostSimulateAdvancedClearing) | **Post** /simulate/advanced/clearing | Simulates an advanced clearing/settlement transaction
[**PostSimulateAuthorization**](SimulateApi.md#PostSimulateAuthorization) | **Post** /simulate/authorization | Simulates an authorization
[**PostSimulateAuthorizationAdvice**](SimulateApi.md#PostSimulateAuthorizationAdvice) | **Post** /simulate/authorization/advice | Simulates an authorization advice transaction
[**PostSimulateClearing**](SimulateApi.md#PostSimulateClearing) | **Post** /simulate/clearing | Simulates a clearing/settlement transaction
[**PostSimulateDirectdeposits**](SimulateApi.md#PostSimulateDirectdeposits) | **Post** /simulate/directdeposits | Simulates the creation of direct deposit
[**PostSimulateFinancial**](SimulateApi.md#PostSimulateFinancial) | **Post** /simulate/financial | Simulates a financial request (PIN debit) transaction with optional cash back
[**PostSimulateFinancialAdvice**](SimulateApi.md#PostSimulateFinancialAdvice) | **Post** /simulate/financial/advice | Simulates a financial advice transaction
[**PostSimulateFinancialBalanceinquiry**](SimulateApi.md#PostSimulateFinancialBalanceinquiry) | **Post** /simulate/financial/balanceinquiry | Simulates a balance inquiry
[**PostSimulateFinancialOriginalcredit**](SimulateApi.md#PostSimulateFinancialOriginalcredit) | **Post** /simulate/financial/originalcredit | Simulates an orignal credit transaction
[**PostSimulateFinancialWithdrawal**](SimulateApi.md#PostSimulateFinancialWithdrawal) | **Post** /simulate/financial/withdrawal | Simulates an ATM withdrawal transaction
[**PostSimulateReversal**](SimulateApi.md#PostSimulateReversal) | **Post** /simulate/reversal | Simulates a reversal transaction


# **PostSimulateAdvancedClearing**
> AdvancedSimulationResponseModel PostSimulateAdvancedClearing(ctx, optional)
Simulates an advanced clearing/settlement transaction



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SimulateApiPostSimulateAdvancedClearingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SimulateApiPostSimulateAdvancedClearingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AdvancedClearingRequestModel**](AdvancedClearingRequestModel.md)|  | 

### Return type

[**AdvancedSimulationResponseModel**](advanced_simulation_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSimulateAuthorization**
> SimulationResponseModel PostSimulateAuthorization(ctx, optional)
Simulates an authorization



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SimulateApiPostSimulateAuthorizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SimulateApiPostSimulateAuthorizationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AuthRequestModel**](AuthRequestModel.md)|  | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSimulateAuthorizationAdvice**
> SimulationResponseModel PostSimulateAuthorizationAdvice(ctx, optional)
Simulates an authorization advice transaction



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SimulateApiPostSimulateAuthorizationAdviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SimulateApiPostSimulateAuthorizationAdviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AuthorizationAdviceModel**](AuthorizationAdviceModel.md)|  | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSimulateClearing**
> SimulationResponseModel PostSimulateClearing(ctx, optional)
Simulates a clearing/settlement transaction



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SimulateApiPostSimulateClearingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SimulateApiPostSimulateClearingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ClearingModel**](ClearingModel.md)|  | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSimulateDirectdeposits**
> DepositDepositResponse PostSimulateDirectdeposits(ctx, body)
Simulates the creation of direct deposit



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DirectDepositRequest**](DirectDepositRequest.md)| Direct deposit simulate request model | 

### Return type

[**DepositDepositResponse**](DepositDepositResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSimulateFinancial**
> SimulationResponseModel PostSimulateFinancial(ctx, body)
Simulates a financial request (PIN debit) transaction with optional cash back



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FinancialRequestModel**](FinancialRequestModel.md)| Financial request model | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSimulateFinancialAdvice**
> SimulationResponseModel PostSimulateFinancialAdvice(ctx, body)
Simulates a financial advice transaction



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthorizationAdviceModel**](AuthorizationAdviceModel.md)| Financial advice request model | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSimulateFinancialBalanceinquiry**
> SimulationResponseModel PostSimulateFinancialBalanceinquiry(ctx, body)
Simulates a balance inquiry



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BalanceInquiryRequestModel**](BalanceInquiryRequestModel.md)| Balance inquiry request model | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSimulateFinancialOriginalcredit**
> SimulationResponseModel PostSimulateFinancialOriginalcredit(ctx, body)
Simulates an orignal credit transaction



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrignalcreditRequestModel**](OrignalcreditRequestModel.md)| Orignal Credit request model | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSimulateFinancialWithdrawal**
> SimulationResponseModel PostSimulateFinancialWithdrawal(ctx, body)
Simulates an ATM withdrawal transaction



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WithdrawalRequestModel**](WithdrawalRequestModel.md)| ATM withdrawal request model | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSimulateReversal**
> SimulationResponseModel PostSimulateReversal(ctx, optional)
Simulates a reversal transaction



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SimulateApiPostSimulateReversalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SimulateApiPostSimulateReversalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ReversalModel**](ReversalModel.md)|  | 

### Return type

[**SimulationResponseModel**](simulation_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

