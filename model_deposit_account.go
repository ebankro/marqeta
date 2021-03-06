/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type DepositAccount struct {
	Token                string `json:"token"`
	UserToken            string `json:"user_token,omitempty"`
	BusinessToken        string `json:"business_token,omitempty"`
	AccountNumber        string `json:"account_number"`
	RoutingNumber        string `json:"routing_number"`
	AllowImmediateCredit bool   `json:"allow_immediate_credit,omitempty"`
}
