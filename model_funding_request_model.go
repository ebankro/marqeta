/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type FundingRequestModel struct {
	UserToken   string  `json:"user_token"`
	OrderNumber string  `json:"order_number"`
	Amount      float32 `json:"amount"`
	// Payment card or ACH account number
	FundingSource  string `json:"funding_source,omitempty"`
	FundingAddress string `json:"funding_address,omitempty"`
	Fundgpadetail  string `json:"fundgpadetail,omitempty"`
	OrderToken     string `json:"order_token,omitempty"`
	CurrencyCode   string `json:"currency_code,omitempty"`
}
