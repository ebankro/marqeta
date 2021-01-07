# MsaOrderRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | [optional] [default to null]
**CampaignToken** | **string** |  | [default to null]
**UserToken** | **string** | Required if &#39;business_token&#39; is null | [optional] [default to null]
**BusinessToken** | **string** | Required if &#39;user_token&#39; is null | [optional] [default to null]
**CurrencyCode** | **string** |  | [default to null]
**PurchaseAmount** | **float32** |  | [default to null]
**RewardAmount** | **float32** | default is 0.00 | [optional] [default to null]
**RewardTriggerAmount** | **float32** | default is 0.01 | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**FundingSourceToken** | **string** |  | [default to null]
**FundingSourceAddressToken** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


