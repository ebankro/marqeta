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

type PanResponse struct {
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime time.Time `json:"created_time"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime time.Time `json:"last_modified_time"`
	// The unique identifier of the card user
	UserToken string `json:"user_token"`
	// The unique identifier of the card
	CardToken string `json:"card_token"`
}