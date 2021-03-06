/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type Issuer struct {
	Success           bool     `json:"success,omitempty"`
	FraudScore        int32    `json:"fraud_score,omitempty"`
	FraudRating       string   `json:"fraud_rating,omitempty"`
	RuleViolations    []string `json:"rule_violations,omitempty"`
	FraudScoreReasons []string `json:"fraud_score_reasons,omitempty"`
	RecommendedAction string   `json:"recommended_action,omitempty"`
	Model             string   `json:"model,omitempty"`
	Message           string   `json:"message,omitempty"`
}
