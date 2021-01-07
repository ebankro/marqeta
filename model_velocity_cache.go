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

type VelocityCache struct {
	CardHolderId int64 `json:"cardHolderId,omitempty"`
	VelocityControlId int64 `json:"velocityControlId,omitempty"`
	UsedAmount float32 `json:"usedAmount,omitempty"`
	UsageCount int32 `json:"usageCount,omitempty"`
	MaxId int64 `json:"maxId,omitempty"`
	WindowStartTime time.Time `json:"windowStartTime,omitempty"`
	VcSignature string `json:"vcSignature,omitempty"`
	CreatedTime time.Time `json:"createdTime,omitempty"`
	LastModifiedTime time.Time `json:"lastModifiedTime,omitempty"`
}
