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

type AcceptedCountriesModel struct {
	Token            string     `json:"token,omitempty"`
	Name             string     `json:"name"`
	IsWhitelist      bool       `json:"is_whitelist"`
	CountryCodes     []string   `json:"country_codes"`
	CreatedTime      *time.Time `json:"created_time,omitempty"`
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
}
