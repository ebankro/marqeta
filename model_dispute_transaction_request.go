/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type DisputeTransactionRequest struct {
	Token                     string  `json:"token"`
	Pan                       string  `json:"pan"`
	ReferenceTransactionToken string  `json:"reference_transaction_token"`
	Reason                    string  `json:"reason"`
	CashAmount                float32 `json:"cash_amount,omitempty"`
	PendingAmount             float32 `json:"pending_amount,omitempty"`
	IncludeAcquirerFees       bool    `json:"include_acquirer_fees,omitempty"`
	CaseManagementIdentifier  string  `json:"case_management_identifier,omitempty"`
}
