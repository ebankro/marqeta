/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TokenServiceProvider struct {
	TokenReferenceId string `json:"token_reference_id,omitempty"`
	// 50 char max
	PanReferenceId string `json:"pan_reference_id"`
	CorrelationId string `json:"correlation_id,omitempty"`
	TokenRequestorId string `json:"token_requestor_id,omitempty"`
	TokenRequestorName string `json:"token_requestor_name,omitempty"`
	TokenType string `json:"token_type,omitempty"`
	TokenPan string `json:"token_pan,omitempty"`
	TokenExpiration string `json:"token_expiration,omitempty"`
	TokenScore string `json:"token_score,omitempty"`
	TokenAssuranceLevel string `json:"token_assurance_level,omitempty"`
	TokenEligibilityDecision string `json:"token_eligibility_decision,omitempty"`
}