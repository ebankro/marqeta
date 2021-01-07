/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Poi struct {
	Other *OtherPoi `json:"other,omitempty"`
	Ecommerce bool `json:"ecommerce,omitempty"`
	// Default = false
	Atm bool `json:"atm,omitempty"`
}
