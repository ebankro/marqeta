# \FundingSourcesApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllACHFundingSources**](FundingSourcesApi.md#GetAllACHFundingSources) | **Get** /fundingsources/program/ach | Returns a list of Program ACH funding sources
[**GetFundingsourcesAchFundingsourcetoken**](FundingSourcesApi.md#GetFundingsourcesAchFundingsourcetoken) | **Get** /fundingsources/ach/{funding_source_token} | Returns a user ACH account
[**GetFundingsourcesAchFundingsourcetokenVerificationamounts**](FundingSourcesApi.md#GetFundingsourcesAchFundingsourcetokenVerificationamounts) | **Get** /fundingsources/ach/{funding_source_token}/verificationamounts | Returns the dollar amounts used to verify the ACH account
[**GetFundingsourcesAddressesBusinessBusinesstoken**](FundingSourcesApi.md#GetFundingsourcesAddressesBusinessBusinesstoken) | **Get** /fundingsources/addresses/business/{business_token} | Lists all addresses for a business
[**GetFundingsourcesAddressesFundingsourceaddresstoken**](FundingSourcesApi.md#GetFundingsourcesAddressesFundingsourceaddresstoken) | **Get** /fundingsources/addresses/{funding_source_address_token} | Returns a user address for a funding source
[**GetFundingsourcesAddressesUserUsertoken**](FundingSourcesApi.md#GetFundingsourcesAddressesUserUsertoken) | **Get** /fundingsources/addresses/user/{user_token} | Lists all addresses for a user
[**GetFundingsourcesBusinessBusinesstoken**](FundingSourcesApi.md#GetFundingsourcesBusinessBusinesstoken) | **Get** /fundingsources/business/{business_token} | Lists all funding sources for a business
[**GetFundingsourcesPaymentcardFundingsourcetoken**](FundingSourcesApi.md#GetFundingsourcesPaymentcardFundingsourcetoken) | **Get** /fundingsources/paymentcard/{funding_source_token} | Returns a specific payment card
[**GetFundingsourcesProgramToken**](FundingSourcesApi.md#GetFundingsourcesProgramToken) | **Get** /fundingsources/program/{token} | Returns a specific program funding source
[**GetFundingsourcesProgramgatewayToken**](FundingSourcesApi.md#GetFundingsourcesProgramgatewayToken) | **Get** /fundingsources/programgateway/{token} | Returns a gateway program funding source
[**GetFundingsourcesUserUsertoken**](FundingSourcesApi.md#GetFundingsourcesUserUsertoken) | **Get** /fundingsources/user/{user_token} | Lists all funding sources for a user
[**PostFundingsourcesAch**](FundingSourcesApi.md#PostFundingsourcesAch) | **Post** /fundingsources/ach | Registers an ACH funding source
[**PostFundingsourcesAddresses**](FundingSourcesApi.md#PostFundingsourcesAddresses) | **Post** /fundingsources/addresses | Creates an account holder address for a funding source
[**PostFundingsourcesPaymentcard**](FundingSourcesApi.md#PostFundingsourcesPaymentcard) | **Post** /fundingsources/paymentcard | Registers a payment card funding source
[**PostFundingsourcesProgram**](FundingSourcesApi.md#PostFundingsourcesProgram) | **Post** /fundingsources/program | Creates a program funding source
[**PostFundingsourcesProgramAch**](FundingSourcesApi.md#PostFundingsourcesProgramAch) | **Post** /fundingsources/program/ach | Registers an ACH funding source for a program
[**PostFundingsourcesProgramgateway**](FundingSourcesApi.md#PostFundingsourcesProgramgateway) | **Post** /fundingsources/programgateway | Creates a gateway program funding source
[**PutFundingsourcesAchFundingsourcetoken**](FundingSourcesApi.md#PutFundingsourcesAchFundingsourcetoken) | **Put** /fundingsources/ach/{funding_source_token} | Verifies a bank account as a funding source
[**PutFundingsourcesAddressesFundingsourceaddresstoken**](FundingSourcesApi.md#PutFundingsourcesAddressesFundingsourceaddresstoken) | **Put** /fundingsources/addresses/{funding_source_address_token} | Updates the account holder address for a funding source
[**PutFundingsourcesFundingsourcetoken**](FundingSourcesApi.md#PutFundingsourcesFundingsourcetoken) | **Put** /fundingsources/paymentcard/{funding_source_token} | Updates a specific payment card 
[**PutFundingsourcesFundingsourcetokenDefault**](FundingSourcesApi.md#PutFundingsourcesFundingsourcetokenDefault) | **Put** /fundingsources/{funding_source_token}/default | Configures a default funding source
[**PutFundingsourcesProgramToken**](FundingSourcesApi.md#PutFundingsourcesProgramToken) | **Put** /fundingsources/program/{token} | Updates a specific program funding source
[**PutFundingsourcesProgramgatewayCustomHeaderToken**](FundingSourcesApi.md#PutFundingsourcesProgramgatewayCustomHeaderToken) | **Put** /fundingsources/programgateway/customheaders/{token} | Updates a specific gateway program funding source Custom headers
[**PutFundingsourcesProgramgatewayToken**](FundingSourcesApi.md#PutFundingsourcesProgramgatewayToken) | **Put** /fundingsources/programgateway/{token} | Updates a specific gateway program funding source


# **GetAllACHFundingSources**
> AchListResponse GetAllACHFundingSources(ctx, optional)
Returns a list of Program ACH funding sources



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FundingSourcesApiGetAllACHFundingSourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiGetAllACHFundingSourcesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of items to retrieve. Count can be between 1 - 10 items. | [default to 5]
 **startIndex** | **optional.Int32**| Indicates from what row to start returning data. | [default to 0]
 **fields** | **optional.String**| Comma delimited list of fields to return (e.g. field_1,field_2,..) | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**AchListResponse**](ACHListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFundingsourcesAchFundingsourcetoken**
> AchResponseModel GetFundingsourcesAchFundingsourcetoken(ctx, fundingSourceToken)
Returns a user ACH account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fundingSourceToken** | **string**| Funding account token | 

### Return type

[**AchResponseModel**](ach_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFundingsourcesAchFundingsourcetokenVerificationamounts**
> AchVerificationModel GetFundingsourcesAchFundingsourcetokenVerificationamounts(ctx, fundingSourceToken)
Returns the dollar amounts used to verify the ACH account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fundingSourceToken** | **string**| Funding account token | 

### Return type

[**AchVerificationModel**](ach_verification_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFundingsourcesAddressesBusinessBusinesstoken**
> CardholderAddressListResponse GetFundingsourcesAddressesBusinessBusinesstoken(ctx, businessToken, optional)
Lists all addresses for a business



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **businessToken** | **string**| Business token | 
 **optional** | ***FundingSourcesApiGetFundingsourcesAddressesBusinessBusinesstokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiGetFundingsourcesAddressesBusinessBusinesstokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**CardholderAddressListResponse**](CardholderAddressListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFundingsourcesAddressesFundingsourceaddresstoken**
> CardholderAddressResponse GetFundingsourcesAddressesFundingsourceaddresstoken(ctx, fundingSourceAddressToken)
Returns a user address for a funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fundingSourceAddressToken** | **string**| Funding source address token | 

### Return type

[**CardholderAddressResponse**](CardholderAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFundingsourcesAddressesUserUsertoken**
> CardholderAddressListResponse GetFundingsourcesAddressesUserUsertoken(ctx, userToken, optional)
Lists all addresses for a user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userToken** | **string**| User token | 
 **optional** | ***FundingSourcesApiGetFundingsourcesAddressesUserUsertokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiGetFundingsourcesAddressesUserUsertokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**CardholderAddressListResponse**](CardholderAddressListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFundingsourcesBusinessBusinesstoken**
> FundingAccountListResponse GetFundingsourcesBusinessBusinesstoken(ctx, businessToken, optional)
Lists all funding sources for a business



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **businessToken** | **string**| Business token | 
 **optional** | ***FundingSourcesApiGetFundingsourcesBusinessBusinesstokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiGetFundingsourcesBusinessBusinesstokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Type, such as a payment card or ACH | 
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**FundingAccountListResponse**](FundingAccountListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFundingsourcesPaymentcardFundingsourcetoken**
> PaymentCardResponseModel GetFundingsourcesPaymentcardFundingsourcetoken(ctx, fundingSourceToken)
Returns a specific payment card



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fundingSourceToken** | **string**| Funding token | 

### Return type

[**PaymentCardResponseModel**](payment_card_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFundingsourcesProgramToken**
> ProgramFundingSourceResponse GetFundingsourcesProgramToken(ctx, token)
Returns a specific program funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Program funding source token | 

### Return type

[**ProgramFundingSourceResponse**](program_funding_source_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFundingsourcesProgramgatewayToken**
> GatewayProgramFundingSourceResponse GetFundingsourcesProgramgatewayToken(ctx, token)
Returns a gateway program funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Gateway program funding source token | 

### Return type

[**GatewayProgramFundingSourceResponse**](gateway_program_funding_source_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFundingsourcesUserUsertoken**
> FundingAccountListResponse GetFundingsourcesUserUsertoken(ctx, userToken, optional)
Lists all funding sources for a user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userToken** | **string**| User token | 
 **optional** | ***FundingSourcesApiGetFundingsourcesUserUsertokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiGetFundingsourcesUserUsertokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Type, such as a payment card or ACH | 
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**FundingAccountListResponse**](FundingAccountListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFundingsourcesAch**
> AchResponseModel PostFundingsourcesAch(ctx, optional)
Registers an ACH funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FundingSourcesApiPostFundingsourcesAchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiPostFundingsourcesAchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AchModel**](AchModel.md)|  | 

### Return type

[**AchResponseModel**](ach_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFundingsourcesAddresses**
> CardholderAddressResponse PostFundingsourcesAddresses(ctx, optional)
Creates an account holder address for a funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FundingSourcesApiPostFundingsourcesAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiPostFundingsourcesAddressesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CardHolderAddressModel**](CardHolderAddressModel.md)|  | 

### Return type

[**CardholderAddressResponse**](CardholderAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFundingsourcesPaymentcard**
> PaymentCardResponseModel PostFundingsourcesPaymentcard(ctx, optional)
Registers a payment card funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FundingSourcesApiPostFundingsourcesPaymentcardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiPostFundingsourcesPaymentcardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TokenRequest**](TokenRequest.md)|  | 

### Return type

[**PaymentCardResponseModel**](payment_card_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFundingsourcesProgram**
> ProgramFundingSourceResponse PostFundingsourcesProgram(ctx, optional)
Creates a program funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FundingSourcesApiPostFundingsourcesProgramOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiPostFundingsourcesProgramOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ProgramFundingSourceRequest**](ProgramFundingSourceRequest.md)|  | 

### Return type

[**ProgramFundingSourceResponse**](program_funding_source_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFundingsourcesProgramAch**
> BaseAchResponseModel PostFundingsourcesProgramAch(ctx, optional)
Registers an ACH funding source for a program



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FundingSourcesApiPostFundingsourcesProgramAchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiPostFundingsourcesProgramAchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BaseAchRequestModel**](BaseAchRequestModel.md)|  | 

### Return type

[**BaseAchResponseModel**](base_ach_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFundingsourcesProgramgateway**
> GatewayProgramFundingSourceResponse PostFundingsourcesProgramgateway(ctx, optional)
Creates a gateway program funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FundingSourcesApiPostFundingsourcesProgramgatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiPostFundingsourcesProgramgatewayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GatewayProgramFundingSourceRequest**](GatewayProgramFundingSourceRequest.md)|  | 

### Return type

[**GatewayProgramFundingSourceResponse**](gateway_program_funding_source_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFundingsourcesAchFundingsourcetoken**
> AchResponseModel PutFundingsourcesAchFundingsourcetoken(ctx, fundingSourceToken, optional)
Verifies a bank account as a funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fundingSourceToken** | **string**|  | 
 **optional** | ***FundingSourcesApiPutFundingsourcesAchFundingsourcetokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiPutFundingsourcesAchFundingsourcetokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AchVerificationModel**](AchVerificationModel.md)|  | 

### Return type

[**AchResponseModel**](ach_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFundingsourcesAddressesFundingsourceaddresstoken**
> CardholderAddressResponse PutFundingsourcesAddressesFundingsourceaddresstoken(ctx, fundingSourceAddressToken, optional)
Updates the account holder address for a funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fundingSourceAddressToken** | **string**| Funding source address token | 
 **optional** | ***FundingSourcesApiPutFundingsourcesAddressesFundingsourceaddresstokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiPutFundingsourcesAddressesFundingsourceaddresstokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CardHolderAddressUpdateModel**](CardHolderAddressUpdateModel.md)|  | 

### Return type

[**CardholderAddressResponse**](CardholderAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFundingsourcesFundingsourcetoken**
> PaymentCardResponseModel PutFundingsourcesFundingsourcetoken(ctx, fundingSourceToken, body)
Updates a specific payment card 



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fundingSourceToken** | **string**| Funding account token | 
  **body** | [**TokenUpdateRequest**](TokenUpdateRequest.md)| Payment card | 

### Return type

[**PaymentCardResponseModel**](payment_card_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFundingsourcesFundingsourcetokenDefault**
> PaymentCardResponseModel PutFundingsourcesFundingsourcetokenDefault(ctx, fundingSourceToken)
Configures a default funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fundingSourceToken** | **string**| Funding account | 

### Return type

[**PaymentCardResponseModel**](payment_card_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFundingsourcesProgramToken**
> ProgramFundingSourceResponse PutFundingsourcesProgramToken(ctx, token, optional)
Updates a specific program funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Program funding source token | 
 **optional** | ***FundingSourcesApiPutFundingsourcesProgramTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiPutFundingsourcesProgramTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProgramFundingSourceUpdateRequest**](ProgramFundingSourceUpdateRequest.md)|  | 

### Return type

[**ProgramFundingSourceResponse**](program_funding_source_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFundingsourcesProgramgatewayCustomHeaderToken**
> GatewayProgramFundingSourceResponse PutFundingsourcesProgramgatewayCustomHeaderToken(ctx, token, optional)
Updates a specific gateway program funding source Custom headers



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Gateway program funding source token | 
 **optional** | ***FundingSourcesApiPutFundingsourcesProgramgatewayCustomHeaderTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiPutFundingsourcesProgramgatewayCustomHeaderTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GatewayProgramCustomHeaderUpdateRequest**](GatewayProgramCustomHeaderUpdateRequest.md)|  | 

### Return type

[**GatewayProgramFundingSourceResponse**](gateway_program_funding_source_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFundingsourcesProgramgatewayToken**
> GatewayProgramFundingSourceResponse PutFundingsourcesProgramgatewayToken(ctx, token, optional)
Updates a specific gateway program funding source



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Gateway program funding source token | 
 **optional** | ***FundingSourcesApiPutFundingsourcesProgramgatewayTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundingSourcesApiPutFundingsourcesProgramgatewayTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GatewayProgramFundingSourceUpdateRequest**](GatewayProgramFundingSourceUpdateRequest.md)|  | 

### Return type

[**GatewayProgramFundingSourceResponse**](gateway_program_funding_source_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

