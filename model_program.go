/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Program struct {
	ProgramId string `json:"program_id"`
	ShortCode string `json:"short_code"`
	LongCode string `json:"long_code"`
}
