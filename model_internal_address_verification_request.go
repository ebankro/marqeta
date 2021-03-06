/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type InternalAddressVerificationRequest struct {
	Compressed bool `json:"compressed,omitempty"`
	// cardholder address
	Address string `json:"address,omitempty"`
	// cardholder city
	City string `json:"city,omitempty"`
	// cardholder state
	State string `json:"state,omitempty"`
	// cardholder country
	Country string `json:"country,omitempty"`
	// cardholder postalcode
	Postalcode         string `json:"postalcode,omitempty"`
	IsCompressed       bool   `json:"is_compressed,omitempty"`
	CardholderName     string `json:"cardholder_name,omitempty"`
	ConfigProviderType string `json:"config_provider_type,omitempty"`
}
