# \BulkIssuancesApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBulkissuances**](BulkIssuancesApi.md#GetBulkissuances) | **Get** /bulkissuances | Lists all bulk issuance requests
[**GetBulkissuancesToken**](BulkIssuancesApi.md#GetBulkissuancesToken) | **Get** /bulkissuances/{token} | Returns a bulk issuance request
[**PostBulkissuances**](BulkIssuancesApi.md#PostBulkissuances) | **Post** /bulkissuances | Creates a bulk issuance request for cards


# **GetBulkissuances**
> BulkCardOrderListResponse GetBulkissuances(ctx, optional)
Lists all bulk issuance requests



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BulkIssuancesApiGetBulkissuancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BulkIssuancesApiGetBulkissuancesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of requests to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

[**BulkCardOrderListResponse**](BulkCardOrderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBulkissuancesToken**
> BulkIssuanceResponse GetBulkissuancesToken(ctx, token)
Returns a bulk issuance request



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Bulk issuance token | 

### Return type

[**BulkIssuanceResponse**](bulk_issuance_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBulkissuances**
> BulkIssuanceResponse PostBulkissuances(ctx, optional)
Creates a bulk issuance request for cards



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BulkIssuancesApiPostBulkissuancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BulkIssuancesApiPostBulkissuancesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BulkIssuanceRequest**](BulkIssuanceRequest.md)|  | 

### Return type

[**BulkIssuanceResponse**](bulk_issuance_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

