/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DigitalWalletTokenRequestorMapModel struct {
	Network string `json:"network"`
	DigitalWalletTokenRequestorId string `json:"digital_wallet_token_requestor_id"`
	DigitalWalletTokenRequestorName string `json:"digital_wallet_token_requestor_name"`
	Token string `json:"token,omitempty"`
}
