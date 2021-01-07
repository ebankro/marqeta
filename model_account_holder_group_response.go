/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AccountHolderGroupResponse struct {
	// 36 char max
	Token string `json:"token,omitempty"`
	// 40 char max
	Name string `json:"name,omitempty"`
	Config *AccountHolderGroupConfig `json:"config,omitempty"`
}
