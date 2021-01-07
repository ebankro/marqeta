# GatewayProgramFundingSourceResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | 250 char max. Empty string (disabled). Must be HTTPS. | [default to null]
**Version** | **string** |  | [default to null]
**Name** | **string** |  | [default to null]
**Active** | **bool** |  | [optional] [default to null]
**Token** | **string** |  | [default to null]
**CreatedTime** | [**time.Time**](time.Time.md) |  | [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) |  | [default to null]
**Account** | **string** |  | [default to null]
**BasicAuthUsername** | **string** | 50 char max. Required if URL is present | [default to null]
**BasicAuthPassword** | **string** | 50 char max. Required if URL is present. Minimum 20 chars with upper and lowercase letters, numbers, and symbols | [default to null]
**TimeoutMillis** | **int64** | Total timeout in milliseconds for gateway processing | [default to null]
**CustomHeader** | **map[string]string** | Custom headers to be passed along with request | [default to null]
**UseMtls** | **bool** | Use MTLS for funding request | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


