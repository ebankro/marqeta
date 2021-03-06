/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type BillingAddress struct {
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Address       string `json:"address,omitempty"`
	Zip           string `json:"zip,omitempty"`
	CompressedZip string `json:"compressed_zip,omitempty"`
}
