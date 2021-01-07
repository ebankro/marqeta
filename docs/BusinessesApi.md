# \BusinessesApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBusinesses**](BusinessesApi.md#GetBusinesses) | **Get** /businesses | Lists all businesses
[**GetBusinessesParenttokenChildren**](BusinessesApi.md#GetBusinessesParenttokenChildren) | **Get** /businesses/{parent_token}/children | Lists all children of a parent business
[**GetBusinessesToken**](BusinessesApi.md#GetBusinessesToken) | **Get** /businesses/{token} | Returns a specific business
[**GetBusinessesTokenNotes**](BusinessesApi.md#GetBusinessesTokenNotes) | **Get** /businesses/{token}/notes | Lists business notes
[**GetBusinessesTokenSsn**](BusinessesApi.md#GetBusinessesTokenSsn) | **Get** /businesses/{token}/ssn | Returns a specific business proprietor&#39;s SSN
[**PostBusinesses**](BusinessesApi.md#PostBusinesses) | **Post** /businesses | Creates a business
[**PostBusinessesLookup**](BusinessesApi.md#PostBusinessesLookup) | **Post** /businesses/lookup | Returns a specific business
[**PostBusinessesTokenNotes**](BusinessesApi.md#PostBusinessesTokenNotes) | **Post** /businesses/{token}/notes | Creates a note for a business
[**PutBusinessesToken**](BusinessesApi.md#PutBusinessesToken) | **Put** /businesses/{token} | Updates a specific business
[**PutBusinessesTokenNotesNotestoken**](BusinessesApi.md#PutBusinessesTokenNotesNotestoken) | **Put** /businesses/{token}/notes/{notes_token} | Updates a specific note for a business


# **GetBusinesses**
> BusinessCardHolderListResponse GetBusinesses(ctx, optional)
Lists all businesses



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BusinessesApiGetBusinessesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessesApiGetBusinessesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of users to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **businessNameDba** | **optional.String**| Business name DBA | 
 **businessNameLegal** | **optional.String**| Business name legal | 
 **searchType** | **optional.String**| Search type | 
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**BusinessCardHolderListResponse**](BusinessCardHolderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBusinessesParenttokenChildren**
> UserCardHolderListResponse GetBusinessesParenttokenChildren(ctx, parentToken, optional)
Lists all children of a parent business



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentToken** | **string**| Token of parent business | 
 **optional** | ***BusinessesApiGetBusinessesParenttokenChildrenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessesApiGetBusinessesParenttokenChildrenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of users to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**UserCardHolderListResponse**](UserCardHolderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBusinessesToken**
> BusinessCardHolderResponse GetBusinessesToken(ctx, token, optional)
Returns a specific business



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Business token | 
 **optional** | ***BusinessesApiGetBusinessesTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessesApiGetBusinessesTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**BusinessCardHolderResponse**](business_card_holder_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBusinessesTokenNotes**
> CardHolderNoteListResponse GetBusinessesTokenNotes(ctx, token, optional)
Lists business notes



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Business token | 
 **optional** | ***BusinessesApiGetBusinessesTokenNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessesApiGetBusinessesTokenNotesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **count** | **optional.Int32**| Number of notes to retrieve | [default to 5]
 **createdBy** | **optional.String**| Created by | 
 **createdByUserRole** | **optional.String**| Comma-delimited list of created by user roles | 
 **includePrivate** | **optional.Bool**| Include private notes and private fields in note response | 
 **searchType** | **optional.String**| Search type | 
 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **sortBy** | **optional.String**| Sort order | [default to -lastModifiedTime]

### Return type

[**CardHolderNoteListResponse**](CardHolderNoteListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBusinessesTokenSsn**
> SsnResponseModel GetBusinessesTokenSsn(ctx, token, optional)
Returns a specific business proprietor's SSN



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Business token | 
 **optional** | ***BusinessesApiGetBusinessesTokenSsnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessesApiGetBusinessesTokenSsnOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullSsn** | **optional.Bool**|  | 

### Return type

[**SsnResponseModel**](ssn_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBusinesses**
> BusinessCardHolderResponse PostBusinesses(ctx, optional)
Creates a business



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BusinessesApiPostBusinessesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessesApiPostBusinessesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BusinessCardholder**](BusinessCardholder.md)|  | 

### Return type

[**BusinessCardHolderResponse**](business_card_holder_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBusinessesLookup**
> BusinessCardholder PostBusinessesLookup(ctx, optional)
Returns a specific business



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BusinessesApiPostBusinessesLookupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessesApiPostBusinessesLookupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DdaRequest**](DdaRequest.md)|  | 

### Return type

[**BusinessCardholder**](business_cardholder.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBusinessesTokenNotes**
> CardholderNoteResponseModel PostBusinessesTokenNotes(ctx, token, optional)
Creates a note for a business



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Business token | 
 **optional** | ***BusinessesApiPostBusinessesTokenNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessesApiPostBusinessesTokenNotesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CardholderNoteRequestModel**](CardholderNoteRequestModel.md)|  | 

### Return type

[**CardholderNoteResponseModel**](cardholder_note_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutBusinessesToken**
> BusinessCardholder PutBusinessesToken(ctx, token, body)
Updates a specific business



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Business token | 
  **body** | [**BusinessCardHolderUpdate**](BusinessCardHolderUpdate.md)| Business object | 

### Return type

[**BusinessCardholder**](business_cardholder.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutBusinessesTokenNotesNotestoken**
> CardholderNoteResponseModel PutBusinessesTokenNotesNotestoken(ctx, token, notesToken, optional)
Updates a specific note for a business



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Business token | 
  **notesToken** | **string**| Notes token | 
 **optional** | ***BusinessesApiPutBusinessesTokenNotesNotestokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessesApiPutBusinessesTokenNotesNotestokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of CardholderNoteUpdateRequestModel**](CardholderNoteUpdateRequestModel.md)|  | 

### Return type

[**CardholderNoteResponseModel**](cardholder_note_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

