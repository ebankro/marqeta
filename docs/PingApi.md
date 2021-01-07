# \PingApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPing**](PingApi.md#GetPing) | **Get** /ping | Returns a heartbeat to the consumer
[**PostPing**](PingApi.md#PostPing) | **Post** /ping | Echo test for sending payload to server


# **GetPing**
> PingResponse GetPing(ctx, )
Returns a heartbeat to the consumer

Tests if the Marqeta server is available and responsive.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PingResponse**](ping_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPing**
> EchoPingResponse PostPing(ctx, optional)
Echo test for sending payload to server



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PingApiPostPingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PingApiPostPingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EchoPingRequest**](EchoPingRequest.md)|  | 

### Return type

[**EchoPingResponse**](echo_ping_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

