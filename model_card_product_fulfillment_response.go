/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type CardProductFulfillmentResponse struct {
	Shipping                *ShippingInformationResponse `json:"shipping,omitempty"`
	CardPersonalization     *CardPersonalization         `json:"card_personalization"`
	PaymentInstrument       string                       `json:"payment_instrument,omitempty"`
	PackageId               string                       `json:"package_id,omitempty"`
	AllZeroCardSecurityCode bool                         `json:"all_zero_card_security_code,omitempty"`
	BinPrefix               string                       `json:"bin_prefix,omitempty"`
	BulkShip                bool                         `json:"bulk_ship,omitempty"`
	PanLength               string                       `json:"pan_length,omitempty"`
	FulfillmentProvider     string                       `json:"fulfillment_provider,omitempty"`
	AllowCardCreation       bool                         `json:"allow_card_creation,omitempty"`
	UppercaseNameLines      bool                         `json:"uppercase_name_lines,omitempty"`
	EnableOfflinePin        bool                         `json:"enable_offline_pin,omitempty"`
}
