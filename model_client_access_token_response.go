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

type ClientAccessTokenResponse struct {
	Application *Application `json:"application,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	Created time.Time `json:"created"`
	// yyyy-MM-ddTHH:mm:ssZ
	Expires time.Time `json:"expires"`
	Token string `json:"token,omitempty"`
	CardToken string `json:"card_token,omitempty"`
}
