/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type JitFunding struct {
	PaymentcardFundingSource    *JitFundingPaymentcardFundingSource    `json:"paymentcard_funding_source,omitempty"`
	ProgramgatewayFundingSource *JitFundingProgramgatewayFundingSource `json:"programgateway_funding_source,omitempty"`
	ProgramFundingSource        *JitFundingProgramFundingSource        `json:"program_funding_source,omitempty"`
}
