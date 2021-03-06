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

type DigitalWalletApplePayProvisionResponse struct {
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime *time.Time `json:"created_time"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime   *time.Time `json:"last_modified_time"`
	CardToken          string     `json:"card_token"`
	EncryptedPassData  string     `json:"encrypted_pass_data"`
	ActivationData     string     `json:"activation_data"`
	EphemeralPublicKey string     `json:"ephemeral_public_key"`
}
