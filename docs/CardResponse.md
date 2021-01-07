# CardResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | [**time.Time**](time.Time.md) | yyyy-MM-ddTHH:mm:ssZ | [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | yyyy-MM-ddTHH:mm:ssZ | [default to null]
**Token** | **string** | 36 char max | [default to null]
**UserToken** | **string** | 36 char max | [default to null]
**CardProductToken** | **string** | 36 char max | [default to null]
**LastFour** | **string** |  | [default to null]
**Pan** | **string** |  | [default to null]
**Expiration** | **string** |  | [default to null]
**ExpirationTime** | [**time.Time**](time.Time.md) | yyyy-MM-ddTHH:mm:ssZ | [default to null]
**CvvNumber** | **string** |  | [optional] [default to null]
**ChipCvvNumber** | **string** |  | [optional] [default to null]
**Barcode** | **string** |  | [default to null]
**PinIsSet** | **bool** |  | [default to null]
**State** | **string** |  | [default to null]
**StateReason** | **string** |  | [default to null]
**FulfillmentStatus** | **string** |  | [default to null]
**ReissuePanFromCardToken** | **string** |  | [optional] [default to null]
**Fulfillment** | [***CardFulfillmentResponse**](CardFulfillmentResponse.md) |  | [optional] [default to null]
**BulkIssuanceToken** | **string** |  | [optional] [default to null]
**TranslatePinFromCardToken** | **string** |  | [optional] [default to null]
**ActivationActions** | [***ActivationActions**](activation_actions.md) |  | [optional] [default to null]
**InstrumentType** | **string** |  | [optional] [default to null]
**Expedite** | **bool** |  | [optional] [default to null]
**Metadata** | **map[string]string** |  | [optional] [default to null]
**ContactlessExemptionCounter** | **int32** |  | [optional] [default to null]
**ContactlessExemptionTotalAmount** | **float32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


