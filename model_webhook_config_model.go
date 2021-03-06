/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type WebhookConfigModel struct {
	// Must be HTTPS
	Url string `json:"url"`
	// Must contain upper and lowercase letters, numbers, and symbols
	Secret            string `json:"secret,omitempty"`
	BasicAuthUsername string `json:"basic_auth_username"`
	// Required if URL is present; must contain upper and lowercase letters, numbers, and symbols
	BasicAuthPassword string `json:"basic_auth_password"`
	// Custom headers to be passed along with request
	CustomHeader map[string]string `json:"custom_header,omitempty"`
	// Use MTLS for webhook
	UseMtls bool `json:"use_mtls,omitempty"`
}
