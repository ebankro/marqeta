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

type ChargebackTransitionResponse struct {
	Token            string `json:"token"`
	State            string `json:"state"`
	PreviousState    string `json:"previous_state"`
	Channel          string `json:"channel"`
	ChargebackToken  string `json:"chargeback_token"`
	Reason           string `json:"reason,omitempty"`
	TransactionToken string `json:"transaction_token,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime *time.Time `json:"created_time"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime *time.Time `json:"last_modified_time"`
	Type_            string     `json:"type"`
	Amount           float32    `json:"amount,omitempty"`
}
