/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type Application struct {
	Token            string `json:"token,omitempty"`
	Program          string `json:"program,omitempty"`
	Environment      string `json:"environment,omitempty"`
	ProgramShortCode string `json:"program_short_code,omitempty"`
	ClientApiBaseUrl string `json:"client_api_base_url,omitempty"`
	AssetsUrl        string `json:"assets_url,omitempty"`
	AccessCode       string `json:"access_code,omitempty"`
}
