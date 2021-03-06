/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type BulkRequestModel struct {
	UserTokens      []string `json:"user_tokens,omitempty"`
	BusinessTokens  []string `json:"business_tokens,omitempty"`
	CardTokens      []string `json:"card_tokens,omitempty"`
	KycTokens       []string `json:"kyc_tokens,omitempty"`
	DdaTokens       []string `json:"dda_tokens,omitempty"`
	DepositAccounts []string `json:"deposit_accounts,omitempty"`
}
