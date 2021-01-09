/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CardInventoryRequest struct {
	Token string `json:"token,omitempty"`
	// Package ID assigned by card fulfillment
	PackageId string `json:"package_id"`
	StartingInventory int32 `json:"starting_inventory"`
}