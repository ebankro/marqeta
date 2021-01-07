/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PingResponse struct {
	Success bool `json:"success,omitempty"`
	Version string `json:"version,omitempty"`
	Revision string `json:"revision,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Env string `json:"env,omitempty"`
	Id string `json:"id,omitempty"`
}
