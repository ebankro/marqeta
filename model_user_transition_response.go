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

type UserTransitionResponse struct {
	Token            string     `json:"token"`
	Status           string     `json:"status"`
	ReasonCode       string     `json:"reason_code"`
	Reason           string     `json:"reason,omitempty"`
	Channel          string     `json:"channel"`
	CreatedTime      *time.Time `json:"created_time,omitempty"`
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	UserToken        string     `json:"user_token,omitempty"`
}
