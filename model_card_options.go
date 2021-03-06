/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type CardOptions struct {
	Cvv            string          `json:"cvv,omitempty"`
	CardPresent    bool            `json:"card_present,omitempty"`
	Expiration     string          `json:"expiration,omitempty"`
	BillingAddress *BillingAddress `json:"billing_address,omitempty"`
}
