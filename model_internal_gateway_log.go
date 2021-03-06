/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type InternalGatewayLog struct {
	GatewayMerchant        *InternalGatewayMerchant `json:"gatewayMerchant,omitempty"`
	OrderId                string                   `json:"orderId,omitempty"`
	RequestMethod          string                   `json:"requestMethod,omitempty"`
	GatewayRequestMethod   string                   `json:"gatewayRequestMethod,omitempty"`
	ResponseCode           string                   `json:"responseCode,omitempty"`
	ResponseSubCode        string                   `json:"responseSubCode,omitempty"`
	ResponseReasonCode     string                   `json:"responseReasonCode,omitempty"`
	ResponseMessage        string                   `json:"responseMessage,omitempty"`
	GatewayResponseMessage string                   `json:"gatewayResponseMessage,omitempty"`
	ResponseStatus         string                   `json:"responseStatus,omitempty"`
	GatewayTransactionId   string                   `json:"gatewayTransactionId,omitempty"`
	Amount                 float32                  `json:"amount,omitempty"`
	ApiDuration            int64                    `json:"apiDuration,omitempty"`
	GatewayDuration        int64                    `json:"gatewayDuration,omitempty"`
	Memo                   string                   `json:"memo,omitempty"`
	GatewayVersion         string                   `json:"gatewayVersion,omitempty"`
	FundingSource          *InternalFundingSource   `json:"funding_source"`
	InternalUser           *InternalUser            `json:"internalUser,omitempty"`
}
