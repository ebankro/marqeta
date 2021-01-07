/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RealTimeFeeAssessment struct {
	TransactionType string `json:"transaction_type,omitempty"`
	InternationalEnabled bool `json:"international_enabled,omitempty"`
	DomesticEnabled bool `json:"domestic_enabled,omitempty"`
}
