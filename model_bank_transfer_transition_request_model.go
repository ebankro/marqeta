/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BankTransferTransitionRequestModel struct {
	Token string `json:"token,omitempty"`
	BankTransferToken string `json:"bank_transfer_token"`
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
	Channel string `json:"channel"`
	BatchNumber string `json:"batch_number,omitempty"`
	ProgramReserveToken string `json:"program_reserve_token,omitempty"`
}