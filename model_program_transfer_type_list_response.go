/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type ProgramTransferTypeListResponse struct {
	Count      int32                        `json:"count,omitempty"`
	StartIndex int32                        `json:"start_index,omitempty"`
	EndIndex   int32                        `json:"end_index,omitempty"`
	IsMore     bool                         `json:"is_more,omitempty"`
	Data       []ProgramTransferTypeReponse `json:"data,omitempty"`
}
