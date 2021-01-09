/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BankTransferRequestModel struct {
	Token string `json:"token,omitempty"`
	Amount float32 `json:"amount"`
	Fees []float32 `json:"fees,omitempty"`
	Memo string `json:"memo,omitempty"`
	FundingSourceToken string `json:"funding_source_token"`
	Type_ string `json:"type,omitempty"`
	// default = USD
	CurrencyCode string `json:"currency_code,omitempty"`
	// default = STANDARD
	TransferSpeed string `json:"transfer_speed,omitempty"`
	StandardEntryClassCode string `json:"standard_entry_class_code,omitempty"`
}