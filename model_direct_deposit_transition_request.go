/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type DirectDepositTransitionRequest struct {
	Token              string `json:"token,omitempty"`
	Channel            string `json:"channel"`
	IdempotentHash     string `json:"idempotentHash,omitempty"`
	DirectDepositToken string `json:"direct_deposit_token"`
	State              string `json:"state"`
	ReasonCode         string `json:"reason_code,omitempty"`
	Reason             string `json:"reason"`
}
