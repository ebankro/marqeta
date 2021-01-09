/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ClearingRetryModel struct {
	OriginalFailedTransactionToken string `json:"original_failed_transaction_token"`
	NewNetworkReference string `json:"new_network_reference,omitempty"`
	NewApprovalCode string `json:"new_approval_code,omitempty"`
	NewStan string `json:"new_stan,omitempty"`
	FindOriginalWindowDays int32 `json:"find_original_window_days,omitempty"`
	NewProcessingCode string `json:"new_processing_code,omitempty"`
}