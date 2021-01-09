/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type OfferUpdateModel struct {
	Active bool `json:"active,omitempty"`
	// 255 char max
	Name string `json:"name,omitempty"`
	// yyyy-MM-ddThh:mm:ssZ
	StartDate string `json:"start_date,omitempty"`
	// yyyy-MM-ddThh:mm:ssZ
	EndDate string `json:"end_date,omitempty"`
}