# \OfferOrdersApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOfferordersToken**](OfferOrdersApi.md#GetOfferordersToken) | **Get** /offerorders/{token} | Returns a purchase order for a specific offer
[**PostOfferorders**](OfferOrdersApi.md#PostOfferorders) | **Post** /offerorders | Creates an offer order


# **GetOfferordersToken**
> OfferOrderResponse GetOfferordersToken(ctx, token)
Returns a purchase order for a specific offer



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Order token | 

### Return type

[**OfferOrderResponse**](offer_order_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOfferorders**
> OfferOrderResponse PostOfferorders(ctx, optional)
Creates an offer order



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OfferOrdersApiPostOfferordersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfferOrdersApiPostOfferordersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OfferOrderRequest**](OfferOrderRequest.md)|  | 

### Return type

[**OfferOrderResponse**](offer_order_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

