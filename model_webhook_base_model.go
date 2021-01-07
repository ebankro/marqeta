/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type WebhookBaseModel struct {
	Name string `json:"name"`
	Active bool `json:"active,omitempty"`
	Config *WebhookConfigModel `json:"config"`
	// An array of event types
	Events []string `json:"events"`
}
