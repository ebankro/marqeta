# DirectDepositAccountRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserToken** | **string** | Required if &#39;business_token&#39; is null | [optional] [default to null]
**BusinessToken** | **string** | Required if &#39;user_token&#39; is null | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]
**AllowImmediateCredit** | **bool** |  | [optional] [default to null]
**Token** | **string** |  | [optional] [default to null]
**CustomerDueDiligence** | [**[]CustomerDueDiligenceRequest**](customer_due_diligence_request.md) | Required if account type &#x3D; Checking | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


