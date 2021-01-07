# \WebhooksApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebhooks**](WebhooksApi.md#GetWebhooks) | **Get** /webhooks | Returns a list of webhook configurations
[**GetWebhooksToken**](WebhooksApi.md#GetWebhooksToken) | **Get** /webhooks/{token} | Returns a webhook configuration
[**PostWebhooks**](WebhooksApi.md#PostWebhooks) | **Post** /webhooks | Creates a webhook configuration
[**PostWebhooksTokenEventtypeEventtoken**](WebhooksApi.md#PostWebhooksTokenEventtypeEventtoken) | **Post** /webhooks/{token}/{event_type}/{event_token} | Replays an event to a webhook
[**PostWebhooksTokenPing**](WebhooksApi.md#PostWebhooksTokenPing) | **Post** /webhooks/{token}/ping | Pings a webhook
[**PutWebhooksCustomHeadersToken**](WebhooksApi.md#PutWebhooksCustomHeadersToken) | **Put** /webhooks/customheaders/{token} | Updates a specific webhook configuration with custom headers
[**PutWebhooksToken**](WebhooksApi.md#PutWebhooksToken) | **Put** /webhooks/{token} | Updates a specific webhook configuration


# **GetWebhooks**
> WebhookResponseModelListResponse GetWebhooks(ctx, optional)
Returns a list of webhook configurations



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksApiGetWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiGetWebhooksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **optional.Bool**| Return only active webhook configurations | [default to false]
 **count** | **optional.Int32**| Number of reward programs to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -createdTime]

### Return type

[**WebhookResponseModelListResponse**](WebhookResponseModelListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhooksToken**
> WebhookResponseModel GetWebhooksToken(ctx, token)
Returns a webhook configuration



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Webhook token | 

### Return type

[**WebhookResponseModel**](webhook_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWebhooks**
> WebhookResponseModel PostWebhooks(ctx, optional)
Creates a webhook configuration



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksApiPostWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiPostWebhooksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookRequestModel**](WebhookRequestModel.md)|  | 

### Return type

[**WebhookResponseModel**](webhook_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWebhooksTokenEventtypeEventtoken**
> PostWebhooksTokenEventtypeEventtoken(ctx, token, eventType, eventToken)
Replays an event to a webhook



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Webhook token | 
  **eventType** | **string**| Event type | 
  **eventToken** | **string**| Event token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWebhooksTokenPing**
> WebhookPingModel PostWebhooksTokenPing(ctx, token)
Pings a webhook

Endpoints must respond with a 200 status code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Ping a webhook | 

### Return type

[**WebhookPingModel**](webhook_ping_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutWebhooksCustomHeadersToken**
> WebhookResponseModel PutWebhooksCustomHeadersToken(ctx, token, optional)
Updates a specific webhook configuration with custom headers



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Webhook token | 
 **optional** | ***WebhooksApiPutWebhooksCustomHeadersTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiPutWebhooksCustomHeadersTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of WebhookUpdateCustomHeaderRequest**](WebhookUpdateCustomHeaderRequest.md)|  | 

### Return type

[**WebhookResponseModel**](webhook_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutWebhooksToken**
> WebhookResponseModel PutWebhooksToken(ctx, token, optional)
Updates a specific webhook configuration



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Webhook token | 
 **optional** | ***WebhooksApiPutWebhooksTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiPutWebhooksTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of WebhookBaseModel**](WebhookBaseModel.md)|  | 

### Return type

[**WebhookResponseModel**](webhook_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

