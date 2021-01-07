# \DigitalWalletProvisionRequestsApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostDigitalwalletprovisionrequestsAndroidpay**](DigitalWalletProvisionRequestsApi.md#PostDigitalwalletprovisionrequestsAndroidpay) | **Post** /digitalwalletprovisionrequests/androidpay | Validates and creates Android-specific provisioning request data
[**PostDigitalwalletprovisionrequestsApplepay**](DigitalWalletProvisionRequestsApi.md#PostDigitalwalletprovisionrequestsApplepay) | **Post** /digitalwalletprovisionrequests/applepay | Validates Apple certificates and creates Apple-specific provisioning request data
[**PostDigitalwalletprovisionrequestsSamsungpay**](DigitalWalletProvisionRequestsApi.md#PostDigitalwalletprovisionrequestsSamsungpay) | **Post** /digitalwalletprovisionrequests/samsungpay | Validates and creates Samsung-specific provisioning request data


# **PostDigitalwalletprovisionrequestsAndroidpay**
> DigitalWalletAndroidPayProvisionResponse PostDigitalwalletprovisionrequestsAndroidpay(ctx, optional)
Validates and creates Android-specific provisioning request data



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DigitalWalletProvisionRequestsApiPostDigitalwalletprovisionrequestsAndroidpayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DigitalWalletProvisionRequestsApiPostDigitalwalletprovisionrequestsAndroidpayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DigitalWalletAndroidPayProvisionRequest**](DigitalWalletAndroidPayProvisionRequest.md)|  | 

### Return type

[**DigitalWalletAndroidPayProvisionResponse**](digital_wallet_android_pay_provision_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDigitalwalletprovisionrequestsApplepay**
> DigitalWalletApplePayProvisionResponse PostDigitalwalletprovisionrequestsApplepay(ctx, optional)
Validates Apple certificates and creates Apple-specific provisioning request data



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DigitalWalletProvisionRequestsApiPostDigitalwalletprovisionrequestsApplepayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DigitalWalletProvisionRequestsApiPostDigitalwalletprovisionrequestsApplepayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DigitalWalletApplePayProvisionRequest**](DigitalWalletApplePayProvisionRequest.md)|  | 

### Return type

[**DigitalWalletApplePayProvisionResponse**](digital_wallet_apple_pay_provision_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDigitalwalletprovisionrequestsSamsungpay**
> DigitalWalletSamsungPayProvisionResponse PostDigitalwalletprovisionrequestsSamsungpay(ctx, optional)
Validates and creates Samsung-specific provisioning request data



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DigitalWalletProvisionRequestsApiPostDigitalwalletprovisionrequestsSamsungpayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DigitalWalletProvisionRequestsApiPostDigitalwalletprovisionrequestsSamsungpayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DigitalWalletAndroidPayProvisionRequest**](DigitalWalletAndroidPayProvisionRequest.md)|  | 

### Return type

[**DigitalWalletSamsungPayProvisionResponse**](digital_wallet_samsung_pay_provision_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

