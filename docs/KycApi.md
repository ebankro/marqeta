# \KycApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKycBusinessBusinesstoken**](KycApi.md#GetKycBusinessBusinesstoken) | **Get** /kyc/business/{business_token} | Lists all KYC results for a business
[**GetKycToken**](KycApi.md#GetKycToken) | **Get** /kyc/{token} | Returns a specific KYC result
[**GetKycUserUsertoken**](KycApi.md#GetKycUserUsertoken) | **Get** /kyc/user/{user_token} | Lists all KYC results for a user
[**PostKyc**](KycApi.md#PostKyc) | **Post** /kyc | Performs a KYC
[**PutKycToken**](KycApi.md#PutKycToken) | **Put** /kyc/{token} | Accepts KYC answers for questions from initial request


# **GetKycBusinessBusinesstoken**
> KycListResponse GetKycBusinessBusinesstoken(ctx, businessToken, optional)
Lists all KYC results for a business



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **businessToken** | **string**| Business token | 
 **optional** | ***KycApiGetKycBusinessBusinesstokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KycApiGetKycBusinessBusinesstokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of items to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

[**KycListResponse**](KYCListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKycToken**
> KycResponse GetKycToken(ctx, token)
Returns a specific KYC result



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| KYC token | 

### Return type

[**KycResponse**](kyc_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKycUserUsertoken**
> KycListResponse GetKycUserUsertoken(ctx, userToken, optional)
Lists all KYC results for a user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userToken** | **string**| User token | 
 **optional** | ***KycApiGetKycUserUsertokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KycApiGetKycUserUsertokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of items to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

[**KycListResponse**](KYCListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostKyc**
> KycResponse PostKyc(ctx, optional)
Performs a KYC



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***KycApiPostKycOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KycApiPostKycOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of KycRequest**](KycRequest.md)|  | 

### Return type

[**KycResponse**](kyc_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutKycToken**
> KycResponse PutKycToken(ctx, token, optional)
Accepts KYC answers for questions from initial request



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| KYC token | 
 **optional** | ***KycApiPutKycTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KycApiPutKycTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of KycSubmitAnswersRequestModel**](KycSubmitAnswersRequestModel.md)|  | 

### Return type

[**KycResponse**](kyc_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

