# \PushToCardApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPushtocardsDisburse**](PushToCardApi.md#GetPushtocardsDisburse) | **Get** /pushtocards/disburse | Lists all push-to-card disbursements
[**GetPushtocardsDisburseToken**](PushToCardApi.md#GetPushtocardsDisburseToken) | **Get** /pushtocards/disburse/{token} | Returns a specific push-to-card disbursement
[**GetPushtocardsPaymentcard**](PushToCardApi.md#GetPushtocardsPaymentcard) | **Get** /pushtocards/paymentcard | Returns all push-to-card payment card details
[**GetPushtocardsPaymentcardToken**](PushToCardApi.md#GetPushtocardsPaymentcardToken) | **Get** /pushtocards/paymentcard/{token} | Returns a specific paymentcard object
[**PostPushtocardsDisburse**](PushToCardApi.md#PostPushtocardsDisburse) | **Post** /pushtocards/disburse | Initiates a push-to-card money disbursement
[**PostPushtocardsPaymentcard**](PushToCardApi.md#PostPushtocardsPaymentcard) | **Post** /pushtocards/paymentcard | Adds an external card to which funds will be pushed


# **GetPushtocardsDisburse**
> PushToCardDisburseListResponse GetPushtocardsDisburse(ctx, optional)
Lists all push-to-card disbursements



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PushToCardApiGetPushtocardsDisburseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PushToCardApiGetPushtocardsDisburseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of push-to-card disbursements to retrieve | [default to 10]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

[**PushToCardDisburseListResponse**](PushToCardDisburseListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPushtocardsDisburseToken**
> PushToCardDisbursementResponse GetPushtocardsDisburseToken(ctx, token, optional)
Returns a specific push-to-card disbursement



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Push-to-card disbursement token | 
 **optional** | ***PushToCardApiGetPushtocardsDisburseTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PushToCardApiGetPushtocardsDisburseTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**PushToCardDisbursementResponse**](push_to_card_disbursement_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPushtocardsPaymentcard**
> PushToCardListResponse GetPushtocardsPaymentcard(ctx, userToken, optional)
Returns all push-to-card payment card details



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userToken** | **string**| User token | 
 **optional** | ***PushToCardApiGetPushtocardsPaymentcardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PushToCardApiGetPushtocardsPaymentcardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of push-to-card payment cards to retrieve | [default to 10]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

[**PushToCardListResponse**](PushToCardListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPushtocardsPaymentcardToken**
> PushToCardResponse GetPushtocardsPaymentcardToken(ctx, token, optional)
Returns a specific paymentcard object



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Push-to-card token | 
 **optional** | ***PushToCardApiGetPushtocardsPaymentcardTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PushToCardApiGetPushtocardsPaymentcardTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**PushToCardResponse**](push_to_card_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPushtocardsDisburse**
> PushToCardDisbursementResponse PostPushtocardsDisburse(ctx, optional)
Initiates a push-to-card money disbursement



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PushToCardApiPostPushtocardsDisburseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PushToCardApiPostPushtocardsDisburseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PushToCardDisburseRequest**](PushToCardDisburseRequest.md)|  | 

### Return type

[**PushToCardDisbursementResponse**](push_to_card_disbursement_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPushtocardsPaymentcard**
> PushToCardResponse PostPushtocardsPaymentcard(ctx, optional)
Adds an external card to which funds will be pushed



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PushToCardApiPostPushtocardsPaymentcardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PushToCardApiPostPushtocardsPaymentcardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PushToCardRequest**](PushToCardRequest.md)|  | 

### Return type

[**PushToCardResponse**](push_to_card_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

