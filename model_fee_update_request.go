/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type FeeUpdateRequest struct {
	Name string `json:"name,omitempty"`
	Amount float32 `json:"amount,omitempty"`
	Tags string `json:"tags,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
	Active bool `json:"active,omitempty"`
	RealTimeAssessment *RealTimeFeeAssessmentRequest `json:"real_time_assessment,omitempty"`
}