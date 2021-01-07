/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GatewayResponse struct {
	Code string `json:"code"`
	Data *JitProgramResponse `json:"data,omitempty"`
}
