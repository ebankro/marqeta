/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VelocityControlRequest struct {
	Token string `json:"token,omitempty"`
	Name string `json:"name,omitempty"`
	Association *SpendControlAssociation `json:"association,omitempty"`
	MerchantScope *MerchantScope `json:"merchant_scope,omitempty"`
	UsageLimit int32 `json:"usage_limit,omitempty"`
	ApprovalsOnly bool `json:"approvals_only,omitempty"`
	IncludePurchases bool `json:"include_purchases,omitempty"`
	IncludeWithdrawals bool `json:"include_withdrawals,omitempty"`
	IncludeTransfers bool `json:"include_transfers,omitempty"`
	IncludeCashback bool `json:"include_cashback,omitempty"`
	IncludeCredits bool `json:"include_credits,omitempty"`
	CurrencyCode string `json:"currency_code"`
	AmountLimit float32 `json:"amount_limit"`
	VelocityWindow string `json:"velocity_window"`
	Active bool `json:"active,omitempty"`
}
