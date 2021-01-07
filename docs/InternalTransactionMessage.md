# InternalTransactionMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessingStartTimeMillis** | **int64** |  | [optional] [default to null]
**AuthorizationTranlog** | [***InternalAuthorizationTransaction**](internal_authorization_transaction.md) |  | [default to null]
**FundingTranlog** | [***FundingTranlog**](funding_tranlog.md) |  | [optional] [default to null]
**VelocityCaches** | [**[]VelocityCache**](VelocityCache.md) |  | [optional] [default to null]
**Context** | [**map[string]interface{}**](interface{}.md) |  | [optional] [default to null]
**CreditDebitIndicator** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


