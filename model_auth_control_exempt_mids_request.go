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

type AuthControlExemptMidsRequest struct {
	Token       string                   `json:"token,omitempty"`
	Name        string                   `json:"name"`
	Association *SpendControlAssociation `json:"association,omitempty"`
	Mid         string                   `json:"mid"`
	StartTime   *time.Time               `json:"start_time,omitempty"`
	EndTime     *time.Time               `json:"end_time,omitempty"`
}
