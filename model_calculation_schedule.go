/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CalculationSchedule struct {
	Token string `json:"token"`
	ValueType string `json:"value_type,omitempty"`
	Name string `json:"name"`
	Steps []float32 `json:"steps"`
	StepValues []float32 `json:"step_values"`
}