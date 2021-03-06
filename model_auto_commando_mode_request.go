/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type AutoCommandoModeRequest struct {
	ProgramGatewayFundingSource *InternalFundingSource       `json:"program_gateway_funding_source"`
	GatewayResponse             *InternalGatewayResponse     `json:"gateway_response"`
	VelocityControlRequest      *VelocityControlCheckRequest `json:"velocity_control_request"`
	MccGroups                   []string                     `json:"mcc_groups,omitempty"`
}
