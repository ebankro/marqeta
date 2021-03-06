/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type DisputeModel struct {
	Reason                   string `json:"reason,omitempty"`
	CaseManagementIdentifier string `json:"case_management_identifier,omitempty"`
}
