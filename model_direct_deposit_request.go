/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type DirectDepositRequest struct {
	Token string `json:"token,omitempty"`
	Amount float32 `json:"amount"`
	Type_ string `json:"type"`
	AccountNumber string `json:"account_number"`
	SettlementDate time.Time `json:"settlement_date"`
	StandardEntryClassCode string `json:"standard_entry_class_code,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	CompanyDiscretionaryData string `json:"company_discretionary_data,omitempty"`
	CompanyIdentification string `json:"company_identification,omitempty"`
	CompanyEntryDescription string `json:"company_entry_description,omitempty"`
	IndividualIdentificationNumber string `json:"individual_identification_number,omitempty"`
	IndividualName string `json:"individual_name,omitempty"`
}