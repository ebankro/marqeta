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

type AccessTokenResponse struct {
	Token string `json:"token,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	Expires time.Time `json:"expires"`
	Application *Application `json:"application,omitempty"`
	UserToken string `json:"user_token,omitempty"`
	MasterRoles []string `json:"master_roles,omitempty"`
	TokenType string `json:"token_type,omitempty"`
	OneTime bool `json:"one_time,omitempty"`
}