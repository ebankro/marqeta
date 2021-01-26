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

type DigitalWalletTokenTransitionResponse struct {
	Token              string                  `json:"token"`
	DigitalWalletToken *DigitalWalletTokenHash `json:"digital_wallet_token"`
	CardSwap           *CardSwapHash           `json:"card_swap,omitempty"`
	Type_              string                  `json:"type"`
	Channel            string                  `json:"channel"`
	State              string                  `json:"state"`
	FulfillmentStatus  string                  `json:"fulfillment_status"`
	Reason             string                  `json:"reason,omitempty"`
	ReasonCode         string                  `json:"reason_code,omitempty"`
	CreatedTime        *time.Time              `json:"created_time,omitempty"`
}
