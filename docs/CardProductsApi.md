# \CardProductsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCardproducts**](CardProductsApi.md#GetCardproducts) | **Get** /cardproducts | Lists all card products
[**GetCardproductsToken**](CardProductsApi.md#GetCardproductsToken) | **Get** /cardproducts/{token} | Returns a specific card product
[**PostCardproducts**](CardProductsApi.md#PostCardproducts) | **Post** /cardproducts | Creates a card product
[**PutCardproductsToken**](CardProductsApi.md#PutCardproductsToken) | **Put** /cardproducts/{token} | Updates a specific card product


# **GetCardproducts**
> CardProductListResponse GetCardproducts(ctx, optional)
Lists all card products



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CardProductsApiGetCardproductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardProductsApiGetCardproductsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of items to retrieve. Count can be between 1 - 10 items. | [default to 5]
 **startIndex** | **optional.Int32**| Indicates from what row to start returning data. | [default to 0]
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

[**CardProductListResponse**](CardProductListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardproductsToken**
> CardProductResponse GetCardproductsToken(ctx, token)
Returns a specific card product



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Card product token | 

### Return type

[**CardProductResponse**](card_product_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCardproducts**
> CardProductResponse PostCardproducts(ctx, body)
Creates a card product



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CardProductRequest**](CardProductRequest.md)| Card product object | 

### Return type

[**CardProductResponse**](card_product_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCardproductsToken**
> CardProductResponse PutCardproductsToken(ctx, token, body)
Updates a specific card product



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Card product token | 
  **body** | [**CardProductUpdateModel**](CardProductUpdateModel.md)| Card product object | 

### Return type

[**CardProductResponse**](card_product_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

