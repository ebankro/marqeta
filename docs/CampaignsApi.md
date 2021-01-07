# \CampaignsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCampaigns**](CampaignsApi.md#GetCampaigns) | **Get** /campaigns | Lists all campaigns
[**GetCampaignsToken**](CampaignsApi.md#GetCampaignsToken) | **Get** /campaigns/{token} | Returns a specific campaign
[**GetCampaignsTokenStores**](CampaignsApi.md#GetCampaignsTokenStores) | **Get** /campaigns/{token}/stores | Lists all stores associated with a campaign
[**PostCampaigns**](CampaignsApi.md#PostCampaigns) | **Post** /campaigns | Creates a campaign
[**PutCampaignsToken**](CampaignsApi.md#PutCampaignsToken) | **Put** /campaigns/{token} | Updates a specific campaign


# **GetCampaigns**
> CampaignListResponse GetCampaigns(ctx, optional)
Lists all campaigns



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CampaignsApiGetCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiGetCampaignsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of items to retrieve. Recommended max is 10. | [default to 5]
 **startIndex** | **optional.Int32**| Indicates from what row to start returning data. Used in conjunction with count. | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Field on which to sort the returned items. Use any field in the model, or system fields lastModifiedTime or CreatedTime. | [default to -lastModifiedTime]

### Return type

[**CampaignListResponse**](CampaignListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignsToken**
> CampaignResponseModel GetCampaignsToken(ctx, token, optional)
Returns a specific campaign



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Campaign token | 
 **optional** | ***CampaignsApiGetCampaignsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiGetCampaignsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**CampaignResponseModel**](campaign_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignsTokenStores**
> StoreListResponse GetCampaignsTokenStores(ctx, token, optional)
Lists all stores associated with a campaign



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Campaign token | 
 **optional** | ***CampaignsApiGetCampaignsTokenStoresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiGetCampaignsTokenStoresOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of items to retrieve. Recommended max is 10. | [default to 5]
 **startIndex** | **optional.Int32**| Indicates from what row to start returning data. Used in conjunction with count. | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Field on which to sort the returned items. Use any field in the model, or system fields lastModifiedTime or CreatedTime. | [default to -lastModifiedTime]

### Return type

[**StoreListResponse**](StoreListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaigns**
> CampaignResponseModel PostCampaigns(ctx, optional)
Creates a campaign



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CampaignsApiPostCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiPostCampaignsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CampaignModel**](CampaignModel.md)|  | 

### Return type

[**CampaignResponseModel**](campaign_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCampaignsToken**
> CampaignResponseModel PutCampaignsToken(ctx, token, optional)
Updates a specific campaign



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Campaign token | 
 **optional** | ***CampaignsApiPutCampaignsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiPutCampaignsTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CampaignUpdateModel**](CampaignUpdateModel.md)|  | 

### Return type

[**CampaignResponseModel**](campaign_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

