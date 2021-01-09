/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ReversalModel struct {
	NetworkFees []NetworkFeeModel `json:"network_fees,omitempty"`
	Webhook *Webhook `json:"webhook,omitempty"`
	OriginalTransactionToken string `json:"original_transaction_token"`
	Amount float32 `json:"amount"`
	FindOriginalWindowDays int32 `json:"find_original_window_days,omitempty"`
	IsAdvice bool `json:"is_advice,omitempty"`
}