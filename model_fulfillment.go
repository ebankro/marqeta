/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type Fulfillment struct {
	Shipping              *Shipping            `json:"shipping,omitempty"`
	CardPersonalization   *CardPersonalization `json:"card_personalization"`
	CardFulfillmentReason string               `json:"card_fulfillment_reason,omitempty"`
}
