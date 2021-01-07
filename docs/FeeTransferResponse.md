# FeeTransferResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | **string** |  | [optional] [default to null]
**Fees** | [**[]FeeDetail**](fee_detail.md) |  | [default to null]
**Token** | **string** | 36 char max | [default to null]
**UserToken** | **string** | Required if &#39;business_token&#39; is null | [default to null]
**BusinessToken** | **string** | Required if &#39;user_token&#39; is null | [default to null]
**CreatedTime** | [**time.Time**](time.Time.md) | yyyy-MM-ddTHH:mm:ssZ | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


