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

type PaymentCardResponseModel struct {
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime time.Time `json:"created_time"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime time.Time `json:"last_modified_time"`
	Type_ string `json:"type"`
	Token string `json:"token"`
	AccountSuffix string `json:"account_suffix"`
	AccountType string `json:"account_type"`
	Active bool `json:"active"`
	IsDefaultAccount bool `json:"is_default_account"`
	ExpDate string `json:"exp_date"`
	// Required if 'business_token' is not present
	UserToken string `json:"user_token,omitempty"`
	// Required if 'user_token' is not present
	BusinessToken string `json:"business_token,omitempty"`
}
