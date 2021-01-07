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

type OfferOrderAggregatedBalances struct {
	CurrencyCode string `json:"currency_code"`
	LedgerBalance float32 `json:"ledger_balance"`
	AvailableBalance float32 `json:"available_balance"`
	CreditBalance float32 `json:"credit_balance"`
	CachedBalance float32 `json:"cached_balance"`
	PendingCredits float32 `json:"pending_credits"`
	ImpactedAmount float32 `json:"impacted_amount,omitempty"`
	Balances map[string]OfferOrderAggregatedBalances `json:"balances"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}
