/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type AutoReloadUpdateModel struct {
	Token                     string                 `json:"token,omitempty"`
	Active                    bool                   `json:"active,omitempty"`
	FundingSourceToken        string                 `json:"funding_source_token,omitempty"`
	FundingSourceAddressToken string                 `json:"funding_source_address_token,omitempty"`
	Association               *AutoReloadAssociation `json:"association,omitempty"`
	OrderScope                *OrderScope            `json:"order_scope,omitempty"`
	CurrencyCode              string                 `json:"currency_code,omitempty"`
}
