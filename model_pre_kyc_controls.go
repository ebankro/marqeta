/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PreKycControls struct {
	CashAccessEnabled bool `json:"cash_access_enabled,omitempty"`
	InternationalEnabled bool `json:"international_enabled,omitempty"`
	// Minimum is 0.01
	BalanceMax float32 `json:"balance_max,omitempty"`
	EnableNonProgramLoads bool `json:"enable_non_program_loads,omitempty"`
	IsReloadablePreKyc bool `json:"is_reloadable_pre_kyc,omitempty"`
}