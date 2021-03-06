/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type ProgramReserveTransactionRequest struct {
	IdempotentHash string  `json:"idempotentHash,omitempty"`
	Token          string  `json:"token"`
	Amount         float32 `json:"amount"`
	CurrencyCode   string  `json:"currency_code"`
	Memo           string  `json:"memo,omitempty"`
	Tags           string  `json:"tags,omitempty"`
	Type_          string  `json:"type"`
}
