# GatewayProgramFundingSourceUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [optional] [default to null]
**Url** | **string** | Empty string (disabled); must be HTTPS | [default to null]
**Active** | **bool** |  | [optional] [default to null]
**BasicAuthUsername** | **string** | Required if URL is present | [default to null]
**BasicAuthPassword** | **string** | Required if URL is present; must contain upper and lowercase letters, numbers, and symbols | [default to null]
**TimeoutMillis** | **int64** | Total timeout in milliseconds for gateway processing | [optional] [default to null]
**CustomHeader** | **map[string]string** | Custom headers | [optional] [default to null]
**UseMtls** | **bool** | Use MTLS for funding request | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


