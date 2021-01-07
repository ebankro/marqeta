# AutoReloadResponseModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | [optional] [default to null]
**Active** | **bool** |  | [optional] [default to null]
**FundingSourceToken** | **string** | Required when order scope is GPA | [optional] [default to null]
**FundingSourceAddressToken** | **string** |  | [optional] [default to null]
**Association** | [***AutoReloadAssociation**](auto_reload_association.md) |  | [optional] [default to null]
**OrderScope** | [***OrderScope**](order_scope.md) | either GPA or MSA is required | [default to null]
**CurrencyCode** | **string** |  | [default to null]
**CreatedTime** | [**time.Time**](time.Time.md) | yyyy-MM-ddTHH:mm:ssZ | [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | yyyy-MM-ddTHH:mm:ssZ | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


