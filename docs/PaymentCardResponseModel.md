# PaymentCardResponseModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | [**time.Time**](time.Time.md) | yyyy-MM-ddTHH:mm:ssZ | [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | yyyy-MM-ddTHH:mm:ssZ | [default to null]
**Type_** | **string** |  | [default to null]
**Token** | **string** |  | [default to null]
**AccountSuffix** | **string** |  | [default to null]
**AccountType** | **string** |  | [default to null]
**Active** | **bool** |  | [default to null]
**IsDefaultAccount** | **bool** |  | [default to null]
**ExpDate** | **string** |  | [default to null]
**UserToken** | **string** | Required if &#39;business_token&#39; is not present | [optional] [default to null]
**BusinessToken** | **string** | Required if &#39;user_token&#39; is not present | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


