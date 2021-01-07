/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ChargebackUpdateRequest struct {
	// Either 'reason_description' or 'reason_code' is required
	ReasonDescription string `json:"reason_description,omitempty"`
	// Either 'reason_code' or 'reason_description' is required
	ReasonCode string `json:"reason_code,omitempty"`
}
