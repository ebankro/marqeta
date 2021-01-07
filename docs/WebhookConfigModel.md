# WebhookConfigModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Must be HTTPS | [default to null]
**Secret** | **string** | Must contain upper and lowercase letters, numbers, and symbols | [optional] [default to null]
**BasicAuthUsername** | **string** |  | [default to null]
**BasicAuthPassword** | **string** | Required if URL is present; must contain upper and lowercase letters, numbers, and symbols | [default to null]
**CustomHeader** | **map[string]string** | Custom headers to be passed along with request | [optional] [default to null]
**UseMtls** | **bool** | Use MTLS for webhook | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


