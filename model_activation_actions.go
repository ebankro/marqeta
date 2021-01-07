/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ActivationActions struct {
	TerminateReissuedSourceCard bool `json:"terminate_reissued_source_card,omitempty"`
	SwapDigitalWalletTokensFromCardToken string `json:"swap_digital_wallet_tokens_from_card_token,omitempty"`
}
