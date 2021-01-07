/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Pos struct {
	PanEntryMode string `json:"pan_entry_mode,omitempty"`
	PinEntryMode string `json:"pin_entry_mode,omitempty"`
	TerminalId string `json:"terminal_id,omitempty"`
	TerminalAttendance string `json:"terminal_attendance,omitempty"`
	TerminalLocation string `json:"terminal_location,omitempty"`
	CardHolderPresence bool `json:"card_holder_presence,omitempty"`
	CardholderAuthenticationMethod string `json:"cardholder_authentication_method,omitempty"`
	CardPresence bool `json:"card_presence,omitempty"`
	PinPresent bool `json:"pin_present,omitempty"`
	TerminalType string `json:"terminal_type,omitempty"`
	CardDataInputCapability string `json:"card_data_input_capability,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Zip string `json:"zip,omitempty"`
	PartialApprovalCapable bool `json:"partial_approval_capable,omitempty"`
	PurchaseAmountOnly bool `json:"purchase_amount_only,omitempty"`
	IsRecurring bool `json:"is_recurring,omitempty"`
	IsInstallment bool `json:"is_installment,omitempty"`
}
