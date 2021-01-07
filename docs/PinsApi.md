# \PinsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostPinsControltoken**](PinsApi.md#PostPinsControltoken) | **Post** /pins/controltoken | Creates a new control token for a PIN
[**PutPins**](PinsApi.md#PutPins) | **Put** /pins | Updates the PIN control token
[**RevealPins**](PinsApi.md#RevealPins) | **Post** /pins/reveal | Updates the PIN-reveal control token


# **PostPinsControltoken**
> ControlTokenResponse PostPinsControltoken(ctx, optional)
Creates a new control token for a PIN

Creates a new control token for a PIN, for the specified card for PIN debit or ATM transactions,or to allow for a pin to be revealed to authorized callers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PinsApiPostPinsControltokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PinsApiPostPinsControltokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ControlTokenRequest**](ControlTokenRequest.md)|  | 

### Return type

[**ControlTokenResponse**](control_token_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPins**
> PutPins(ctx, optional)
Updates the PIN control token

Updates a PIN identified by its control token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PinsApiPutPinsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PinsApiPutPinsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PinRequest**](PinRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevealPins**
> RevealPins(ctx, optional)
Updates the PIN-reveal control token

Reveals pin for card associated with given control token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PinsApiRevealPinsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PinsApiRevealPinsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PinRevealRequest**](PinRevealRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

