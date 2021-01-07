# \UsersApi

All URIs are relative to *https://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /users | Lists all users
[**GetUsersAuthClientaccesstokenToken**](UsersApi.md#GetUsersAuthClientaccesstokenToken) | **Get** /users/auth/clientaccesstoken/{token} | Returns a client access token
[**GetUsersParenttokenChildren**](UsersApi.md#GetUsersParenttokenChildren) | **Get** /users/{parent_token}/children | Lists all children of a parent user
[**GetUsersPhonenumberPhonenumber**](UsersApi.md#GetUsersPhonenumberPhonenumber) | **Get** /users/phonenumber/{phone_number} | Lists all users who match a phone number
[**GetUsersToken**](UsersApi.md#GetUsersToken) | **Get** /users/{token} | Returns a specific user
[**GetUsersTokenNotes**](UsersApi.md#GetUsersTokenNotes) | **Get** /users/{token}/notes | Lists cardholder notes
[**GetUsersTokenSsn**](UsersApi.md#GetUsersTokenSsn) | **Get** /users/{token}/ssn | Returns a specific user&#39;s SSN
[**PostUsers**](UsersApi.md#PostUsers) | **Post** /users | Creates a user
[**PostUsersAuthChangepassword**](UsersApi.md#PostUsersAuthChangepassword) | **Post** /users/auth/changepassword | Updates a user password
[**PostUsersAuthClientaccesstoken**](UsersApi.md#PostUsersAuthClientaccesstoken) | **Post** /users/auth/clientaccesstoken | Creates a client access token
[**PostUsersAuthLogin**](UsersApi.md#PostUsersAuthLogin) | **Post** /users/auth/login | Logs in a user
[**PostUsersAuthLogout**](UsersApi.md#PostUsersAuthLogout) | **Post** /users/auth/logout | Logs out a user
[**PostUsersAuthOnetime**](UsersApi.md#PostUsersAuthOnetime) | **Post** /users/auth/onetime | Creates a one-time token
[**PostUsersAuthResetpassword**](UsersApi.md#PostUsersAuthResetpassword) | **Post** /users/auth/resetpassword | Generates a reset password email
[**PostUsersAuthResetpasswordToken**](UsersApi.md#PostUsersAuthResetpasswordToken) | **Post** /users/auth/resetpassword/{token} | Resets a user password
[**PostUsersAuthVerifyemail**](UsersApi.md#PostUsersAuthVerifyemail) | **Post** /users/auth/verifyemail | Generates an email verification request
[**PostUsersAuthVerifyemailToken**](UsersApi.md#PostUsersAuthVerifyemailToken) | **Post** /users/auth/verifyemail/{token} | Verifies the email token
[**PostUsersLookup**](UsersApi.md#PostUsersLookup) | **Post** /users/lookup | Lists all users
[**PostUsersTokenNotes**](UsersApi.md#PostUsersTokenNotes) | **Post** /users/{token}/notes | Creates a note for the cardholder
[**PutUsersToken**](UsersApi.md#PutUsersToken) | **Put** /users/{token} | Updates a specific user
[**PutUsersTokenNotesNotestoken**](UsersApi.md#PutUsersTokenNotesNotestoken) | **Put** /users/{token}/notes/{notes_token} | Updates a specific note for a cardholder


# **GetUsers**
> UserCardHolderListResponse GetUsers(ctx, optional)
Lists all users



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiGetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of users to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **searchType** | **optional.String**| Search type | 
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

# **GetUsersAuthClientaccesstokenToken**
> ClientAccessTokenResponse GetUsersAuthClientaccesstokenToken(ctx, token, optional)
Returns a client access token



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Client access token | 
 **optional** | ***UsersApiGetUsersAuthClientaccesstokenTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetUsersAuthClientaccesstokenTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationToken** | **optional.String**|  | 

### Return type

[**ClientAccessTokenResponse**](client_access_token_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersParenttokenChildren**
> UserCardHolderListResponse GetUsersParenttokenChildren(ctx, parentToken, optional)
Lists all children of a parent user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentToken** | **string**| Token of parent cardholder | 
 **optional** | ***UsersApiGetUsersParenttokenChildrenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetUsersParenttokenChildrenOpts struct

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

# **GetUsersPhonenumberPhonenumber**
> UserCardHolderListResponse GetUsersPhonenumberPhonenumber(ctx, phoneNumber, optional)
Lists all users who match a phone number



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phoneNumber** | **string**| Phone number | 
 **optional** | ***UsersApiGetUsersPhonenumberPhonenumberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetUsersPhonenumberPhonenumberOpts struct

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

# **GetUsersToken**
> UserCardHolderResponse GetUsersToken(ctx, token, optional)
Returns a specific user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| User token | 
 **optional** | ***UsersApiGetUsersTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetUsersTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**UserCardHolderResponse**](user_card_holder_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersTokenNotes**
> CardHolderNoteListResponse GetUsersTokenNotes(ctx, token, optional)
Lists cardholder notes



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| User token | 
 **optional** | ***UsersApiGetUsersTokenNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetUsersTokenNotesOpts struct

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

# **GetUsersTokenSsn**
> SsnResponseModel GetUsersTokenSsn(ctx, token, optional)
Returns a specific user's SSN



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| User token | 
 **optional** | ***UsersApiGetUsersTokenSsnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetUsersTokenSsnOpts struct

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

# **PostUsers**
> UserCardHolderResponse PostUsers(ctx, optional)
Creates a user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiPostUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiPostUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CardHolderModel**](CardHolderModel.md)|  | 

### Return type

[**UserCardHolderResponse**](user_card_holder_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersAuthChangepassword**
> PostUsersAuthChangepassword(ctx, body)
Updates a user password



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PasswordUpdateModel**](PasswordUpdateModel.md)| Password update object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersAuthClientaccesstoken**
> ClientAccessTokenResponse PostUsersAuthClientaccesstoken(ctx, optional)
Creates a client access token



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiPostUsersAuthClientaccesstokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiPostUsersAuthClientaccesstokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ClientAccessTokenRequest**](ClientAccessTokenRequest.md)|  | 

### Return type

[**ClientAccessTokenResponse**](client_access_token_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersAuthLogin**
> LoginResponseModel PostUsersAuthLogin(ctx, optional)
Logs in a user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiPostUsersAuthLoginOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiPostUsersAuthLoginOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LoginRequestModel**](LoginRequestModel.md)| User login object | 

### Return type

[**LoginResponseModel**](login_response_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersAuthLogout**
> PostUsersAuthLogout(ctx, )
Logs out a user



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersAuthOnetime**
> AccessTokenResponse PostUsersAuthOnetime(ctx, optional)
Creates a one-time token



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiPostUsersAuthOnetimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiPostUsersAuthOnetimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OneTimeRequestModel**](OneTimeRequestModel.md)| One-time object | 

### Return type

[**AccessTokenResponse**](access_token_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersAuthResetpassword**
> PostUsersAuthResetpassword(ctx, optional)
Generates a reset password email



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiPostUsersAuthResetpasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiPostUsersAuthResetpasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ResetUserPasswordEmailModel**](ResetUserPasswordEmailModel.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersAuthResetpasswordToken**
> PostUsersAuthResetpasswordToken(ctx, token, optional)
Resets a user password



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Reset password verification token | 
 **optional** | ***UsersApiPostUsersAuthResetpasswordTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiPostUsersAuthResetpasswordTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ResetUserPasswordModel**](ResetUserPasswordModel.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersAuthVerifyemail**
> PostUsersAuthVerifyemail(ctx, )
Generates an email verification request



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersAuthVerifyemailToken**
> PostUsersAuthVerifyemailToken(ctx, token)
Verifies the email token



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Email verification token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersLookup**
> UserCardHolderListResponse PostUsersLookup(ctx, optional)
Lists all users



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiPostUsersLookupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiPostUsersLookupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UserCardHolderSearchModel**](UserCardHolderSearchModel.md)|  | 
 **count** | **optional.Int32**| Number of users to retrieve | [default to 5]
 **startIndex** | **optional.Int32**| Start index | [default to 0]
 **searchType** | **optional.String**| Search type | 
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

# **PostUsersTokenNotes**
> CardholderNoteResponseModel PostUsersTokenNotes(ctx, token, optional)
Creates a note for the cardholder



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| User token | 
 **optional** | ***UsersApiPostUsersTokenNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiPostUsersTokenNotesOpts struct

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

# **PutUsersToken**
> CardHolderModel PutUsersToken(ctx, token, body)
Updates a specific user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| User token | 
  **body** | [**UserCardHolderUpdateModel**](UserCardHolderUpdateModel.md)| User object | 

### Return type

[**CardHolderModel**](card_holder_model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUsersTokenNotesNotestoken**
> CardholderNoteResponseModel PutUsersTokenNotesNotestoken(ctx, token, notesToken, optional)
Updates a specific note for a cardholder



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| User token | 
  **notesToken** | **string**| Notes token | 
 **optional** | ***UsersApiPutUsersTokenNotesNotestokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiPutUsersTokenNotesNotestokenOpts struct

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

