# \PeerTransfersApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPeertransfersToken**](PeerTransfersApi.md#GetPeertransfersToken) | **Get** /peertransfers/{token} | Returns details of a previous transfer
[**GetPeertransfersUserUserorbusinesstoken**](PeerTransfersApi.md#GetPeertransfersUserUserorbusinesstoken) | **Get** /peertransfers/user/{user_or_business_token} | Returns all peer transfers for a user
[**GetPeertransfersUserUserorbusinesstokenRecipient**](PeerTransfersApi.md#GetPeertransfersUserUserorbusinesstokenRecipient) | **Get** /peertransfers/user/{user_or_business_token}/recipient | Returns received peer transfers for a user
[**GetPeertransfersUserUserorbusinesstokenSender**](PeerTransfersApi.md#GetPeertransfersUserUserorbusinesstokenSender) | **Get** /peertransfers/user/{user_or_business_token}/sender | Returns sent peer transfers for a user
[**PostPeertransfers**](PeerTransfersApi.md#PostPeertransfers) | **Post** /peertransfers | Performs a peer transfer from one user to another


# **GetPeertransfersToken**
> PeerTransferResponse GetPeertransfersToken(ctx, token)
Returns details of a previous transfer



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**PeerTransferResponse**](peer_transfer_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPeertransfersUserUserorbusinesstoken**
> PeerTransferResponse GetPeertransfersUserUserorbusinesstoken(ctx, userOrBusinessToken, optional)
Returns all peer transfers for a user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userOrBusinessToken** | **string**| User or business token | 
 **optional** | ***PeerTransfersApiGetPeertransfersUserUserorbusinesstokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PeerTransfersApiGetPeertransfersUserUserorbusinesstokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of transfers to retrieve | [default to 25]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**PeerTransferResponse**](peer_transfer_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPeertransfersUserUserorbusinesstokenRecipient**
> PeerTransferResponse GetPeertransfersUserUserorbusinesstokenRecipient(ctx, userOrBusinessToken, optional)
Returns received peer transfers for a user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userOrBusinessToken** | **string**| User or business token | 
 **optional** | ***PeerTransfersApiGetPeertransfersUserUserorbusinesstokenRecipientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PeerTransfersApiGetPeertransfersUserUserorbusinesstokenRecipientOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of transfers to retrieve | [default to 25]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**PeerTransferResponse**](peer_transfer_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPeertransfersUserUserorbusinesstokenSender**
> PeerTransferResponse GetPeertransfersUserUserorbusinesstokenSender(ctx, userOrBusinessToken, optional)
Returns sent peer transfers for a user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userOrBusinessToken** | **string**| User or business token | 
 **optional** | ***PeerTransfersApiGetPeertransfersUserUserorbusinesstokenSenderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PeerTransfersApiGetPeertransfersUserUserorbusinesstokenSenderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of transfers to retrieve | [default to 25]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**PeerTransferResponse**](peer_transfer_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPeertransfers**
> PeerTransferResponse PostPeertransfers(ctx, optional)
Performs a peer transfer from one user to another



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PeerTransfersApiPostPeertransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PeerTransfersApiPostPeertransfersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PeerTransferRequest**](PeerTransferRequest.md)|  | 

### Return type

[**PeerTransferResponse**](peer_transfer_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

