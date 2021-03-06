# AdvancedClearingRequestModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkFees** | [**[]NetworkFeeModel**](network_fee_model.md) |  | [optional] [default to null]
**Webhook** | [***Webhook**](webhook.md) |  | [optional] [default to null]
**Mid** | **string** |  | [optional] [default to null]
**Amount** | **float32** |  | [default to null]
**SourceAmount** | **float32** |  | [default to null]
**ReconciliationAmount** | **float32** |  | [default to null]
**ReplacementAmount** | **float32** |  | [optional] [default to null]
**CardholderBillingAmount** | **float32** |  | [optional] [default to null]
**Cashback** | **float32** |  | [optional] [default to null]
**LocalTransactionAmount** | **float32** |  | [optional] [default to null]
**LocalCurrencyCode** | **string** |  | [optional] [default to null]
**CardholderBillingConversionRate** | **float64** |  | [optional] [default to null]
**CardholderBillingCurrency** | **string** |  | [optional] [default to null]
**CardToken** | **string** |  | [default to null]
**CardHash** | **string** |  | [default to null]
**AcquirerReferenceId** | **string** |  | [optional] [default to null]
**Rrn** | **string** |  | [optional] [default to null]
**Stan** | **string** |  | [optional] [default to null]
**ProcessingCode** | **string** |  | [optional] [default to null]
**AcquirerFee** | [***MoneyModel**](MoneyModel.md) |  | [optional] [default to null]
**IssuerFee** | [***MoneyModel**](MoneyModel.md) |  | [optional] [default to null]
**FunctionCode** | **string** |  | [optional] [default to null]
**ReasonCode** | **string** |  | [optional] [default to null]
**ApprovalCode** | **string** |  | [optional] [default to null]
**TransactionDate** | [**time.Time**](time.Time.md) | yyyy-MM-dd, yyyy-MM-ddThh:mm:ssZ | [optional] [default to null]
**LocalTransactionDate** | [**time.Time**](time.Time.md) | yyyy-MM-dd, yyyy-MM-ddThh:mm:ssZ | [optional] [default to null]
**SettlementDate** | [**time.Time**](time.Time.md) | yyyy-MM-dd, yyyy-MM-ddThh:mm:ssZ | [optional] [default to null]
**NetworkReferenceId** | **string** |  | [optional] [default to null]
**FindOriginalWindowDays** | **int32** |  | [optional] [default to null]
**BatchNumber** | **string** |  | [optional] [default to null]
**BatchFileName** | **string** |  | [optional] [default to null]
**SequenceNumber** | **string** |  | [optional] [default to null]
**MultiClearingCount** | **string** |  | [optional] [default to null]
**Network** | **string** |  | [optional] [default to null]
**SubNetwork** | **string** | Defaults to VISANET if network is VISA | [optional] [default to null]
**CardAcceptor** | [***CardAcceptorModel**](card_acceptor_model.md) |  | [optional] [default to null]
**CurrencyCode** | **string** |  | [default to null]
**OriginalDataElements** | [***OriginalDataElements**](original_data_elements.md) |  | [optional] [default to null]
**PrecedingRelatedTransactionToken** | **string** |  | [optional] [default to null]
**SendExpirationDate** | **bool** |  | [optional] [default to null]
**SimulateBatchForClearingRecordHash** | **bool** |  | [optional] [default to null]
**CrossBorderIndicator** | **string** |  | [optional] [default to null]
**TokenPan** | **string** |  | [optional] [default to null]
**PaymentChannelIndicator** | **string** |  | [optional] [default to null]
**IsInstallment** | **bool** |  | [optional] [default to null]
**IsRecurring** | **bool** |  | [optional] [default to null]
**NetworkMetadata** | [***NetworkMetadata**](network_metadata.md) |  | [optional] [default to null]
**AuthorizationSourceCode** | **string** |  | [optional] [default to null]
**InterchangeRateDescriptor** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


