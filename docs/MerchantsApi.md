# \MerchantsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchants**](MerchantsApi.md#GetMerchants) | **Get** /merchants | Lists all merchants
[**GetMerchantsToken**](MerchantsApi.md#GetMerchantsToken) | **Get** /merchants/{token} | Returns a specific merchant
[**GetMerchantsTokenStores**](MerchantsApi.md#GetMerchantsTokenStores) | **Get** /merchants/{token}/stores | Lists the stores associated with a specific merchant
[**PostMerchants**](MerchantsApi.md#PostMerchants) | **Post** /merchants | Creates a merchant
[**PutMerchantsToken**](MerchantsApi.md#PutMerchantsToken) | **Put** /merchants/{token} | Updates a specific merchant


# **GetMerchants**
> MerchantListResponse GetMerchants(ctx, optional)
Lists all merchants



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MerchantsApiGetMerchantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantsApiGetMerchantsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of items to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**MerchantListResponse**](MerchantListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMerchantsToken**
> MerchantResponseModel GetMerchantsToken(ctx, token, optional)
Returns a specific merchant



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Merchant token | 
 **optional** | ***MerchantsApiGetMerchantsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantsApiGetMerchantsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**MerchantResponseModel**](merchant_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMerchantsTokenStores**
> StoreListResponse GetMerchantsTokenStores(ctx, token, optional)
Lists the stores associated with a specific merchant



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Merchant token | 
 **optional** | ***MerchantsApiGetMerchantsTokenStoresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantsApiGetMerchantsTokenStoresOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of items to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**StoreListResponse**](StoreListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMerchants**
> MerchantResponseModel PostMerchants(ctx, optional)
Creates a merchant



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MerchantsApiPostMerchantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantsApiPostMerchantsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MerchantModel**](MerchantModel.md)|  | 

### Return type

[**MerchantResponseModel**](merchant_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMerchantsToken**
> MerchantResponseModel PutMerchantsToken(ctx, token, optional)
Updates a specific merchant



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Merchant token | 
 **optional** | ***MerchantsApiPutMerchantsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantsApiPutMerchantsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MerchantUpdateModel**](MerchantUpdateModel.md)|  | 

### Return type

[**MerchantResponseModel**](merchant_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

