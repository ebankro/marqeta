/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type Shipping struct {
	Method           string                     `json:"method,omitempty"`
	ReturnAddress    *FulfillmentAddressRequest `json:"return_address,omitempty"`
	RecipientAddress *FulfillmentAddressRequest `json:"recipient_address,omitempty"`
	// 255 char max
	CareOfLine string `json:"care_of_line,omitempty"`
}
