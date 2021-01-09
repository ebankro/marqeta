/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProgramTransferTypeRequest struct {
	Token string `json:"token,omitempty"`
	ProgramFundingSourceToken string `json:"program_funding_source_token"`
	Tags string `json:"tags,omitempty"`
	Memo string `json:"memo"`
}