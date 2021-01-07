# \CardsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCards**](CardsApi.md#GetCards) | **Get** /cards | Lists cards by the last 4 digits
[**GetCardsBarcodeBarcode**](CardsApi.md#GetCardsBarcodeBarcode) | **Get** /cards/barcode/{barcode} | Returns a card&#39;s metadata
[**GetCardsCardHashShowpan**](CardsApi.md#GetCardsCardHashShowpan) | **Get** /cards/{card_hash}/showpanbyhash | Returns a specific card - PAN visible
[**GetCardsMerchantMerchanttoken**](CardsApi.md#GetCardsMerchantMerchanttoken) | **Get** /cards/merchant/{merchant_token} | Returns a merchant onboarding card
[**GetCardsMerchantMerchanttokenShowpan**](CardsApi.md#GetCardsMerchantMerchanttokenShowpan) | **Get** /cards/merchant/{merchant_token}/showpan | Returns a specific card - PAN visible
[**GetCardsToken**](CardsApi.md#GetCardsToken) | **Get** /cards/{token} | Returns a specific card
[**GetCardsTokenShowpan**](CardsApi.md#GetCardsTokenShowpan) | **Get** /cards/{token}/showpan | Returns a specific card - PAN visible
[**GetCardsUserToken**](CardsApi.md#GetCardsUserToken) | **Get** /cards/user/{token} | Lists all cards for a specific user
[**PostCards**](CardsApi.md#PostCards) | **Post** /cards | Creates a card
[**PostCardsGetbypan**](CardsApi.md#PostCardsGetbypan) | **Post** /cards/getbypan | Returns user and card tokens for the specified PAN
[**PostCardsMerchantMerchanttoken**](CardsApi.md#PostCardsMerchantMerchanttoken) | **Post** /cards/merchant/{merchant_token} | Creates a merchant onboarding card
[**PutCardsToken**](CardsApi.md#PutCardsToken) | **Put** /cards/{token} | Updates a specific card


# **GetCards**
> CardListResponse GetCards(ctx, lastFour, optional)
Lists cards by the last 4 digits



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lastFour** | **string**| Last four digits of card number | 
 **optional** | ***CardsApiGetCardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiGetCardsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of cards to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**CardListResponse**](CardListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardsBarcodeBarcode**
> CardResponse GetCardsBarcodeBarcode(ctx, barcode, optional)
Returns a card's metadata



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **barcode** | **string**| Barcode | 
 **optional** | ***CardsApiGetCardsBarcodeBarcodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiGetCardsBarcodeBarcodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**CardResponse**](card_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardsCardHashShowpan**
> CardResponse GetCardsCardHashShowpan(ctx, cardHash, optional)
Returns a specific card - PAN visible



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardHash** | **string**| Card Hash | 
 **optional** | ***CardsApiGetCardsCardHashShowpanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiGetCardsCardHashShowpanOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **showCvvNumber** | **optional.Bool**|  | 

### Return type

[**CardResponse**](card_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardsMerchantMerchanttoken**
> CardResponse GetCardsMerchantMerchanttoken(ctx, merchantToken, optional)
Returns a merchant onboarding card



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **merchantToken** | **string**| Merchant token | 
 **optional** | ***CardsApiGetCardsMerchantMerchanttokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiGetCardsMerchantMerchanttokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**CardResponse**](card_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardsMerchantMerchanttokenShowpan**
> CardResponse GetCardsMerchantMerchanttokenShowpan(ctx, merchantToken, optional)
Returns a specific card - PAN visible



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **merchantToken** | **string**| Merchant token | 
 **optional** | ***CardsApiGetCardsMerchantMerchanttokenShowpanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiGetCardsMerchantMerchanttokenShowpanOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **showCvvNumber** | **optional.Bool**|  | 

### Return type

[**CardResponse**](card_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardsToken**
> CardResponse GetCardsToken(ctx, token, optional)
Returns a specific card



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Card token | 
 **optional** | ***CardsApiGetCardsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiGetCardsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **expand** | **optional.String**| Object to expand | 

### Return type

[**CardResponse**](card_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardsTokenShowpan**
> CardResponse GetCardsTokenShowpan(ctx, token, optional)
Returns a specific card - PAN visible



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Card token | 
 **optional** | ***CardsApiGetCardsTokenShowpanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiGetCardsTokenShowpanOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **showCvvNumber** | **optional.Bool**|  | 

### Return type

[**CardResponse**](card_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardsUserToken**
> CardListResponse GetCardsUserToken(ctx, token, optional)
Lists all cards for a specific user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| User token | 
 **optional** | ***CardsApiGetCardsUserTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiGetCardsUserTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of items to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**CardListResponse**](CardListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCards**
> CardResponse PostCards(ctx, optional)
Creates a card



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CardsApiPostCardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiPostCardsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CardRequest**](CardRequest.md)|  | 
 **showCvvNumber** | **optional.Bool**| Show CVV | [default to false]
 **showPan** | **optional.Bool**| Show PAN | [default to false]

### Return type

[**CardResponse**](card_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCardsGetbypan**
> PanResponse PostCardsGetbypan(ctx, optional)
Returns user and card tokens for the specified PAN



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CardsApiPostCardsGetbypanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiPostCardsGetbypanOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PanRequest**](PanRequest.md)|  | 

### Return type

[**PanResponse**](pan_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCardsMerchantMerchanttoken**
> CardResponse PostCardsMerchantMerchanttoken(ctx, merchantToken, optional)
Creates a merchant onboarding card



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **merchantToken** | **string**| Merchant token | 
 **optional** | ***CardsApiPostCardsMerchantMerchanttokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiPostCardsMerchantMerchanttokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MerchantCardRequest**](MerchantCardRequest.md)|  | 

### Return type

[**CardResponse**](card_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCardsToken**
> CardResponse PutCardsToken(ctx, token, optional)
Updates a specific card



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Card token | 
 **optional** | ***CardsApiPutCardsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiPutCardsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CardUpdateRequest**](CardUpdateRequest.md)|  | 

### Return type

[**CardResponse**](card_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

