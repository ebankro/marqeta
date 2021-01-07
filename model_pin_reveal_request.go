/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PinRevealRequest struct {
	ControlToken string `json:"control_token"`
	// Verification method required
	CardholderVerificationMethod string `json:"cardholder_verification_method"`
}
