# \OffersApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOffers**](OffersApi.md#GetOffers) | **Get** /offers | Lists all offers
[**GetOffersToken**](OffersApi.md#GetOffersToken) | **Get** /offers/{token} | Returns a specific offer
[**PostOffers**](OffersApi.md#PostOffers) | **Post** /offers | Creates an offer
[**PutOffersToken**](OffersApi.md#PutOffersToken) | **Put** /offers/{token} | Updates an offer


# **GetOffers**
> OfferListResponse GetOffers(ctx, optional)
Lists all offers



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OffersApiGetOffersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OffersApiGetOffersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of items to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**OfferListResponse**](OfferListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOffersToken**
> OfferResponseModel GetOffersToken(ctx, token, optional)
Returns a specific offer



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Offer token | 
 **optional** | ***OffersApiGetOffersTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OffersApiGetOffersTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**OfferResponseModel**](offer_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOffers**
> OfferResponseModel PostOffers(ctx, optional)
Creates an offer



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OffersApiPostOffersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OffersApiPostOffersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OfferModel**](OfferModel.md)|  | 

### Return type

[**OfferResponseModel**](offer_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutOffersToken**
> OfferResponseModel PutOffersToken(ctx, token, optional)
Updates an offer



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Offer token | 
 **optional** | ***OffersApiPutOffersTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OffersApiPutOffersTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OfferUpdateModel**](OfferUpdateModel.md)|  | 

### Return type

[**OfferResponseModel**](offer_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

