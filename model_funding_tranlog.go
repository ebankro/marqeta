/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type FundingTranlog struct {
	ReturnedBalances string `json:"returnedBalances,omitempty"`
	Tags string `json:"tags,omitempty"`
	TransactionType string `json:"transactionType"`
	Token string `json:"token"`
	RefTransaction *InternalAuthorizationTransaction `json:"ref_transaction,omitempty"`
	Node string `json:"node"`
	Network string `json:"network"`
	SubNetwork string `json:"subNetwork"`
	Mid string `json:"mid,omitempty"`
	Tid string `json:"tid,omitempty"`
	Stan string `json:"stan,omitempty"`
	CaName string `json:"ca_name,omitempty"`
	CaStreet string `json:"ca_street,omitempty"`
	CaZip string `json:"ca_zip,omitempty"`
	CaCity string `json:"ca_city,omitempty"`
	CaRegion string `json:"ca_region,omitempty"`
	CaCountry string `json:"ca_country,omitempty"`
	FunctionCode string `json:"functionCode,omitempty"`
	ReasonCode string `json:"reasonCode,omitempty"`
	ResponseCode string `json:"responseCode,omitempty"`
	ApprovalNumber string `json:"approvalNumber,omitempty"`
	DisplayMessage string `json:"displayMessage,omitempty"`
	Date time.Time `json:"date,omitempty"`
	TransmissionDate time.Time `json:"transmissionDate,omitempty"`
	LocalTransactionDate time.Time `json:"localTransactionDate,omitempty"`
	CaptureDate time.Time `json:"captureDate,omitempty"`
	SettlementDate time.Time `json:"settlementDate,omitempty"`
	Itc string `json:"itc,omitempty"`
	Irc string `json:"irc,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
	Amount float32 `json:"amount,omitempty"`
	AdditionalAmount float32 `json:"additionalAmount,omitempty"`
	AcquirerFee float32 `json:"acquirerFee,omitempty"`
	IssuerFee float32 `json:"issuerFee,omitempty"`
	Rc string `json:"rc,omitempty"`
	Extrc string `json:"extrc,omitempty"`
	Duration int32 `json:"duration,omitempty"`
	Cardholder *InternalUser `json:"cardholder,omitempty"`
	ActingCardholder *InternalUser `json:"actingCardholder,omitempty"`
	Card *InternalCard `json:"card,omitempty"`
	Account *InternalAccount `json:"account,omitempty"`
	Account2 *InternalAccount `json:"account2,omitempty"`
	Mcc string `json:"mcc,omitempty"`
	NetworkReferenceId string `json:"networkReferenceId,omitempty"`
	AcquirerReferenceId string `json:"acquirerReferenceId,omitempty"`
	RetrievalReferenceNumber string `json:"retrievalReferenceNumber,omitempty"`
	ForwardingInstId string `json:"forwardingInstId,omitempty"`
	NetworkMid string `json:"networkMid,omitempty"`
	RequestAmount float32 `json:"requestAmount,omitempty"`
	TransactionState string `json:"transactionState,omitempty"`
	RemoteHost string `json:"remoteHost,omitempty"`
	ResponseAmount float32 `json:"responseAmount,omitempty"`
	ExpirationTime time.Time `json:"expirationTime,omitempty"`
	IncomingNetworkRequestITC string `json:"incomingNetworkRequestITC,omitempty"`
	DigitalWalletToken *InternalDigitalWallet `json:"digital_wallet_token,omitempty"`
	TranlogAttributes map[string]string `json:"tranlogAttributes,omitempty"`
	Payload *TransactionModel `json:"payload,omitempty"`
	Layer int32 `json:"layer,omitempty"`
	TransactionName string `json:"transaction_name,omitempty"`
	Originator string `json:"originator,omitempty"`
	Acquirer string `json:"acquirer,omitempty"`
	Gpaorder *InternalGpaOrder `json:"gpaorder"`
	GatewayLog *InternalGatewayLog `json:"gatewayLog"`
}
