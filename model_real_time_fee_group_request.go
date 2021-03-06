/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type RealTimeFeeGroupRequest struct {
	Name      string   `json:"name,omitempty"`
	Active    bool     `json:"active,omitempty"`
	FeeTokens []string `json:"fee_tokens,omitempty"`
}
