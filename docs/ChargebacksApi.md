# \ChargebacksApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChargebacks**](ChargebacksApi.md#GetChargebacks) | **Get** /chargebacks | List all chargebacks
[**GetChargebacksChargebacktokenTransitions**](ChargebacksApi.md#GetChargebacksChargebacktokenTransitions) | **Get** /chargebacks/{chargeback_token}/transitions | Lists all chargeback transitions that are related to a chargeback
[**GetChargebacksToken**](ChargebacksApi.md#GetChargebacksToken) | **Get** /chargebacks/{token} | Returns a specific chargeback
[**GetChargebacksTransitionsToken**](ChargebacksApi.md#GetChargebacksTransitionsToken) | **Get** /chargebacks/transitions/{token} | Returns a specific chargeback transition
[**PostChargebackAllocationAcknowledgment**](ChargebacksApi.md#PostChargebackAllocationAcknowledgment) | **Post** /chargebacks/allocationacknowledgement | Creates a chargeback allocation acknowledgement
[**PostChargebacks**](ChargebacksApi.md#PostChargebacks) | **Post** /chargebacks | Creates a chargeback
[**PostChargebacksTransitions**](ChargebacksApi.md#PostChargebacksTransitions) | **Post** /chargebacks/transitions | Creates a chargeback transition
[**PutChargebacksToken**](ChargebacksApi.md#PutChargebacksToken) | **Put** /chargebacks/{token} | Updates chargeback data
[**PutChargebacksTokenGrantprovisionalcredit**](ChargebacksApi.md#PutChargebacksTokenGrantprovisionalcredit) | **Put** /chargebacks/{token}/grantprovisionalcredit | Grants provisional credit
[**PutChargebacksTokenReverseprovisionalcredit**](ChargebacksApi.md#PutChargebacksTokenReverseprovisionalcredit) | **Put** /chargebacks/{token}/reverseprovisionalcredit | Reverses provisional credit


# **GetChargebacks**
> ChargebackListResponse GetChargebacks(ctx, optional)
List all chargebacks



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChargebacksApiGetChargebacksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChargebacksApiGetChargebacksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of chargebacks to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **networkReferenceId** | **optional.String**| Network reference ID | 
 **transactionToken** | **optional.String**| Transaction token | 
 **amount** | **optional.String**| Amount | 
 **states** | **optional.String**| Comma-delimited list of chargeback states to display e.g. INITIATED | REPRESENTMENT | PREARBITRATION | ARBITRATION | CASE_WON | CASE_LOST | NETWORK_REJECTED | WITHDRAWN | WRITTEN_OFF_ISSUER | WRITTEN_OFF_PROGRAM | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]
 **networkCaseId** | **optional.String**|  | 

### Return type

[**ChargebackListResponse**](ChargebackListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChargebacksChargebacktokenTransitions**
> ChargebackTransitionListResponse GetChargebacksChargebacktokenTransitions(ctx, chargebackToken, optional)
Lists all chargeback transitions that are related to a chargeback



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chargebackToken** | **string**| Chargeback token | 
 **optional** | ***ChargebacksApiGetChargebacksChargebacktokenTransitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChargebacksApiGetChargebacksChargebacktokenTransitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of transitions to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**ChargebackTransitionListResponse**](ChargebackTransitionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChargebacksToken**
> ChargebackResponse GetChargebacksToken(ctx, token)
Returns a specific chargeback



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**ChargebackResponse**](chargeback_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChargebacksTransitionsToken**
> ChargebackTransitionResponse GetChargebacksTransitionsToken(ctx, token)
Returns a specific chargeback transition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Chargeback transition token | 

### Return type

[**ChargebackTransitionResponse**](chargeback_transition_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostChargebackAllocationAcknowledgment**
> ChargebackResponse PostChargebackAllocationAcknowledgment(ctx, optional)
Creates a chargeback allocation acknowledgement



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChargebacksApiPostChargebackAllocationAcknowledgmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChargebacksApiPostChargebackAllocationAcknowledgmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ChargebackRequest**](ChargebackRequest.md)|  | 

### Return type

[**ChargebackResponse**](chargeback_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostChargebacks**
> ChargebackResponse PostChargebacks(ctx, optional)
Creates a chargeback



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChargebacksApiPostChargebacksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChargebacksApiPostChargebacksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ChargebackRequest**](ChargebackRequest.md)|  | 

### Return type

[**ChargebackResponse**](chargeback_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostChargebacksTransitions**
> ChargebackTransitionResponse PostChargebacksTransitions(ctx, optional)
Creates a chargeback transition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChargebacksApiPostChargebacksTransitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChargebacksApiPostChargebacksTransitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ChargebackTransitionRequest**](ChargebackTransitionRequest.md)|  | 

### Return type

[**ChargebackTransitionResponse**](chargeback_transition_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutChargebacksToken**
> ChargebackResponse PutChargebacksToken(ctx, token, optional)
Updates chargeback data



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 
 **optional** | ***ChargebacksApiPutChargebacksTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChargebacksApiPutChargebacksTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ChargebackUpdateRequest**](ChargebackUpdateRequest.md)|  | 

### Return type

[**ChargebackResponse**](chargeback_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutChargebacksTokenGrantprovisionalcredit**
> ChargebackResponse PutChargebacksTokenGrantprovisionalcredit(ctx, token)
Grants provisional credit



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**ChargebackResponse**](chargeback_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutChargebacksTokenReverseprovisionalcredit**
> ChargebackResponse PutChargebacksTokenReverseprovisionalcredit(ctx, token)
Reverses provisional credit



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**ChargebackResponse**](chargeback_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

