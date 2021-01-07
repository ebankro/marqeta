# \AcceptedCountriesApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAcceptedcountries**](AcceptedCountriesApi.md#GetAcceptedcountries) | **Get** /acceptedcountries | Lists all accepted countries
[**GetAcceptedcountriesToken**](AcceptedCountriesApi.md#GetAcceptedcountriesToken) | **Get** /acceptedcountries/{token} | Returns a specific accepted country


# **GetAcceptedcountries**
> AcceptedCountriesListResponse GetAcceptedcountries(ctx, optional)
Lists all accepted countries



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AcceptedCountriesApiGetAcceptedcountriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AcceptedCountriesApiGetAcceptedcountriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of accepted countries to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **name** | **optional.String**| Name | 
 **whitelist** | **optional.Bool**| Whitelist | 
 **searchType** | **optional.String**| Search type | 
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**AcceptedCountriesListResponse**](AcceptedCountriesListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAcceptedcountriesToken**
> AcceptedCountriesModel GetAcceptedcountriesToken(ctx, token, optional)
Returns a specific accepted country



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Accepted country token | 
 **optional** | ***AcceptedCountriesApiGetAcceptedcountriesTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AcceptedCountriesApiGetAcceptedcountriesTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**AcceptedCountriesModel**](accepted_countries_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

