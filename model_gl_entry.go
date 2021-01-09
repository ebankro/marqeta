/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GlEntry struct {
	Detail string `json:"detail,omitempty"`
	Tag string `json:"tag,omitempty"`
	Amount float32 `json:"amount"`
	Layer string `json:"layer"`
	Account string `json:"account"`
	Type_ string `json:"type"`
}