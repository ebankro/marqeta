/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type Funding struct {
	Amount        float32                    `json:"amount,omitempty"`
	Source        *FundingSourceModel        `json:"source"`
	SourceAddress *CardholderAddressResponse `json:"source_address,omitempty"`
	GatewayLog    *GatewayLogModel           `json:"gateway_log,omitempty"`
}
