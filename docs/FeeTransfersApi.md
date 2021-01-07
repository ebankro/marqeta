# \FeeTransfersApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeetransfersToken**](FeeTransfersApi.md#GetFeetransfersToken) | **Get** /feetransfers/{token} | Returns a fee transfer
[**PostFeetransfers**](FeeTransfersApi.md#PostFeetransfers) | **Post** /feetransfers | Creates a fee transfer


# **GetFeetransfersToken**
> FeeTransferResponse GetFeetransfersToken(ctx, token)
Returns a fee transfer



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**FeeTransferResponse**](fee_transfer_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFeetransfers**
> FeeTransferResponse PostFeetransfers(ctx, optional)
Creates a fee transfer



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FeeTransfersApiPostFeetransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeeTransfersApiPostFeetransfersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FeeTransferRequest**](FeeTransferRequest.md)|  | 

### Return type

[**FeeTransferResponse**](fee_transfer_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

