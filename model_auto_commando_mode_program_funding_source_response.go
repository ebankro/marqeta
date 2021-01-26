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

type AutoCommandoModeProgramFundingSourceResponse struct {
	Name             string     `json:"name"`
	Active           bool       `json:"active,omitempty"`
	Token            string     `json:"token"`
	CreatedTime      *time.Time `json:"created_time"`
	LastModifiedTime *time.Time `json:"last_modified_time"`
	Account          string     `json:"account"`
}
