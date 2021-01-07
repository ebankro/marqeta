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

type ProgramTransferTypeReponse struct {
	Token string `json:"token"`
	ProgramFundingSourceToken string `json:"program_funding_source_token"`
	Tags string `json:"tags,omitempty"`
	Memo string `json:"memo,omitempty"`
	CreatedTime time.Time `json:"created_time,omitempty"`
	LastModifiedTime time.Time `json:"last_modified_time,omitempty"`
}
