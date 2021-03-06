/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

import (
	"time"
)

type CampaignModel struct {
	Active    bool       `json:"active"`
	Name      string     `json:"name"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
	Token     string     `json:"token"`
	// Enclose tokens in brackets
	StoreTokens []string `json:"store_tokens,omitempty"`
}
