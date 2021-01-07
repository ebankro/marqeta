/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TerminalModel struct {
	Tid string `json:"tid,omitempty"`
	PartialApprovalCapable string `json:"partial_approval_capable,omitempty"`
	CardholderPresence string `json:"cardholder_presence,omitempty"`
	CardPresence string `json:"card_presence,omitempty"`
	Channel string `json:"channel,omitempty"`
	ProcessingType string `json:"processing_type,omitempty"`
	PinPresent string `json:"pin_present,omitempty"`
}
