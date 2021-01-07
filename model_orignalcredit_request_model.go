/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type OrignalcreditRequestModel struct {
	Amount float32 `json:"amount"`
	CardToken string `json:"card_token"`
	Mid string `json:"mid"`
	ScreeningScore string `json:"screening_score,omitempty"`
	CardAcceptor *CardAcceptorModel `json:"card_acceptor,omitempty"`
	Type_ string `json:"type"`
	SenderData *OriginalCreditSenderData `json:"sender_data,omitempty"`
	Webhook *Webhook `json:"webhook,omitempty"`
}
