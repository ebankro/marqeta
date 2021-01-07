# \ProgramTransfersApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProgramtransfers**](ProgramTransfersApi.md#GetProgramtransfers) | **Get** /programtransfers | Lists all program transfers
[**GetProgramtransfersToken**](ProgramTransfersApi.md#GetProgramtransfersToken) | **Get** /programtransfers/{token} | Returns a specific program transfer
[**GetProgramtransfersTypes**](ProgramTransfersApi.md#GetProgramtransfersTypes) | **Get** /programtransfers/types | Lists all program transfer types
[**GetProgramtransfersTypesTypetoken**](ProgramTransfersApi.md#GetProgramtransfersTypesTypetoken) | **Get** /programtransfers/types/{type_token} | Returns a specific program transfer type
[**PostProgramtransfers**](ProgramTransfersApi.md#PostProgramtransfers) | **Post** /programtransfers | Transfers to a program funding source
[**PostProgramtransfersTypes**](ProgramTransfersApi.md#PostProgramtransfersTypes) | **Post** /programtransfers/types | Creates a program transfer type
[**PutProgramtransfersTypesTypetoken**](ProgramTransfersApi.md#PutProgramtransfersTypesTypetoken) | **Put** /programtransfers/types/{type_token} | Updates a specific program transfer type


# **GetProgramtransfers**
> ProgramTransferListResponse GetProgramtransfers(ctx, optional)
Lists all program transfers



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProgramTransfersApiGetProgramtransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgramTransfersApiGetProgramtransfersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of program transfers to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]
 **userToken** | **optional.String**| User token | 
 **businessToken** | **optional.String**| Business token | 
 **typeToken** | **optional.String**| Program type token | 

### Return type

[**ProgramTransferListResponse**](ProgramTransferListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramtransfersToken**
> ProgramTransferResponse GetProgramtransfersToken(ctx, token)
Returns a specific program transfer



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Program transfer token | 

### Return type

[**ProgramTransferResponse**](program_transfer_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramtransfersTypes**
> ProgramTransferTypeListResponse GetProgramtransfersTypes(ctx, optional)
Lists all program transfer types



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProgramTransfersApiGetProgramtransfersTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgramTransfersApiGetProgramtransfersTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of program transfer types to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**ProgramTransferTypeListResponse**](ProgramTransferTypeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramtransfersTypesTypetoken**
> ProgramTransferTypeReponse GetProgramtransfersTypesTypetoken(ctx, typeToken)
Returns a specific program transfer type



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **typeToken** | **string**| Type token | 

### Return type

[**ProgramTransferTypeReponse**](program_transfer_type_reponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostProgramtransfers**
> ProgramTransferResponse PostProgramtransfers(ctx, optional)
Transfers to a program funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProgramTransfersApiPostProgramtransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgramTransfersApiPostProgramtransfersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ProgramTransfer**](ProgramTransfer.md)|  | 

### Return type

[**ProgramTransferResponse**](program_transfer_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostProgramtransfersTypes**
> ProgramTransferTypeReponse PostProgramtransfersTypes(ctx, optional)
Creates a program transfer type



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProgramTransfersApiPostProgramtransfersTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgramTransfersApiPostProgramtransfersTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ProgramTransferTypeRequest**](ProgramTransferTypeRequest.md)|  | 

### Return type

[**ProgramTransferTypeReponse**](program_transfer_type_reponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutProgramtransfersTypesTypetoken**
> ProgramTransferTypeReponse PutProgramtransfersTypesTypetoken(ctx, typeToken, optional)
Updates a specific program transfer type



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **typeToken** | **string**| Type token | 
 **optional** | ***ProgramTransfersApiPutProgramtransfersTypesTypetokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgramTransfersApiPutProgramtransfersTypesTypetokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProgramTransferTypeRequest**](ProgramTransferTypeRequest.md)|  | 

### Return type

[**ProgramTransferTypeReponse**](program_transfer_type_reponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

