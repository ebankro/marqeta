/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type AddressVerification struct {
	Name          string `json:"name,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
	Zip           string `json:"zip,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
}
